package bot

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Модуль Telegram Bot: обработка команд, выдача доступа, webhooks.

// Bot содержит Telegram API и Redis
type Bot struct {
	API   *tgbotapi.BotAPI
	Redis *redis.Client
}

// InitBot инициализирует Telegram Bot API
func InitBot(redisClient *redis.Client) (*Bot, error) {
	token := os.Getenv("TELEGRAM_TOKEN")
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &Bot{API: api, Redis: redisClient}, nil
}

// HandleTelegramWebhook обрабатывает входящие webhook Telegram
func (b *Bot) HandleTelegramWebhook(w http.ResponseWriter, r *http.Request) {
	var update tgbotapi.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if update.Message != nil {
		userID := update.Message.From.ID
		// Проверяем подписку
		isSub, _ := b.Redis.Get(r.Context(), "subscribed:"+intToStr(userID)).Result()
		if isSub == "1" {
			b.GrantAccess(userID)
			b.API.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Доступ выдан!"))
		} else {
			b.API.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Сначала оплатите подписку!"))
		}
	}
	w.WriteHeader(http.StatusOK)
}

// GrantAccess — функция-заглушка для выдачи доступа (например, добавление в канал)
func (b *Bot) GrantAccess(userID int64) {
	// Здесь должна быть логика добавления пользователя в канал/группу через Telegram API
	log.Printf("Выдаю доступ пользователю %d", userID)
}

func intToStr(id int64) string {
	return fmt.Sprintf("%d", id)
}
