package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/stazoloto/todo/internal/domain"
	"github.com/stazoloto/todo/internal/usecase"
	"github.com/stazoloto/todo/pkg/logger"
)

type TodoHandler struct {
	usecase *usecase.TodoUsecase
}

func NewTodoHandler(usecase *usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{
		usecase: usecase,
	}
}

func (h *TodoHandler) HandleTodos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logger.FromContext(ctx)

	switch r.Method {
	case http.MethodGet:
		h.getAllTodos(w, r.WithContext(ctx))
	case http.MethodPost:
		h.createTodo(w, r.WithContext(ctx))
	case http.MethodPut:
		h.updateTodo(w, r.WithContext(ctx))
	case http.MethodDelete:
		h.deleteTodo(w, r.WithContext(ctx))
	default:
		logger.Error(ctx, "Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func (h *TodoHandler) getAllTodos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logger.FromContext(ctx)

	todos, err := h.usecase.GetAllTodos(ctx)
	if err != nil {
		logger.Error(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todos)
}


func (h *TodoHandler) createTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logger.FromContext(ctx)

	var todo domain.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		logger.Error(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := h.usecase.CreateTodo(ctx, todo)
	if err != nil {
		logger.Error(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	todo.ID = id
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) updateTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logger.FromContext(ctx)

	var todo domain.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		logger.Error(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := h.usecase.UpdateTodo(ctx, todo)
	if err != nil {
		logger.Error(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *TodoHandler) deleteTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logger.FromContext(ctx)

	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		logger.Error(ctx, "Invalid ID")
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	err = h.usecase.DeleteTodo(ctx, id)
	if err != nil {
		logger.Error(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

