package task_management

import (
	"CLI_Task_Manager/task"
	"errors"
)

type TaskManager struct {
	tasks  map[int]*task.Task
	nextID int
	name   string
}

func MakeTaskManager(name string) (error, *TaskManager) {

	if name == "" {
		return errors.New("Cannot Make Task With No Name..."), nil
	}

	return nil, &TaskManager{name: name, tasks: make(map[int]*task.Task), nextID: 0}

}

func (tm *TaskManager) CreateTask(name string, description string, status int, done bool, id int) error {

	task, err := task.NewTask(id, name, description, status, done)

	if err != nil {
		return err
	} else if tm.tasks[id] != nil {
		return errors.New("Cannot Utilize The Same ID...")
	}

	tm.tasks[id] = task
	return nil

}

func (tm *TaskManager) ListTasks() []task.Task {

	tasks := make([]task.Task, 0, len(tm.tasks))
	for _, t := range tm.tasks {
		tasks = append(tasks, *t)
	}

	return tasks

}

func (tm *TaskManager) DeleteTask(id int) error {

	if tm.tasks[id] == nil {
		return errors.New("No Task With This ID...")
	}

	delete(tm.tasks, id)
	return nil

}

func (tm *TaskManager) GetTask(id int) *task.Task {

	return tm.tasks[id]

}

func (tm *TaskManager) GetTaskManagerName() string {

	return tm.name

}

func (tm *TaskManager) GetNextID() int {

	return tm.nextID

}

func (tm *TaskManager) IncrementID() {

	tm.nextID++

}
