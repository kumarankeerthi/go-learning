package chat

type ChatService interface {
	ViewChat(sender string, recipient string) ([]Chat, error)
	SendMessage(sender string, recipient string, message string) error
}
