package chat

type ChatRepository interface {
	ViewChat(sender string, recipient string) ([]Messages, error)
	SendMessage(chat Chat) error
}
