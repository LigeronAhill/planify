package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/LigeronAhill/planify/internal/telegram/commands"
	"github.com/LigeronAhill/planify/internal/telegram/state"
	"github.com/LigeronAhill/planify/internal/telegram/types"
)

func (b *Bot) commandHandler(ctx context.Context, update *types.Update) error {
	chatID := update.Message.Chat.ID
	messageID := update.Message.MessageID
	userId := update.Message.From.ID
	receivedCommand := update.Message.Text
	switch receivedCommand {
	case commands.Start:
		// TODO: add user to DB
		b.usersState[userId] = state.Common
		message := &types.SendMessageRequest{
			ChatID:      chatID,
			Text:        "Выберите пункт меню:",
			ReplyMarkup: makeKeyboard(),
		}
		return b.sendMessage(ctx, message)
	case commands.Help:
		message := &types.SendMessageRequest{
			ChatID:    update.Message.Chat.ID,
			Text:      `Справка: доступные команды и т\.д\.`,
			ParseMode: types.MardownV2,
		}
		return b.sendMessage(ctx, message)
	case commands.Cancel:
		return nil
	case commands.AddTask:
		err := b.deleteMessage(ctx, chatID, messageID)
		if err != nil {
			return err
		}
		b.usersState[userId] = state.AddingTask
		message := &types.SendMessageRequest{
			ChatID: update.Message.Chat.ID,
			Text:   "Введите название задачи",
		}
		return b.sendMessage(ctx, message)
	case commands.MyTasks:
		return nil
	case commands.DoneTasks:
		return nil
	case commands.Statistics:
		return nil
	default:
		message := &types.SendMessageRequest{
			ChatID:    update.Message.Chat.ID,
			Text:      fmt.Sprintf("Получена неизвестная команда: *%s*", strings.ReplaceAll(receivedCommand, "/", "")),
			ParseMode: types.MardownV2,
		}
		return b.sendMessage(ctx, message)
	}
}

func makeKeyboard() *types.ReplyKeyboardMarkup {
	topButtons := []*types.KeyboardButton{
		{
			Text: commands.AddTask,
		},
		{
			Text: commands.MyTasks,
		},
	}
	bottomButtons := []*types.KeyboardButton{
		{
			Text: commands.DoneTasks,
		},
		{
			Text: commands.Statistics,
		},
	}
	var kb [][]*types.KeyboardButton
	kb = append(kb, topButtons)
	kb = append(kb, bottomButtons)
	return &types.ReplyKeyboardMarkup{
		Keyboard:       kb,
		IsPersistent:   true,
		ResizeKeyboard: true,
		Selective:      true,
	}
}
