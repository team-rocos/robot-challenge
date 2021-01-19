package warehouse

import (
	"errors"
	"github.com/solimov-advanced/robot-challenge/librobot"
)

// Warehouse provides an abstraction of a simulated warehouse containing both robots and crates.
type Warehouse struct {
	robots    []librobot.Robot
	crates    []librobot.Position
	x         uint
	y         uint
	isRunning bool
	name      string
}

// NewWarehouse returns a crate warehouse with a specified grid size.
func NewWarehouse(name string, sizeX, sizeY uint) librobot.CrateWarehouse {
	newWarehouse := Warehouse{
		robots:    make([]librobot.Robot, 0),
		crates:    make([]librobot.Position, 0),
		x:         sizeX,
		y:         sizeY,
		isRunning: false,
		name:      name,
	}
	println("- New 'Warehouse' was created, with size: (", sizeX, ", ", sizeY, ")")
	return &newWarehouse
}

func (w Warehouse) Name() string {
	return w.name
}

func (w Warehouse) Size() (x, y uint) {
	return w.x, w.y
}

func (w Warehouse) Robots() []librobot.Robot {
	return w.robots
}

func (w Warehouse) Crates() []librobot.Position {
	return w.crates
}

// AddCrate tries to add a new crate at a specified position.
func (w *Warehouse) AddCrate(x, y uint) error {
	if w.crates == nil {
		w.crates = make([]librobot.Position, 0)
	}
	isOccupied := false
	for _, c := range w.crates {
		if c.X == x && c.Y == y {
			isOccupied = true
			break
		}
	}

	if !isOccupied {
		w.crates = append(w.crates, librobot.Position{
			X: x,
			Y: y,
		})
		println("- New 'Crate' was created, at location: (", x, ", ", y, ")")
		return nil
	} else {
		return errors.New("proposed crate position is occupied")
	}
}

// DelCrate tries to delete a crate at a specified position.
func (w *Warehouse) DelCrate(x, y uint) error {
	index := -1
	for i, c := range w.crates {
		if c.X == x && c.Y == y {
			index = i
			break
		}
	}

	if index > -1 {
		w.crates = append(w.crates[:index], w.crates[index+1:]...)
		println("- Removed 'Crate' at location: (", x, ", ", y, ")")
		return nil
	} else {
		return errors.New("no crate at position")
	}
}

// AddRobot tries to add a new robot to the warehouse.
func (w *Warehouse) AddRobot(robot librobot.Robot) error {
	x := robot.State().X
	y := robot.State().Y
	if w.x <= x || w.y <= y {
		return errors.New("robot is outside warehouse bounds")
	}

	for _, r := range w.robots {
		if r.Name() == robot.Name() {
			return errors.New("a robot in this warehouse has this name")
		}
		if r.State().X == x && r.State().Y == y {
			return errors.New("a robot in this warehouse occupies this position")
		}
	}

	w.robots = append(w.robots, robot)
	println("- New 'Robot' named \""+robot.Name()+"\" was added, at location: (", robot.State().X, ", ", robot.State().Y, ")")
	return nil
}

// GiveTask tries to issue a task string to a robot with specified name.
func (w *Warehouse) GiveTask(robotName string, task string) error {
	index := -1
	for i, r := range (*w).Robots() {
		if r.Name() == robotName {
			index = i
			break
		}
	}

	if index > -1 {
		(*w).Robots()[index].EnqueueTask(task)
		return nil
	} else {
		return errors.New("robot name not found")
	}
}

// ALL DONE.
