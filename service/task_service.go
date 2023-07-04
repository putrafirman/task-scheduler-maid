package service

import "context"

type taskService struct {
}

type TaskService interface {
	DeployVersion(ctx context.Context) error
}

func NewTaskService() TaskService {
	return &taskService{}
}

func (d *taskService) DeployVersion(ctx context.Context) error {
	return nil
}
