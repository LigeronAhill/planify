package repository

import (
	"context"
	"fmt"
	"log/slog"
	"path"
	"time"

	"github.com/LigeronAhill/planify/internal/e"
	"github.com/LigeronAhill/planify/internal/models"
)

func (r *Repository) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	op := "получение пользователей из базы данных"
	slog.Info(op)
	f := path.Join(r.queriesPath, "users", "get_all.sql")
	rows, err := r.QueryRows(ctx, f)
	if err != nil {
		return nil, e.Wrap(op, err)
	}
	defer rows.Close()
	var result []*models.User
	for rows.Next() {
		var id int
		var firstName, lastName, username string
		var created, updated time.Time
		err = rows.Scan(&id, &firstName, &lastName, &username, &created, &updated)
		if err != nil {
			return nil, e.Wrap(op, err)
		}
		user := &models.User{
			UserID:    id,
			FirstName: firstName,
			LastName:  lastName,
			Username:  username,
			Created:   created,
			Updated:   updated,
		}
		result = append(result, user)
	}
	return result, nil
}

func (r *Repository) InsertUser(ctx context.Context, user *models.User) (*models.User, error) {
	op := fmt.Sprintf("добавление или обновление пользователя в базе данных с id: %d", user.UserID)
	slog.Info(op)
	f := path.Join(r.queriesPath, "users", "insert.sql")
	row, err := r.QueryRow(
		ctx,
		f,
		user.UserID,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Created,
		user.Updated,
	)
	if err != nil {
		return nil, e.Wrap(op, err)
	}
	var id int
	var firstName, lastName, username string
	var created, updated time.Time
	err = row.Scan(&id, &firstName, &lastName, &username, &created, &updated)
	if err != nil {
		return nil, e.Wrap(op, err)
	}
	createdUser := &models.User{
		UserID:    id,
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Created:   created,
		Updated:   updated,
	}
	return createdUser, nil
}

func (r *Repository) GetUser(ctx context.Context, user_id int) (*models.User, error) {
	op := fmt.Sprintf("получение пользователя из базы данных с id: %d", user_id)
	slog.Info(op)
	f := path.Join(r.queriesPath, "users", "get.sql")
	row, err := r.QueryRow(ctx, f, user_id)
	if err != nil {
		return nil, e.Wrap(op, err)
	}
	var id int
	var firstName, lastName, username string
	var created, updated time.Time
	err = row.Scan(&id, &firstName, &lastName, &username, &created, &updated)
	if err != nil {
		return nil, e.Wrap(op, err)
	}
	createdUser := &models.User{
		UserID:    user_id,
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Created:   created,
		Updated:   updated,
	}
	return createdUser, nil
}

func (r *Repository) DeleteUser(ctx context.Context, user_id int) error {
	op := fmt.Sprintf("удаление пользователя из базы данных с id: %d", user_id)
	slog.Info(op)
	f := path.Join(r.queriesPath, "users", "delete.sql")
	err := r.Exec(ctx, f, user_id)
	if err != nil {
		return e.Wrap(op, err)
	}
	return nil
}
