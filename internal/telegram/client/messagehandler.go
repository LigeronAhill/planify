package client

import (
	"context"
	"fmt"

	"github.com/LigeronAhill/planify/internal/enums"
	"github.com/LigeronAhill/planify/internal/telegram/state"
	"github.com/LigeronAhill/planify/internal/telegram/types"
)

func (b *Bot) messageHandler(ctx context.Context, update *types.Update) error {
	userID := update.Message.From.ID
	chatID := update.Message.Chat.ID
	if userState, ok := b.usersState[userID]; ok {
		switch userState {
		case state.Common:
			message := &types.SendMessageRequest{
				ChatID: chatID,
				Text:   "Пока я не могу обрабатывать текстовые сообщения",
			}
			return b.sendMessage(ctx, message)
		case state.AddingTask:
			// TODO: add task to DB
			b.usersState[userID] = state.Common
			taskName := update.Message.Text
			var keyboard [][]*types.InlineKeyboardButton
			priority := &types.InlineKeyboardButton{
				Text:         "Установить высокий приоритет",
				CallbackData: fmt.Sprintf("taskPriority high %s", taskName),
			}

			keyboard = append(keyboard, []*types.InlineKeyboardButton{priority})
			mu := &types.InlineKeyboardMarkup{
				InlineKeyboard: keyboard,
			}
			message := &types.SendMessageRequest{
				ChatID:      chatID,
				Text:        fmt.Sprintf("Название: '%s'\nПриоритет: %s", taskName, enums.NORMAL),
				ReplyMarkup: mu,
			}
			return b.sendMessage(ctx, message)
		default:
			message := &types.SendMessageRequest{
				ChatID: chatID,
				Text:   "Внутренняя ошибка",
			}
			return b.sendMessage(ctx, message)

		}
	}
	return nil
}
