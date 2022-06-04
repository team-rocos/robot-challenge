package service

import (
	"github.com/edwardkcyu/robot-challenge/a-restful/internal/util"
	"github.com/edwardkcyu/robot-challenge/a-restful/thirdparty"
)

type IRobotService interface {
	EnqueueTask(command string) error
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

func (s *RobotService) EnqueueTask(command string) error {
	s.log.Info("run command")
	return nil
}
