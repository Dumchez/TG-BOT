package repository

type Bucket string

const (
	AccessTokens  Bucket = "access_tokens"
	RequestTokens Bucket = "request_tokens"
)

type TokenRepository interface {
	Save(bucket Bucket, chatId int64, token string) error
	Get(bucket Bucket, chatId int64) (string, error)
}
