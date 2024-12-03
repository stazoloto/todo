package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	httplib "net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/stazoloto/todo/internal/delivery/http"
	"github.com/stazoloto/todo/internal/infrastructure/postgres"
	"github.com/stazoloto/todo/internal/usecase"
	"github.com/stazoloto/todo/pkg/logger"

	_ "github.com/lib/pq"
)

type App struct {
	db     *sql.DB
	logger *logger.Logger
}

func NewApp() *App {

	logger := logger.NewLogger()

	err := godotenv.Load()
	if err != nil {
		logger.Error(context.Background(), fmt.Sprintf("Error loading .env file: %v", err))
	}

	// Получение переменных окружения
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Строка подключения
	dbConnStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	// БД подключение
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return &App{
		db:     db,
		logger: logger,
	}
}

func (a *App) Run() error {
	todoRepo := postgres.NewTodoRepository(a.db)
	todoUsecase := usecase.NewTodoUsecase(todoRepo)
	TodoHandler := http.NewTodoHandler(todoUsecase)

	httplib.HandleFunc("/todos", func(w httplib.ResponseWriter, r *httplib.Request) {
		ctx := a.logger.NewContext(r.Context())
		TodoHandler.HandleTodos(w, r.WithContext(ctx))
	})

	fmt.Println("Server started at :9000")
	return httplib.ListenAndServe(":9000", nil)
}
