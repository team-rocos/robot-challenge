package simcli

import "github.com/solimov-advanced/robot-challenge/librobot"

// activeTask provides an abstraction of a robots task.
type activeTask struct {
	robotId  string
	taskId   string
	position chan librobot.RobotState
	err      chan error
}
