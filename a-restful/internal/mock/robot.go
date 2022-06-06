package mock

import (
	"errors"

	"github.com/edwardkcyu/robot-challenge/a-restful/thirdparty"
)

type MockRobot struct {
	state               thirdparty.RobotState
	hasEnqueueTaskError bool
	hasCancelTaskError  bool
}

func NewMockRobot(x uint, y uint) *MockRobot {
	return &MockRobot{
		state: thirdparty.RobotState{
			X: x,
			Y: y,
		},
	}
}

func (r *MockRobot) WithHasEnqueueTaskError(hasEnqueueTaskError bool) *MockRobot {
	r.hasEnqueueTaskError = hasEnqueueTaskError
	return r
}

func (r *MockRobot) WithHasCancelTaskError(hasCancelTaskError bool) *MockRobot {
	r.hasCancelTaskError = hasCancelTaskError
	return r
}

func (r *MockRobot) EnqueueTask(commands string) (
	taskID string, position chan thirdparty.RobotState, err chan error,
) {
	position = make(chan thirdparty.RobotState)
	err = make(chan error)

	if r.hasEnqueueTaskError {
		go func() {
			position <- thirdparty.RobotState{}
		}()

		go func() {
			err <- errors.New("enqueue task failed")
		}()

		return "", position, err

	}

	go func() {
		position <- thirdparty.RobotState{
			X:        1,
			Y:        2,
			HasCrate: false,
		}
	}()

	go func() {
		err <- nil
	}()

	return "task1", position, err
}

func (r *MockRobot) CancelTask(taskID string) error {
	if r.hasCancelTaskError {
		return errors.New("cancel task failed")
	}
	return nil
}

func (r *MockRobot) CurrentState() thirdparty.RobotState {
	return r.state
}

func (r *MockRobot) QueryTask(taskId string) (thirdparty.Task, error) {
	task := thirdparty.Task{
		Status: thirdparty.TaskStatusExecuting,
	}
	return task, nil
}
