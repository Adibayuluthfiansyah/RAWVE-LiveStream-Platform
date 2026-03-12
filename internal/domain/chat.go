package domain

import "time"

type Stream struct {
	ID        string    `json:"id" gorm:"primaryKey;type:varchar(100)"`
	Title     string    `json:"title" gorm:"not null;default:'Untitled Stream'"`
	IsLive    bool      `json:"is_live" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type Message struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	StreamID  string    `json:"stream_id" gorm:"index;not null;type:varchar(100)"`
	UserID    string    `json:"user_id" gorm:"index;not null;type:varchar(100)"`
	Content   string    `json:"content" gorm:"not null;type:text"`
	CreatedAt time.Time `json:"created_at" gorm:"index;autoCreateTime"`
}

type ChatRepository interface {
	CreateStream(stream *Stream) error
	UpdateStreamStatus(streamID string, isLive bool) error
	SaveMessage(message *Message) error
	GetMessagesByStreamID(streamID string, limit int) ([]Message, error)
}

type ChatUseCase interface {
	StartStream(stream *Stream) error
	EndStream(streamID string) error
	SendMessage(streamID string, userID string, content string) (*Message, error)
	GetChatHistory(streamID string) ([]Message, error)
}
