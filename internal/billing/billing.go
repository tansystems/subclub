package billing

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/stripe/stripe-go/v78/webhook"
)

// Модуль Stripe: создание сессий оплаты, обработка webhook, проверка подписки.

// HandleStripeWebhook обрабатывает Stripe webhook
func HandleStripeWebhook(redisClient *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const MaxBodyBytes = int64(65536)
		r.Body = http.MaxBytesReader(w, r.Body, MaxBodyBytes)
		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		sigHeader := r.Header.Get("Stripe-Signature")
		webhookSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")
		event, err := webhook.ConstructEvent(payload, sigHeader, webhookSecret)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if event.Type == "checkout.session.completed" {
			var session struct {
				ID                string `json:"id"`
				CustomerEmail     string `json:"customer_email"`
				ClientReferenceID string `json:"client_reference_id"`
			}
			if err := json.Unmarshal(event.Data.Raw, &session); err == nil {
				ctx := context.Background()
				if session.ClientReferenceID != "" {
					redisClient.Set(ctx, "subscribed:"+session.ClientReferenceID, "1", 0)
				}
			}
		}
		w.WriteHeader(http.StatusOK)
	}
}

// IsSubscribed проверяет, есть ли активная подписка у пользователя
func IsSubscribed(redisClient *redis.Client, userID string) (bool, error) {
	ctx := context.Background()
	val, err := redisClient.Get(ctx, "subscribed:"+userID).Result()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return val == "1", nil
}
