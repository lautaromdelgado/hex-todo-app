package aplication

import (
	"time"

	"github.com/lautaromdelgado/hex-todo-app/internal/domain/model"
	"github.com/lautaromdelgado/hex-todo-app/internal/domain/ports"
)

type TaskServiceImp struct {
	repo ports.RepositoryTask
}

func NewTaskService(repositoryTask ports.RepositoryTask) *TaskServiceImp {
	return &TaskServiceImp{repo: repositoryTask}
}

func (t *TaskServiceImp) CreateTask(title string) (*model.Task, error) {
	task := &model.Task{
		Task:      "Lavar los platos",
		Completed: false,
		CreatedIn: time.Now(),
	}
	err := t.repo.Save(task)
	return task, err
}

func (t *TaskServiceImp) ListTask() ([]*model.Task, error) {
	return t.repo.GetAll()
}

func (t *TaskServiceImp) CompleteTask(id uint) error {
	return t.repo.Complete(id)
}
