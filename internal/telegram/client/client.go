package client

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/LigeronAhill/planify/internal/e"
	"github.com/LigeronAhill/planify/internal/telegram/types"
)

// https://api.telegram.org/bot<token>/METHOD_NAME
const BASE_PATH = "https://api.telegram.org"

type Bot struct {
	client  *http.Client
	token   string
	baseUrl *url.URL
	limit   int
	offset  int
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
	bot := &Bot{
		client,
		token,
		baseUrl,
		10,
		0,
	}
	me, err := bot.getMe(ctx)
	if err != nil {
		return nil, e.Wrap(op, err)
	}
	slog.Info(op, slog.String("me", me.Username))
	return bot, nil
}

func (b *Bot) request(ctx context.Context, method string, body io.Reader) ([]byte, error) {
	op := fmt.Sprintf("%s request", method)
	slog.Debug(op)
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

func (b *Bot) echo(ctx context.Context, in chan *types.Update) {
	for u := range in {
		chatID := u.Message.Chat.ID
		text := u.Message.Text
		err := b.sendTextMessage(ctx, chatID, text)
		if err != nil {
			slog.Error(err.Error())
		}
	}
}

func (b *Bot) Run(ctx context.Context) {
	updatesChannel := make(chan *types.Update)
	go b.producer(ctx, updatesChannel)
	b.echo(ctx, updatesChannel)
}
