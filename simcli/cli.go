package simcli

import (
	"bufio"
	"fmt"
	"github.com/solimov-advanced/robot-challenge/librobot"
	"github.com/solimov-advanced/robot-challenge/robot"
	"github.com/solimov-advanced/robot-challenge/warehouse"
	"os"
	"strconv"
	"strings"
	"time"
)

// Cli is an abstraction of the command-line interface for the simulator.
type Cli struct {
	warehouses []librobot.Warehouse
	tasks      []activeTask
	exit       bool
	isRunning  bool
}

// NewSimulatorCli returns a new empty Cli with initialised slices
func NewSimulatorCli() Cli {
	cli := Cli{
		warehouses: make([]librobot.Warehouse, 0),
		tasks:      make([]activeTask, 0),
		exit:       false,
		isRunning:  false,
	}
	printHelp()
	cli.isRunning = true
	go cli.loopCli()
	return cli
}

func (c Cli) IsRunning() bool {
	return c.isRunning
}

func (c *Cli) loopCli() {
	c.isRunning = true
	c.exit = false
	reader := bufio.NewReader(os.Stdin)
	for c.isRunning {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		c.parseInput(text)

		time.Sleep(10 * time.Millisecond)

		//for _, t := range c.tasks {
		//	pos := <-t.position
		//	fmt.Println("Task", t.taskId,", Robot:", t.robotId,", x:", pos.X,", y:", pos.Y,", hasCrate:", pos.HasCrate)
		//}
		if c.exit {
			c.isRunning = false
		}
	}
}

func (c *Cli) parseInput(input string) {
	fields := strings.Fields(input)
	if len(fields) < 1 {
		return
	}
	switch fields[0] {
	case "exit":
		fmt.Println("quitting...")
		c.exit = true
		return
	case "quit":
		fmt.Println("quitting...")
		c.exit = true
		return
	case "help":
		printHelp()
		return
	case "add":
		if len(fields) > 1 {
			if fields[1] == "warehouse" {
				if len(fields) > 4 {
					id := fields[2]
					x, err1 := strconv.Atoi(fields[3])
					y, err2 := strconv.Atoi(fields[4])
					if err1 == nil && 0 < x && err2 == nil && 0 < y {
						c.AddWarehouse(id, x, y)
						return
					}
				}
			} else if fields[1] == "robot" {
				if len(fields) > 5 {
					wId := fields[2]
					rId := fields[3]
					x, err1 := strconv.Atoi(fields[4])
					y, err2 := strconv.Atoi(fields[5])
					if err1 == nil && 0 < x && err2 == nil && 0 < y {
						c.AddRobot(wId, rId, x, y)
						return
					}
				}
			} else if fields[1] == "task" {
				if len(fields) > 4 {
					wId := fields[2]
					rId := fields[3]
					commands := fields[4]
					c.AddTask(wId, rId, commands)
					return
				}
			} else if fields[1] == "crate" {
				if len(fields) > 4 {
					wId := fields[2]
					x, err1 := strconv.Atoi(fields[3])
					y, err2 := strconv.Atoi(fields[4])
					if err1 == nil && 0 < x && err2 == nil && 0 < y {
						c.AddCrate(wId, x, y)
						return
					}
				}
			}
		}
		fmt.Println("usage: [add warehouse \"id\" x y]")
		fmt.Println("usage: [add robot \"r_id\" \"w_id\" x y]")
		fmt.Println("usage: [add crate \"w_id\" x y]")
		return
	case "list":
		fmt.Println("|-idx-|------Name------|-Robots-|---Size---|")
		for n, w := range c.warehouses {
			x, y := w.Size()
			fmt.Println("| ", n, " | ", w.Name(), " | ", len(w.Robots()), " | (", x, ", ", y, ") |")
		}
		return
	case "status":
		if len(fields) > 1 {
			if fields[1] == "warehouse" {
				if len(fields) > 2 {
					id := fields[2]
					i := c.findWarehouseId(id)
					if i < 0 {
						fmt.Println("Error: warehouse not found")
						return
					}
					fmt.Println("warehouse", id, "status:")
					fmt.Println("Robots:")
					for _, r := range c.warehouses[i].Robots() {
						s := r.State()
						fmt.Println(r.Name(), " @(", s.X, ", ", s.Y, ") hasCrate:", strconv.FormatBool(s.HasCrate))
					}
					temp, ok := c.warehouses[i].(librobot.CrateWarehouse)
					if ok {
						fmt.Println("Crates:")
						for _, c := range temp.Crates() {
							fmt.Println("(", c.X, ", ", c.Y, ")")
						}
					}
					return
				}
			} else if fields[1] == "robot" {
				if len(fields) > 2 {
					id := fields[2]
					for _, w := range c.warehouses {
						wId := w.Name()
						i := c.findRobotId(wId, id)
						if i > -1 {
							r := w.Robots()[i]
							fmt.Println("robot", r.Name(), "status:")
							fmt.Println(r.StatusInfo())
							return
						}
					}
					fmt.Println("Error: robot not found")
					return
				}
			}
		}
		return
	}
}

