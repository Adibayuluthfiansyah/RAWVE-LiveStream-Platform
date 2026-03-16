package domain

import "time"

type Stream struct {
	ID                string    `json:"id" gorm:"primaryKey;type:varchar(100)"`
	Title             string    `json:"title" gorm:"not null;default:'Untitled Stream'"`
	Category          string    `json:"category" gorm:"size:100"`
	ThumbnailURL      string    `json:"thumbnail_url" gorm:"size:255"`
	IsLive            bool      `json:"is_live" gorm:"default:false"`
	EnableDonation    bool      `json:"enable_donation" gorm:"default:true"`
	FollowersOnlyChat bool      `json:"followers_only_chat" gorm:"default:false"`
	CreatedAt         time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt         time.Time `json:"updated_at" gorm:"autoUpdateTime"`
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
	GetActiveStreams() ([]Stream, error)
}

type ChatUseCase interface {
	StartStream(stream *Stream) error
	EndStream(streamID string) error
	SendMessage(streamID string, userID string, content string) (*Message, error)
	GetChatHistory(streamID string) ([]Message, error)
	GetActiveStreams() ([]Stream, error)
}
