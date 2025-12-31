package task_management

import "CLI_Task_Manager/task"

type taskManagerInterface interface {
	CreateTask(name string, description string, status int, done bool, id int) error
	ListTasks() []task.Task
	DeleteTask(id int) error
	GetTask(id int) *task.Task
	GetTaskManagerName() string
	GetNextID() int
	IncrementID()
}
