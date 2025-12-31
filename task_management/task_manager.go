package task_management

import (
	"CLI_Task_Manager/task"
	"encoding/json"
	"errors"
	"os"
)

type TaskManager struct {
	tasks    map[int]*task.Task
	nextID   int
	name     string
	filepath string
}

type TaskManagerDTO struct {
	Tasks    map[int]task.TaskDTO `json:"tasks"`
	NextID   int                  `json:"nextID"`
	Name     string               `json:"name"`
	Filepath string               `json:"path"`
}

func (tm *TaskManager) ToDTO() TaskManagerDTO {
	tasksDTO := make(map[int]task.TaskDTO)
	for id, t := range tm.tasks {
		tasksDTO[id] = t.ToDTO()
	}

	return TaskManagerDTO{
		Tasks:    tasksDTO,
		NextID:   tm.nextID,
		Name:     tm.name,
		Filepath: tm.filepath,
	}
}

func (tm *TaskManager) FromDTO(dto TaskManagerDTO) error {
	tm.tasks = make(map[int]*task.Task)

	for id, taskDTO := range dto.Tasks {
		t, err := task.NewTaskFromDTO(taskDTO)
		if err != nil {
			return err
		}
		tm.tasks[id] = t
	}

	tm.nextID = dto.NextID
	tm.name = dto.Name
	tm.filepath = dto.Filepath

	return nil
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

func (tm *TaskManager) SaveToFile() error {

	if tm.filepath == "" {
		return errors.New("Cannot Save To Empty FilePath...")
	}

	file, err := os.Create(tm.filepath)
	if err != nil {
		return err
	}

	defer file.Close()

	dto := tm.ToDTO()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(dto)

}

func (tm *TaskManager) SetPath(path string) {

	tm.filepath = path

}

func LoadFromFile(filepath string) (*TaskManager, error) {

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var dto TaskManagerDTO
	dto.Filepath = filepath
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&dto)
	if err != nil {
		return nil, err
	}

	tm := &TaskManager{}
	err = tm.FromDTO(dto)
	if err != nil {
		return nil, err
	}

	return tm, nil

}
