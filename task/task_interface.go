package task

type TaskInterface interface {
	GetStatus() int
	SetStatus(state int) error
	GetName() string
	SetName(name string) error
	GetDescription() string
	SetDescription(desc string) error
	IsDone() bool
	SetDone(done bool) error
	GetID() int
	setID(id int) error
}
