package thirdparty

type RealRobot struct {
	state RobotState
}

func NewRealRobot() *RealRobot {
	return &RealRobot{}
}

func (r *RealRobot) EnqueueTask(commands string) (
	taskID string, position chan RobotState, err chan error,
) {
	position = make(chan RobotState)
	err = make(chan error)
	return "task1", position, err
}

func (r *RealRobot) CancelTask(taskID string) error {
	return nil
}

func (r *RealRobot) CurrentState() RobotState {
	return r.state
}
