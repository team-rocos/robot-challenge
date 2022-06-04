package mock

import "github.com/edwardkcyu/robot-challenge/a-restful/thirdparty"

type MockRobot struct {
	state thirdparty.RobotState
}

func NewMockRobot() *MockRobot {
	return &MockRobot{}
}

func (r *MockRobot) EnqueueTask(commands string) (
	taskID string, position chan thirdparty.RobotState, err chan error,
) {
	position = make(chan thirdparty.RobotState)
	err = make(chan error)
	return "task1", position, err
}

func (r *MockRobot) CancelTask(taskID string) error {
	return nil
}

func (r *MockRobot) CurrentState() thirdparty.RobotState {
	return r.state
}
