package usecase

import "github.com/Adibayuluthfiansyah/Go-LiveChat/internal/domain"

type UserUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(ur domain.UserRepository) domain.UserUseCase {
	return &UserUsecase{
		userRepo: ur,
	}
}

// sync from auth
func (u *UserUsecase) SyncUserFromAuth(user *domain.User) error {
	return u.userRepo.CreateOrUpdate(user)
}

// update profile
func (u *UserUsecase) UpdateProfile(userID string, displayName, bio, category string) error {
	user := &domain.User{
		ID:          userID,
		DisplayName: displayName,
		Bio:         bio,
		Category:    category,
	}
	return u.userRepo.UpdateProfile(user)
}
