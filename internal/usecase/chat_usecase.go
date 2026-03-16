package usecase

import "github.com/Adibayuluthfiansyah/Go-LiveChat/internal/domain"

type chatUsecase struct {
	chatRepo domain.ChatRepository
}

func NewChatUsecase(cr domain.ChatRepository) domain.ChatUseCase {
	return &chatUsecase{
		chatRepo: cr,
	}
}

// start stream
func (u *chatUsecase) StartStream(stream *domain.Stream) error {
	stream.IsLive = true
	return u.chatRepo.CreateStream(stream)
}

// end stream
func (u *chatUsecase) EndStream(streamID string) error {
	return u.chatRepo.UpdateStreamStatus(streamID, false)
}

// send message
func (u *chatUsecase) SendMessage(streamID string, userID string, content string) (*domain.Message, error) {
	msg := &domain.Message{
		StreamID: streamID,
		UserID:   userID,
		Content:  content,
	}
	err := u.chatRepo.SaveMessage(msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

// get chat history
func (u *chatUsecase) GetChatHistory(streamID string) ([]domain.Message, error) {
	return u.chatRepo.GetMessagesByStreamID(streamID, 60)
}

// get active stream
func (u *chatUsecase) GetActiveStreams() ([]domain.Stream, error) {
	return u.chatRepo.GetActiveStreams()
}
