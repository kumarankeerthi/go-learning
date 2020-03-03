package chat

type ChatService interface {
	ViewChat(sender string, recipient string) ([]Messages, error)
	SendMessage(chat Chat) error
}
