FROM golang:1.20.3-alpine3.17 AS builder

COPY . /github.com/Dumchez/telegram-bot-app/
WORKDIR /github.com/Dumchez/telegram-bot-app/

RUN go mod download
RUN go build -o ./bin/bot cmd/bot/main.go

FROM alpine:latest 

WORKDIR /root/

COPY --from=0 /github.com/Dumchez/telegram-bot-app/bin/bot .
COPY --from=0 /github.com/Dumchez/telegram-bot-app/configs configs/

EXPOSE 80

CMD [ "./bot" ]


