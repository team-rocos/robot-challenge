package tasking

// Task provides an abstraction of a queue of commands.
type Task struct {
	Id       string
	Commands string
}

func (t Task) IsEmpty() bool {
	return len(t.Commands) == 0
}

// ALL DONE.
