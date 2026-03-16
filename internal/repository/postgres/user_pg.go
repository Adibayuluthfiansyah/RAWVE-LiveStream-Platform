package postgres

import (
	"github.com/Adibayuluthfiansyah/Go-LiveChat/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userPG struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userPG{
		db: db,
	}
}

// create or update
func (r *userPG) CreateOrUpdate(user *domain.User) error {
	return r.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(user).Error
}

// get by id clerk/oAuth
func (r *userPG) GetByID(id string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("id =?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// update profile
func (r *userPG) UpdateProfile(user *domain.User) error {
	return r.db.Model(&domain.User{}).Where("id = ?", user.ID).Updates(map[string]interface{}{
		"display_name": user.DisplayName,
		"bio":          user.Bio,
		"category":     user.Category,
	}).Error
}
