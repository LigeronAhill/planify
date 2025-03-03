package storage

import (
	"os/exec"
	"path"
	"strings"

	"github.com/LigeronAhill/planify/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(fileName string) (*Repository, error) {
	cmdOut, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return nil, err
	}
	projectPath := strings.TrimSpace(string(cmdOut))

	filePath := path.Join(projectPath, "storage", fileName)
	db, err := gorm.Open(sqlite.Open(filePath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repository{db}, nil
}

func (r *Repository) Migrate() error {
	err := r.db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}
	err = r.db.AutoMigrate(&models.Task{})
	if err != nil {
		return err
	}
	return nil
}
