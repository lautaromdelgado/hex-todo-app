package ports

import "github.com/lautaromdelgado/hex-todo-app/internal/domain/model"

type ServiceTask interface {
	CreateTask(title string) (*model.Task, error)
	ListTask() ([]*model.Task, error)
	CompleteTask(id uint) error
}
