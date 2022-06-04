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
	TaskStatusUnknown   TaskStatus = "UNKNOWN"
	TaskStatusPending              = "PENDING"
	TaskStatusExecuting            = "EXECUTING"
	TaskStatusDone                 = "DONE"
)

type Task struct {
	Status TaskStatus
}
