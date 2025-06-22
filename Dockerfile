FROM golang:1.21-alpine as builder
WORKDIR /app
COPY . .
RUN go mod tidy && cd cmd/subclub && go build -o /subclub

FROM alpine:latest
WORKDIR /app
COPY --from=builder /subclub /subclub
COPY config/config.yaml ./config/config.yaml
COPY .env .env
EXPOSE 8080
CMD ["/subclub"]
