package thirdparty

type Warehouse interface {
	Robots() []Robot
}

type Robot interface {
	EnqueueTask(commands string) (taskID string, position chan RobotState, err chan error)

	CancelTask(taskID string) error

	CurrentState() RobotState

	// Additional methods needed
	QueryTask(taskID string) (Task, error)
}

type RobotState struct {
	X        uint
	Y        uint
	HasCrate bool
}

type TaskStatus string

const (
	TaskStatusUnknown   = TaskStatus("UNKNOWN")
	TaskStatusPending   = TaskStatus("PENDING")
	TaskStatusExecuting = TaskStatus("EXECUTING")
	TaskStatusDone      = TaskStatus("DONE")
)

type Task struct {
	Status TaskStatus
}

type Season int64

const (
	Summer Season = 0
	Autumn        = 1
	Winter        = 2
	Spring        = 3
)
