package service

import (
	todo "github.com/darianfd99/todo-app/pkg"
	"github.com/darianfd99/todo-app/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{
		repo:     repo,
		listRepo: listRepo,
	}
}

func (s *TodoItemService) Create(userId int, listId int, item todo.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		//List does not exists or does not belongs to user
		return 0, err
	}

	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId int, listId int) ([]todo.TodoItem, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		//List does not exists or does not belongs to user
		return nil, err
	}

	return s.repo.GetAll(listId)
}
