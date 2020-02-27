package chat

type ChatRepository interface {
	ViewChat(sender string, recipient string) ([]Chat, error)
	SendMessage(sender string, recipient string, message string) error
}
