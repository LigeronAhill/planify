package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/LigeronAhill/planify/internal/e"
	"github.com/LigeronAhill/planify/internal/enums"
	"github.com/LigeronAhill/planify/internal/telegram/types"
)

func (b *Bot) callbackHandler(ctx context.Context, update *types.Update) error {
	op := "answer callback query"
	data := update.CallbackQuery.Data
	queryID := update.CallbackQuery.ID
	if strings.HasPrefix(data, "taskPriority") {
		priorityString := strings.Split(data, " ")[1]
		taskName := strings.Split(data, " ")[2]
		priority := enums.NORMAL
		priorityButton := &types.InlineKeyboardButton{}
		switch priorityString {
		case "high":
			priority = enums.HIGH
			priorityButton = &types.InlineKeyboardButton{
				Text:         "Установить обычный приоритет",
				CallbackData: fmt.Sprintf("taskPriority normal %s", taskName),
			}
		case "normal":
			priority = enums.NORMAL
			priorityButton = &types.InlineKeyboardButton{
				Text:         "Установить высокий приоритет",
				CallbackData: fmt.Sprintf("taskPriority high %s", taskName),
			}
		}
		chatID := update.CallbackQuery.Message.Chat.ID
		messageID := update.CallbackQuery.Message.MessageID
		var keyboard [][]*types.InlineKeyboardButton

		keyboard = append(keyboard, []*types.InlineKeyboardButton{priorityButton})
		mu := &types.InlineKeyboardMarkup{
			InlineKeyboard: keyboard,
		}
		message := &types.EditMessageTextRequest{
			ChatID:      chatID,
			MessageID:   messageID,
			Text:        fmt.Sprintf("Название: '%s'\nПриоритет: %s", taskName, priority),
			ReplyMarkup: mu,
		}
		err := b.EditMessageText(ctx, message)
		if err != nil {
			return e.Wrap(op, err)
		}
		return b.answerCallbackQuery(ctx, queryID, fmt.Sprintf("Устанавливаем приоритет %s для задачи %s", priority, taskName))
	}
	return nil
}
