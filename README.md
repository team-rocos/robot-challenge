Note: 

It's my first time using go and my first time using the interface concept. 

I haven't ever written code for anyone other than myself either!

I wasn't entirely sure how the channels returned by "EnqueueTask" were meant to be used.


# Robot Warehouse Simulation Library

## Packages (a brief description)
**librobot**  - interfaces for Warehouse and Robot objects, implemented by **warehouse** and **robot** packages

**warehouse** - simulated warehouse has a discrete size and has functions to add/remove robots & crates, start/stop robots and issue robot tasks

**robot**     - simulated robot runs a go routine to manage it's state and state logic. it is implemented by the **warehouse** package and implements the **tasking** and       **commanding** sub-packages

**robot/tasking**     - allows a robot functions to get next task, add a task, and get the next command from an active task

**robot/commanding**  - handles interpretation of a command and returns corresponding command struct that implements the generic command interface

## Description
Simulates a warehouse space and robots moving through it with discrete coordinate positions.
Robots act independently and reference the warehouse they are contained within, in this way the warehouse provides Robots with knowledge of their environment (position of other walls, robots and crates) checking the position of objects using the simulated-warehouse takes the place of sensors in a real setting.

Robots each run a go routine that manages their software-state and logic (such as starting/stopping tasks and commands)

Robots have a queue of tasks, each task is defined as a series of commands
A command is defined as an action, usually changing the robot's physical-state

Commands are run in a go routine that exits upon completion (in this simulator thats 1 second of real-time)

Command go files implement an interface and allow segregation of command logic

The commands.go file has a function to parse the robots command and return the correct corresponding Command implementation (i.e. move.go or crane.go)
