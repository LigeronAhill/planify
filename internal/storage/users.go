package storage

import (
	"github.com/LigeronAhill/planify/internal/models"
)

func (r *Repository) InsertUser(user *models.User) error {
	existingUser, err := r.GetUser(user.ID)
	if err != nil {
		res := r.db.Create(user)
		if res.Error != nil {
			return res.Error
		}
	} else {
		user.CreatedAt = existingUser.CreatedAt
		result := r.db.Save(user)
		if result.Error != nil {
			return result.Error
		}

	}
	return nil
}

func (r *Repository) GetUser(id uint) (*models.User, error) {
	var user *models.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *Repository) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *Repository) DeleteUser(id uint) error {
	result := r.db.Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
