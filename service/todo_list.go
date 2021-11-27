package service

import (
	todo "github.com/darianfd99/todo-app/pkg"
	"github.com/darianfd99/todo-app/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId int, ListId int) (todo.TodoList, error) {
	return s.repo.GetById(userId, ListId)
}

func (s *TodoListService) Delete(userId int, ListId int) error {
	return s.repo.Delete(userId, ListId)
}

func (s *TodoListService) Update(userId int, ListId int, input todo.UpdateListInput) error {
	return s.repo.Update(userId, ListId, input)
}
