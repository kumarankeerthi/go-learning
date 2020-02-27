package chat

type chatService struct {
	chatRepositoy ChatRepository
}

func NewChatService(chatRepo ChatRepository) *chatService {
	return &chatService{chatRepositoy: chatRepo}
}

func (s *chatService) ViewChat(sender string, recipient string) ([]Chat, error) {
	return s.chatRepositoy.ViewChat(sender, recipient)
}

func (s *chatService) SendMessage(sender string, recipient string, message string) error {
	return s.chatRepositoy.SendMessage(sender, recipient, message)
}
