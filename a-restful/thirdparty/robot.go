package thirdparty

type Warehouse interface {
	Robots() []Robot
}

type Robot interface {
	EnqueueTask(commands string) (taskID string, position chan RobotState, err chan error)

	CancelTask(taskID string) error

	CurrentState() RobotState
}

type RobotState struct {
	X        uint
	Y        uint
	HasCrate bool
}
