package main

import (
	"github.com/solimov-advanced/robot-challenge/simcli"
	"time"
)

func main() {
	//testRobotCollision()
	testRobotMoveCrate()
}

func exampleCommandLineUsage() {
	cli := simcli.NewSimulatorCli()
	cli.Start()
	// enter commands in the commandline, type "quit" to end
	for true {
		time.Sleep(time.Second * 1)
		if !cli.IsRunning() {
			break
		}
	}
}

// Here, robot r1 crosses in front of robot r2 forcing r2 to wait.
func testRobotCollision() {
	cli := simcli.NewSimulatorCli()
	cli.Start()
	cli.AddWarehouse("w1", 10, 10)
	cli.AddRobot("w1", "r1", 1, 4)
	cli.AddRobot("w1", "r2", 3, 2)
	cli.AddTask("w1", "r1", "EEEEE")
	cli.AddTask("w1", "r2", "NNNNN")

	for true {
		time.Sleep(time.Second * 1)
		cli.QuitIfNoTasks()
		if !cli.IsRunning() {
			break
		}
	}
}

// Here, robot r1 moves to pick up a crate and moves to another location and places the crate.
func testRobotMoveCrate() {
	cli := simcli.NewSimulatorCli()
	cli.Start()
	cli.AddWarehouse("w1", 10, 10)
	cli.AddRobot("w1", "r1", 1, 4)
	cli.AddCrate("w1", 3, 5)
	cli.AddTask("w1", "r1", "EENG")
	cli.AddTask("w1", "r1", "EESED")

	for true {
		time.Sleep(time.Second * 1)
		cli.QuitIfNoTasks()
		if !cli.IsRunning() {
			break
		}
	}
}
