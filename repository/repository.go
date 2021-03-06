package repository

import (
	todo "github.com/darianfd99/todo-app/pkg"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId int, listId int) (todo.TodoList, error)
	Update(userId int, listId int, input todo.UpdateListInput) error
	Delete(userId int, listId int) error
}

type TodoItem interface {
	Create(listId int, item todo.TodoItem) (int, error)
	GetAll(listId int) ([]todo.TodoItem, error)
}

type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
