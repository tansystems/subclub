SubClub

SubClub is a backend service for delivering private content with automatic Telegram access management. Designed for creators who monetize via subscriptions.

Stack

Go
Redis
Stripe
Telegram Bot API
Docker
Structure

cmd/subclub/ — main service
internal/auth/ — user authentication
internal/billing/ — Stripe integration
internal/bot/ — Telegram bot logic
internal/panel/ — author dashboard
internal/storage/ — Redis database
config/ — environment configs

Quick Start
1) Copy .env.example to .env and fill in the variables
2) Start Redis and run:
go run cmd/subclub/main.go
Or using Docker:
docker build -t subclub .
docker run --env-file=.env -p 8080:8080 subclub

.env Configuration
| Variable                | Description                                          | Example            |
| ----------------------- | ---------------------------------------------------- | ------------------ |
| REDIS\_ADDR             | Redis server address                                 | localhost:6379     |
| REDIS\_PASS             | Redis password (leave blank if not used)             |                    |
| STRIPE\_SECRET          | Stripe API secret key (get from Stripe Dashboard)    | sk\_test\_...      |
| STRIPE\_WEBHOOK\_SECRET | Stripe Webhook secret (get when setting up endpoint) | whsec\_...         |
| TELEGRAM\_TOKEN         | Telegram bot token (create via @BotFather)           | 123456\:ABC-DEF... |
| JWT\_SECRET             | Secret key for signing JWT tokens                    | supersecret        |
| PORT                    | Port to run the service                              | 8080               |

Example:

REDIS_ADDR=localhost:6379

REDIS_PASS=

STRIPE_SECRET=sk_test_...

STRIPE_WEBHOOK_SECRET=whsec_...

TELEGRAM_TOKEN=123456:ABC-DEF...

JWT_SECRET=supersecret

PORT=8080

TODO

Stripe webhook handling
Telegram webhook implementation
Author panel UI
User authentication flow
