package service

import (
	"github.com/Eagoker/todo-list"
	"github.com/Eagoker/todo-list/pkg/repository"
)

type Authorization interface{
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)

}

type TodoList interface{
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
}

type TodoItem interface{

}

type Service struct{
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		TodoList: NewTodoListService(repo.TodoList),
	}
}