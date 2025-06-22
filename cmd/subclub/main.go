package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v78"

	// "github.com/stripe/stripe-go/v78/webhook"
	// "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"subclub/internal/auth"
	"subclub/internal/billing"
	"subclub/internal/bot"
	"subclub/internal/panel"
)

func main() {
	// Загрузка .env
	_ = godotenv.Load()

	// Инициализация Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})
	defer redisClient.Close()

	// Инициализация Stripe
	stripe.Key = os.Getenv("STRIPE_SECRET")

	// Инициализация Telegram Bot
	telegramBot, err := bot.InitBot(redisClient)
	if err != nil {
		log.Fatalf("Ошибка Telegram Bot: %v", err)
	}

	// Web сервер и маршруты
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "SubClub работает!")
	})

	r.HandleFunc("/stripe/webhook", billing.HandleStripeWebhook(redisClient))
	r.HandleFunc("/telegram/webhook", telegramBot.HandleTelegramWebhook)
	r.HandleFunc("/register", auth.RegisterHandler(redisClient)).Methods("POST")
	r.HandleFunc("/login", auth.LoginHandler(redisClient)).Methods("POST")
	r.HandleFunc("/panel", panel.AuthorPanelHandler)
	// TODO: добавить обработчики: Stripe webhook, Telegram webhook, панель автора

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Сервер запущен на :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
