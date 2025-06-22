# SubClub

Сервис для приватного контента и автоматической выдачи доступа в Telegram.

## Стек
- Go
- Redis
- Stripe
- Telegram Bot API
- Docker

## Структура
- `cmd/subclub/` — основной сервис
- `internal/auth/` — авторизация
- `internal/billing/` — Stripe
- `internal/bot/` — Telegram Bot
- `internal/panel/` — панель автора
- `internal/storage/` — Redis
- `config/` — конфиги

## Быстрый старт
1. Скопируйте `.env.example` в `.env` и заполните переменные:
2. Запустите Redis и выполните:
   ```bash
   go run cmd/subclub/main.go
   ```
3. Для Docker:
   ```bash
   docker build -t subclub .
   docker run --env-file=.env -p 8080:8080 subclub
   ```

## Описание переменных .env

| Переменная              | Описание                                                                 | Пример значения                |
|-------------------------|--------------------------------------------------------------------------|---------------------------------|
| `REDIS_ADDR`            | Адрес Redis-сервера                                                      | `localhost:6379`                |
| `REDIS_PASS`            | Пароль Redis (если не требуется — оставить пустым)                       |                                 |
| `STRIPE_SECRET`         | Секретный ключ Stripe (API Key) — получить в [Stripe Dashboard](https://dashboard.stripe.com/apikeys) | `sk_test_...`                   |
| `STRIPE_WEBHOOK_SECRET` | Секрет Stripe Webhook — получить при создании webhook endpoint в Stripe   | `whsec_...`                     |
| `TELEGRAM_TOKEN`        | Токен Telegram-бота — получить у [@BotFather](https://t.me/BotFather)     | `123456:ABC-DEF...`             |
| `JWT_SECRET`            | Секрет для подписи JWT-токенов (любой длинный случайный текст)            | `supersecret`                   |
| `PORT`                  | Порт, на котором запускается сервис                                      | `8080`                          |

**Пример заполненного .env:**
```
REDIS_ADDR=localhost:6379
REDIS_PASS=
STRIPE_SECRET=sk_test_51N...yourkey...
STRIPE_WEBHOOK_SECRET=whsec_...yourwebhooksecret...
TELEGRAM_TOKEN=123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11
JWT_SECRET=supersecret
PORT=8080
```

## TODO
- Реализовать Stripe webhook
- Реализовать Telegram webhook
- Панель автора
- Авторизация пользователей