func (c *Cli) AddWarehouse(id string, x, y int) {
	index := c.findWarehouseId(id)
	if index > -1 {
		fmt.Println("Error: id is used by other warehouse")
		return
	}
	newWarehouse := warehouse.NewWarehouse(id, uint(x), uint(y))
	c.warehouses = append(c.warehouses, newWarehouse)
}

func (c *Cli) AddRobot(warehouseId, id string, x, y int) {
	wIdx := c.findWarehouseId(warehouseId)
	if wIdx < 0 {
		fmt.Println("Error: warehouse not found")
		return
	}
	rIdx := c.findRobotId(warehouseId, id)
	if rIdx > -1 {
		fmt.Println("Error: id is used by other newRobot")
		return
	}
	newRobot := robot.NewRobot(id, uint(x), uint(y), robot.MobileDiagonalCraneBot, c.warehouses[wIdx])
	newRobot.Start()
	c.warehouses[wIdx].AddRobot(&newRobot)
}

func (c *Cli) AddCrate(warehouseId string, x, y int) {
	wIdx := c.findWarehouseId(warehouseId)
	if wIdx < 0 {
		fmt.Println("Error: warehouse not found")
		return
	}

	temp, ok := c.warehouses[wIdx].(librobot.CrateWarehouse)
	if ok {
		temp.AddCrate(uint(x), uint(y))
	}
}

func (c *Cli) AddTask(warehouseId, robotId, commands string) {
	wIdx := c.findWarehouseId(warehouseId)
	if wIdx < 0 {
		fmt.Println("Error: warehouse not found")
		return
	}
	rIdx := c.findRobotId(warehouseId, robotId)
	if rIdx < 0 {
		fmt.Println("Error: r not found")
		return
	}
	r := c.warehouses[wIdx].Robots()[rIdx]
	taskId, pos, err := r.EnqueueTask(commands)
	task := activeTask{
		robotId:  robotId,
		taskId:   taskId,
		position: pos,
		err:      err,
	}
	c.tasks = append(c.tasks, task)
}

func (c *Cli) QuitIfNoTasks() {
	for _, w := range c.warehouses {
		for _, r := range w.Robots() {
			if r.HasTasks() {
				return
			}
		}
	}
	c.exit = true
}

func (c Cli) findWarehouseId(id string) int {
	for i, w := range c.warehouses {
		if w.Name() == id {
			return i
		}
	}
	return -1
}

func (c Cli) findRobotId(warehouseId, id string) int {
	w_index := -1
	for i, w := range c.warehouses {
		if w.Name() == warehouseId {
			w_index = i
			break
		}
	}
	if w_index > -1 {
		for i, r := range c.warehouses[w_index].Robots() {
			if r.Name() == id {
				return i
			}
		}
	}
	return -1
}

func printHelp() {
	fmt.Println("------------------ROBOT-WAREHOUSE-SIMULATOR------------------")
	fmt.Println("COMMANDS:")
	fmt.Println("[help]")
	fmt.Println("		-show list of available commands.")
	fmt.Println("[list]")
	fmt.Println("		-lists all warehouses statuses.")
	fmt.Println("[add warehouse \"id\" x y]")
	fmt.Println("		-creates a new warehouse with name=\"id\" size x,y.")
	fmt.Println("[add robot \"w_id\" \"r_id\" x y]")
	fmt.Println("		-adds a new robot to warehouse=\"w_id\" with name=\"r_id\" at position x,y.")
	fmt.Println("[add crate \"w_id\" x y]")
	fmt.Println("		-adds a new crate to warehouse=\"w_id\" at position x,y.")
	fmt.Println("[add task \"w_id\" \"r_id\" \"commands\"]")
	fmt.Println("		-issues a new task to a robot in warehouse=\"w_id\" with name=\"r_id\".")
	fmt.Println("[status warehouse \"id\"")
	fmt.Println("		-displays info about warehouse=\"w_id\".")
	fmt.Println("[status robot \"id\"")
	fmt.Println("		-displays info about robot=\"r_id\".")
	fmt.Println("[exit] or [quit]")
}
