package storage

import (
	"context"
	"database/sql"
	"log/slog"
	"os"
	"os/exec"
	"path"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

const (
	op = "база данных"
)

type Repository struct {
	db          *sql.DB
	queriesPath string
}

func New(ctx context.Context, fileName string) (*Repository, error) {
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

	queriesPath := path.Join(projectPath, "storage", "queries")

	repository := &Repository{db, queriesPath}

	var version string
	row, err := repository.QueryRow(ctx, "version.sql")
	if err != nil {
		return nil, err
	}
	err = row.Scan(&version)
	if err != nil {
		return nil, err
	}
	slog.Info(op, slog.String("version", version))
	return repository, nil
}

func (r *Repository) Close() error {
	return r.db.Close()
}

func (r *Repository) QueryRow(ctx context.Context, fileName string, args ...any) (*sql.Row, error) {
	qPath := path.Join(r.queriesPath, fileName)
	qb, err := os.ReadFile(qPath)
	if err != nil {
		return nil, err
	}
	q := string(qb)
	result := r.db.QueryRowContext(ctx, q, args...)
	return result, nil
}

func (r *Repository) Query(ctx context.Context, fileName string, args ...any) (*sql.Rows, error) {
	qPath := path.Join(r.queriesPath, fileName)
	qb, err := os.ReadFile(qPath)
	if err != nil {
		return nil, err
	}
	q := string(qb)
	result, err := r.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Repository) Exec(ctx context.Context, fileName string, args ...any) (*sql.Rows, error) {
	qPath := path.Join(r.queriesPath, fileName)
	qb, err := os.ReadFile(qPath)
	if err != nil {
		return nil, err
	}
	q := string(qb)
	result, err := r.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}
