package domain

import "time"

type User struct {
	ID          string    `json:"id" gorm:"primaryKey;type:varchar(100)"`
	Username    string    `json:"username" gorm:"uniqueIndex;not null;size:50"`
	DisplayName string    `json:"display_name" gorm:"size:100"`
	Email       string    `json:"email" gorm:"uniqueIndex;not null;size:100"`
	Bio         string    `json:"bio" gorm:"type:text"`
	Category    string    `json:"category" gorm:"size:50"`
	AvatarURL   string    `json:"avatar_url" gorm:"size:255"`
	SocialLinks string    `json:"social_links" gorm:"type:jsonb"`
	StreamKey   string    `json:"stream_key" gorm:"UniqueIndex;size:100"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type UserRepository interface {
	CreateOrUpdate(user *User) error
	GetByID(id string) (*User, error)
	UpdateProfile(user *User) error
}

type UserUseCase interface {
	SyncUserFromAuth(user *User) error
	UpdateProfile(userID string, displayName string, bio, category string) error
}
