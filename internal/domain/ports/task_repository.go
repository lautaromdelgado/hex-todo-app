package ports

import "github.com/lautaromdelgado/hex-todo-app/internal/domain/model"

type RepositoryTask interface {
	Save(task *model.Task) error
	GetAll() ([]*model.Task, error)
	Complete(id uint) error
}
