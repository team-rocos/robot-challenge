package service

import (
	"fmt"
	"strings"

	"github.com/edwardkcyu/robot-challenge/a-restful/internal/util"
	"github.com/edwardkcyu/robot-challenge/a-restful/thirdparty"
)

const (
	MaxEast  = 9
	MinEast  = 0
	MaxNorth = 9
	MinNorth = 0
)

const (
	North string = "N"
	East         = "E"
	South        = "S"
	West         = "W"
)

type IRobotService interface {
	EnqueueTask(command string) (string, error)
	ValidateMovement(command string) error
	CancelTask(taskID string) error
}

type RobotService struct {
	log   *util.Logger
	robot thirdparty.Robot
}

func NewRobotService(robot thirdparty.Robot) *RobotService {
	return &RobotService{
		log:   util.NewLogger("RobotService"),
		robot: robot,
	}
}

func (s *RobotService) EnqueueTask(command string) (string, error) {
	taskID, _, _ := s.robot.EnqueueTask(command)

	return taskID, nil
}

func (s *RobotService) ValidateMovement(command string) error {
	position := s.robot.CurrentState()
	moves := strings.Split(command, " ")

	x := position.X
	y := position.Y
	for _, move := range moves {
		switch move {
		case North:
			y = y + 1
		case East:
			x = x + 1
		case South:
			y = y - 1
		case West:
			x = x - 1
		}
	}

	if y > MaxNorth {
		return fmt.Errorf("the move is out of the north boundary: (%d, %d)", x, y)
	}

	if y < MinNorth {
		return fmt.Errorf("the move is out of the south boundary: (%d, %d)", x, y)
	}

	if x > MaxEast {
		return fmt.Errorf("the move is out of the east bound: (%d, %d)", x, y)
	}

	if x < MinEast {
		return fmt.Errorf("the move is out of the west bound: (%d, %d)", x, y)
	}

	return nil
}

func (s *RobotService) CancelTask(taskID string) error {
	err := s.robot.CancelTask(taskID)
	if err != nil {
		return fmt.Errorf("robot failed to cancel task: %w", err)
	}

	return nil
}
