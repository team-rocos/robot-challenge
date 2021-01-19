package robot

import (
	"github.com/solimov-advanced/robot-challenge/librobot"
	"github.com/solimov-advanced/robot-challenge/robot/commanding"
	"github.com/solimov-advanced/robot-challenge/robot/tasking"
	"strconv"
	"time"
)

// Mode type is used by the robots operational finite-mode-machine.
type Mode uint

const (
	Off Mode = iota
	AwaitingTask
	GettingCommand
	RunningCommand
)

// Robot provides abstraction of a simulated robot.
type Robot struct {
	model            Model
	name             string
	state            librobot.RobotState
	nextState        librobot.RobotState
	environment      librobot.Warehouse
	mode             Mode
	tasks            tasking.TaskQueue
	tasker           tasking.TaskProcessor
	commander        commanding.Command
	isRunning        bool
	turnRobotOffFlag bool
	turnRobotOnFlag  bool
	position         chan librobot.RobotState
	err              chan error
}

// NewRobot returns a Robot with specified name, model and position, requires a reference to its environment.
func NewRobot(name string, x, y uint, model Model, env librobot.Warehouse) Robot {
	newRobot := Robot{
		model:       model,
		name:        name,
		state:       librobot.RobotState{X: x, Y: y},
		nextState:   librobot.RobotState{X: x, Y: y},
		environment: env,
		mode:        Off,
		tasks:       tasking.TaskQueue{},
		tasker:      tasking.TaskProcessor{},
		commander:   nil,
		position:    make(chan librobot.RobotState),
		err:         make(chan error),
	}
	return newRobot
}

// Start begins the loop go routine.
func (r *Robot) Start() {
	r.turnRobotOnFlag = true
	if !r.isRunning {
		go r.loop()
	}
}

// Stop terminates the loop go routine that operates the robot, waits for the last command to finish.
func (r *Robot) Stop() {
	r.turnRobotOffFlag = false
}

// loop go routine that loops and operates the robot with a finite-mode-machine.
func (r *Robot) loop() {
	r.isRunning = true
	for r.mode != Off || r.turnRobotOnFlag {
		r.updateMode()
		r.modeLogic()
	}
	r.isRunning = false
}

// updateMode is called by loop to check for mode transitions, mode is only changed here.
func (r *Robot) updateMode() {
	switch r.mode {
	case Off:
		if r.turnRobotOnFlag {
			r.mode = AwaitingTask
			r.turnRobotOnFlag = false
			println("---- Robot '" + r.name + "' was turned ON")
		}
		return
	case AwaitingTask:
		if r.turnRobotOffFlag {
			r.mode = Off
			r.turnRobotOffFlag = false
			println("---- Robot '" + r.name + "' was turned OFF")
		} else if !r.tasker.IsEndReached() {
			r.mode = GettingCommand
		}
		return
	case GettingCommand:
		if r.turnRobotOffFlag {
			r.mode = Off
			r.turnRobotOffFlag = false
			println("---- Robot '" + r.name + "' was turned OFF")
		} else if r.commander == nil && r.tasker.IsEndReached() {
			r.mode = AwaitingTask
			println("---- Robot '" + r.name + "' completed a Task")
		} else if r.commander != nil && !r.commander.HasStarted() {
			r.mode = RunningCommand
		}
		return
	case RunningCommand:
		if r.commander == nil {
			r.mode = GettingCommand
		}
		return
	}
}

// modeLogic is called by loop to affect logic according to mode, mode is never changed here.
func (r *Robot) modeLogic() {
	switch r.mode {
	case Off:
		// do nothing
		return
	case AwaitingTask:
		if r.tasks.Count() > 0 {
			nextTask := r.tasks.GetNextTask()
			if !nextTask.IsEmpty() {
				r.tasker.NewActiveTask(nextTask)
				println("---- Robot '"+r.name+"' started Task[", nextTask.Commands, "]")
			}
		}
		return
	case GettingCommand:
		cmd := r.tasker.GetNextCommand(getRobotCommands(r.model))
		if len(cmd) > 0 {
			r.commander = commanding.ParseCommand(cmd)
		}
		return
	case RunningCommand:
		if !r.commander.HasStarted() {
			finalState := r.commander.FinalState(r.state)

			if r.state.X != finalState.X || r.state.Y != finalState.Y {
				maxX, maxY := r.environment.Size()
				if finalState.X >= maxX && finalState.Y >= maxY {
					println("---- Robot '" + r.name + "' aborted out-of-bounds move command '" + r.commander.GetCommand() + "'")
					r.commander = nil
					return
				}

				for _, bot := range r.environment.Robots() {
					state := bot.State()
					nextState := bot.NextState()
					if r.name == bot.Name() {
						continue
					} else if (finalState.X == state.X && finalState.Y == state.Y) ||
						(finalState.X == nextState.X && finalState.Y == nextState.Y) {
						println("---- Robot '" + r.name + "' waiting for other robot to move...")
						time.Sleep(300 * time.Millisecond) // wait is only used to stop message flooding console
						return
					}
				}
			}

			r.nextState = finalState
			r.commander.Start()
			println("---- Robot '" + r.name + "' began command '" + r.commander.GetCommand() + "' ") //[" + r.tasker.DebugString() + "]")
			return
		} else if r.commander.HasCompleted() {

			if r.state.HasCrate != r.nextState.HasCrate {
				temp, ok := r.environment.(librobot.CrateWarehouse)
				if ok {
					switch r.nextState.HasCrate {
					case true:
						err := temp.DelCrate(r.state.X, r.state.Y)
						if err != nil {
							println("---- Robot '" + r.name + "' failed to grab crate")
							r.commander = nil
							return
						}
						break
					case false:
						err := temp.AddCrate(r.state.X, r.state.Y)
						if err != nil {
							println("---- Robot '" + r.name + "' failed to drop crate")
							r.commander = nil
							return
						}
					}
				}
			}

			r.state = r.nextState
			print("---- Robot '" + r.name + "' finished command '" + r.commander.GetCommand() + "' ")
			print("new Mode: (X: ", r.state.X, ", Y: ", r.state.Y, ")")
			if r.state.HasCrate {
				println(" HasCrate: T")
			} else {
				println(" HasCrate: F")
			}
			r.commander = nil
			return
		}
		return
	}
}

// EnqueueTask appends a new task to the robots task queue, returns the tasks id.
func (r *Robot) EnqueueTask(commands string) (id string, position chan librobot.RobotState, err chan error) {
	id = r.tasks.AddTask(commands)
	position = r.position
	err = r.err
	return
}

// CancelTask tries to remove a task with the corresponding id, will not cancel a task that was already started.
func (r *Robot) CancelTask(id string) error {
	// todo: handle cancelling an active task?
	return r.tasks.RemoveTask(id)
}

// Mode returns the robots position and crate-held status.
func (r Robot) State() librobot.RobotState {
	return r.state
}

// Mode returns the robots next position and crate-held status, as a result of a command.
func (r Robot) NextState() librobot.RobotState {
	return r.nextState
}

// Name returns the robot name.
func (r Robot) Name() string {
	return r.name
}

// StatusInfo returns a message containing the robot status, used for debugging.
func (r Robot) StatusInfo() string {
	return "isRunning: " + strconv.FormatBool(r.isRunning) + ", mode: " + strconv.Itoa(int(r.mode)) + ", tasks: " + strconv.Itoa(r.tasks.Count())
}

// ALL DONE.
