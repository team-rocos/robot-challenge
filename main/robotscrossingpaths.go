package main

import (
	"github.com/solimov-advanced/robot-challenge/simcli"
	"time"
)

/*
	Here, robot r1 crosses in front of robot r2 forcing r2 to wait.
*/

func main() {
	cli := simcli.NewSimulatorCli()
	cli.AddWarehouse("w1", 10, 10)
	cli.AddRobot("w1", "r1", 1, 4)
	cli.AddRobot("w1", "r2", 3, 2)
	cli.AddTask("w1", "r1", "EEEEE")
	cli.AddTask("w1", "r2", "NNNNN")

	for cli.IsRunning() {
		time.Sleep(time.Second * 1)
	}
}
