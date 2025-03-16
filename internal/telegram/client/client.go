package client

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/LigeronAhill/planify/internal/e"
	"github.com/LigeronAhill/planify/internal/telegram/commands"
	"github.com/LigeronAhill/planify/internal/telegram/state"
	"github.com/LigeronAhill/planify/internal/telegram/types"
)

// https://api.telegram.org/bot<token>/METHOD_NAME
const BASE_PATH = "https://api.telegram.org"

type Bot struct {
	client     *http.Client
	token      string
	baseUrl    *url.URL
	limit      int
	offset     int
	usersState map[uint]state.UserState
}

func New(ctx context.Context, token string) (*Bot, error) {
	op := "создание клиента"
	slog.Info(op)
	client := http.DefaultClient
	basePath, err := url.Parse(BASE_PATH)
	if err != nil {
		return nil, e.Wrap(op, err)
	}
	tokenString := fmt.Sprintf("bot%s/", token)
	baseUrl := basePath.JoinPath(tokenString)
	usersState := make(map[uint]state.UserState)
	bot := &Bot{
		client,
		token,
		baseUrl,
		10,
		0,
		usersState,
	}
	me, err := bot.getMe(ctx)
	if err != nil {
		return nil, e.Wrap(op, err)
	}
	slog.Info(op, slog.String("me", me.Username))
	return bot, nil
}

func (b *Bot) request(ctx context.Context, method string, body io.Reader) ([]byte, error) {
	op := "request"
	slog.Debug(op, slog.String("method", method))
	url := b.baseUrl.JoinPath(method)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		return nil, e.Wrap(op, err)
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := b.client.Do(req)
	if err != nil {
		return nil, e.Wrap(op, err)
	}
	defer res.Body.Close()
	result, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, e.Wrap(op, err)
	}
	return result, nil
}

func (b *Bot) producer(ctx context.Context, out chan *types.Update) {
	errCount := 0
	for {
		updates, err := b.getUpdates(ctx)
		if err != nil {
			slog.Error(err.Error())
			errCount += 1
			if errCount > 10 {
				slog.Error("error count greater than 10")
				break
			}
		} else if len(updates) != 0 {
			for _, update := range updates {
				out <- update
			}
		}
		time.Sleep(time.Second * 1)
	}
}

func (b *Bot) router(ctx context.Context, updatesChannel chan *types.Update) {
	for {
		select {
		case update := <-updatesChannel:
			if update.CallbackQuery != nil {
				err := b.callbackHandler(ctx, update)
				if err != nil {
					slog.Error(err.Error())
				}
			} else if slices.Contains(commands.List(), update.Message.Text) {
				err := b.commandHandler(ctx, update)
				if err != nil {
					slog.Error(err.Error())
				}
			} else if len(update.Message.Text) != 0 {
				err := b.messageHandler(ctx, update)
				if err != nil {
					slog.Error(err.Error())
				}
			}
		default:
			continue
		}
	}
}

func (b *Bot) echo(ctx context.Context, in chan *types.Update) {
	for u := range in {
		chatID := u.Message.Chat.ID
		text := u.Message.Text
		topButtons := []*types.KeyboardButton{
			{
				Text: "Button 1",
			},
			{
				Text: "Button 2",
			},
		}
		bottomButtons := []*types.KeyboardButton{
			{
				Text: "Button 3",
			},
			{
				Text: "Button 4",
			},
		}
		var kb [][]*types.KeyboardButton
		kb = append(kb, topButtons)
		kb = append(kb, bottomButtons)
		markup := types.ReplyKeyboardMarkup{
			Keyboard:       kb,
			ResizeKeyboard: true,
			Selective:      true,
		}
		message := types.SendMessageRequest{
			ChatID:      chatID,
			Text:        text,
			ReplyMarkup: markup,
		}
		err := b.sendMessage(ctx, &message)
		if err != nil {
			slog.Error(err.Error())
		}
	}
}

func (b *Bot) Run(ctx context.Context) {
	updatesChannel := make(chan *types.Update)
	go b.producer(ctx, updatesChannel)
	b.router(ctx, updatesChannel)
}
