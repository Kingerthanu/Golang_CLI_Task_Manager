package task

import "errors"

type Task struct {
	status      int
	name        string
	description string
	done        bool
	id          int
}

func (task *Task) GetStatus() int {

	return task.status

}

func (task *Task) SetStatus(state int) error {

	if state < -1 {
		return errors.New("Invalid State ([-1, INT_MAX])...")
	}

	task.status = state
	return nil

}

func (task *Task) GetName() string {

	return task.name

}

func (task *Task) SetName(name string) error {

	if name == "" {
		return errors.New("Empty Name Is Invalid...")
	}

	task.name = name
	return nil

}

func (task *Task) GetDescription() string {

	return task.description

}

func (task *Task) SetDescription(desc string) error {

	if desc == "" {
		return errors.New("Empty Description Is Invalid...")
	}

	task.description = desc
	return nil

}

func (task *Task) IsDone() bool {

	return task.done

}

func (task *Task) SetDone(done bool) error {

	task.done = done
	return nil

}

func (task *Task) GetID() int {

	return task.id

}

func (task *Task) setID(id int) error {

	if id < 0 {
		return errors.New("Cannot Set ID To Negative...")
	}

	task.id = id
	return nil

}

func NewTask(id int, name, description string, status int, done bool) (*Task, error) {

	var task Task

	if err := task.SetName(name); err != nil {
		return nil, err
	}

	if err := task.SetDescription(description); err != nil {
		return nil, err
	}

	if err := task.SetStatus(status); err != nil {
		return nil, err
	}

	if err := task.SetDone(done); err != nil {
		return nil, err
	}

	if err := task.setID(id); err != nil {
		return nil, err
	}

	return &task, nil

}
