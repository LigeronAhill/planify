package storage

import (
	"database/sql"
	"log/slog"
	"os/exec"
	"path"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

const (
	op = "база данных"
)

type Repository struct {
	db *sql.DB
}

func New(fileName string) (*Repository, error) {
	cmdOut, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return nil, err
	}
	projectPath := strings.TrimSpace(string(cmdOut))
	filePath := path.Join(projectPath, "storage", fileName)
	db, err := sql.Open("sqlite3", filePath)
	if err != nil {
		return nil, err
	}
	var version string
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)
	if err != nil {
		return nil, err
	}
	slog.Info(op, slog.String("version", version))
	return &Repository{db}, nil
}

func (r *Repository) Close() error {
	return r.db.Close()
}
