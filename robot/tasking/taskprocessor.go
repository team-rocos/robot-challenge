package tasking

// TaskProcessor provides an abstraction of a queue of commands that are being actively processed.
type TaskProcessor struct {
	id       string
	commands string
	index    int
}

// NewActiveTask returns a new TaskProcessor for a given Task.
func (t *TaskProcessor) NewActiveTask(task Task) {
	t.id = task.Id
	t.commands = task.Commands
	t.index = 0
}

// IsEndReached returns true when there are no more commands to be processed.
func (t TaskProcessor) IsEndReached() bool {
	return t.index >= len(t.commands)
}

// GetNextCommand returns the next command in the TaskProcessor that matches one of the provided commands.
// Command strings have a maximum length of 2 characters.
func (t *TaskProcessor) GetNextCommand(commands []string) string {

	for !t.IsEndReached() {
		if t.index+1 < len(t.commands) && len(commands) > 0 {
			// check if next two chars make a valid command
			nextCmd := t.commands[t.index : t.index+2]
			isValid := false
			for _, cmd := range commands {
				if cmd == nextCmd {
					isValid = true
					break
				}
			}
			if isValid {
				t.index += 2
				return nextCmd
			}
		}

		// check if single char is a valid command
		nextCmd := t.commands[t.index : t.index+1]
		t.index += 1
		isValid := false
		for _, cmd := range commands {
			if cmd == nextCmd {
				isValid = true
				break
			}
		}
		if isValid {
			return nextCmd
		}
	}
	return ""
}

// ALL DONE.
