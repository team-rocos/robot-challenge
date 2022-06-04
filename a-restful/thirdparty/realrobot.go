package thirdparty

import (
	"math/rand"
	"strconv"
)

type RealRobot struct {
	state RobotState
}

func NewRealRobot() *RealRobot {
	return &RealRobot{}
}

func (r *RealRobot) EnqueueTask(commands string) (
	taskID string, position chan RobotState, err chan error,
) {
	rand.Uint32()
	r.state = RobotState{
		X: uint(rand.Intn(9)),
		Y: uint(rand.Intn(9)),
	}

	taskID = strconv.Itoa(rand.Intn(100000))
	position = make(chan RobotState)
	err = make(chan error)

	return taskID, position, err
}

func (r *RealRobot) CancelTask(taskID string) error {
	return nil
}

func (r *RealRobot) CurrentState() RobotState {
	return r.state
}
