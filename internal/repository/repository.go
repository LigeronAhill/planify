package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/LigeronAhill/planify/internal/e"
	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	db             *sql.DB
	migrationsPath string
	queriesPath    string
}

func New(ctx context.Context, fileName string) (*Repository, error) {
	op := "соединение с базой данных"
	cmdOut, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return nil, e.Wrap(op, err)
	}
	projectPath := strings.TrimSpace(string(cmdOut))

	filePath := path.Join(projectPath, "storage", fileName)
	migrationsPath := path.Join(projectPath, "storage", "migrations")
	queriesPath := path.Join(projectPath, "storage", "queries")
	db, err := sql.Open("sqlite3", filePath)
	if err != nil {
		return nil, e.Wrap(op, err)
	}
	var version string
	err = db.QueryRowContext(ctx, "SELECT SQLITE_VERSION()").Scan(&version)
	if err != nil {
		return nil, e.Wrap(op, err)
	}
	slog.Info(op, slog.String("version", version))

	return &Repository{db, migrationsPath, queriesPath}, nil
}

func (r *Repository) Close() error {
	return r.db.Close()
}

func (r *Repository) Migrate(ctx context.Context) error {
	op := "миграция базы данных"
	slog.Info(op)
	files, err := getFileNames(r.migrationsPath)
	if err != nil {
		return e.Wrap(op, err)
	}
	for _, file := range files {
		query, err := getQuery(file)
		if err != nil {
			return e.Wrap(op, err)
		}
		stmt, err := r.db.PrepareContext(ctx, query)
		if err != nil {
			return e.Wrap(op, err)
		}
		_, err = stmt.ExecContext(ctx)
		if err != nil {
			return e.Wrap(op, err)
		}
	}
	return nil
}

func getFileNames(dir string) ([]string, error) {
	op := "получение списка файлов"
	var files []string
	entries, err := os.ReadDir(dir)
	if err != nil {
		return files, e.Wrap(op, err)
	}
	for _, e := range entries {
		files = append(files, e.Name())
	}
	sort.Slice(files, func(i, j int) bool {
		name1 := files[i]
		name2 := files[j]
		strNum1 := strings.Split(name1, "_")[0]
		strNum2 := strings.Split(name2, "_")[0]
		num1, err := strconv.Atoi(strNum1)
		if err != nil {
			return false
		}
		num2, err := strconv.Atoi(strNum2)
		if err != nil {
			return false
		}
		return num1 < num2
	})
	for i := 0; i < len(files); i++ {
		files[i] = path.Join(dir, files[i])
	}
	return files, nil
}

func getQuery(file string) (string, error) {
	op := fmt.Sprintf("чтение запроса из файла '%s'", file)
	data, err := os.ReadFile(file)
	if err != nil {
		return "", e.Wrap(op, err)
	}
	return string(data), nil
}

func (r *Repository) QueryRows(ctx context.Context, file string, args ...any) (*sql.Rows, error) {
	op := fmt.Sprintf("запрос в базе данных из файла '%s'", file)
	query, err := getQuery(file)
	if err != nil {
		return nil, e.Wrap(op, err)
	}
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, e.Wrap(op, err)
	}
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, e.Wrap(op, err)
	}
	return rows, nil
}

func (r *Repository) QueryRow(ctx context.Context, file string, args ...any) (*sql.Row, error) {
	op := fmt.Sprintf("запрос одной строки в базе данных из файла '%s'", file)
	query, err := getQuery(file)
	if err != nil {
		return nil, e.Wrap(op, err)
	}
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, e.Wrap(op, err)
	}
	row := stmt.QueryRowContext(ctx, args...)
	return row, nil
}

func (r *Repository) Exec(ctx context.Context, file string, args ...any) error {
	op := fmt.Sprintf("запрос без ответа в базе данных из файла '%s'", file)
	query, err := getQuery(file)
	if err != nil {
		return e.Wrap(op, err)
	}
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return e.Wrap(op, err)
	}
	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		return e.Wrap(op, err)
	}
	return nil
}
