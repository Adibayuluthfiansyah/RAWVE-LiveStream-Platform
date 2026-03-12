package postgres

import (
	"github.com/Adibayuluthfiansyah/Go-LiveChat/internal/domain"
	"gorm.io/gorm"
)

type chatPG struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) domain.ChatRepository {
	return &chatPG{
		db: db,
	}
}

// create stream
func (r *chatPG) CreateStream(stream *domain.Stream) error {
	return r.db.Create(stream).Error
}

// update stream
func (r *chatPG) UpdateStreamStatus(streamID string, isLive bool) error {
	return r.db.Model(&domain.Stream{}).Where("id = ?", streamID).Update("is_live", isLive).Error
}

// save message
func (r *chatPG) SaveMessage(message *domain.Message) error {
	return r.db.Create(message).Error
}

// get recent message
func (r *chatPG) GetMessagesByStreamID(streamID string, limit int) ([]domain.Message, error) {
	var messages []domain.Message
	err := r.db.Where("stream_id = ?", streamID).Order("created_at desc").Limit(limit).Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}
