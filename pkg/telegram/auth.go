package telegram

import (
	"context"
	"fmt"

	"github.com/Dumchez/telegram-bot-app/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) initAuthorizationProcess(message *tgbotapi.Message) error {
	authLink, err := b.generateAuthorizationLink(int(message.Chat.ID))
	if err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(b.messages.Start, authLink))

	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) getAccessToken(chatId int64) (string, error) {
	return b.tokenRepository.Get(repository.AccessTokens, chatId)

}

func (b *Bot) generateAuthorizationLink(chatId int) (string, error) {
	redirectURL := b.generateRedirectURL(chatId)
	reqToken, err := b.pocketClient.GetRequestToken(context.Background(), redirectURL)
	if err != nil {
		return "", err
	}

	if err := b.tokenRepository.Save(repository.RequestTokens, int64(chatId), reqToken); err != nil {
		return "", err
	}

	return b.pocketClient.GetAuthorizationURL(reqToken, redirectURL)
}

func (b *Bot) generateRedirectURL(chatId int) string {
	return fmt.Sprintf("%s?chat_id=%d", b.redirectURL, chatId)
}
