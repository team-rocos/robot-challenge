package tasking

import (
	"errors"
	"strconv"
)

// TaskQueue provides an abstraction for a queue of Task.
type TaskQueue struct {
	nextId int
	tasks  []Task
}

// AddTask appends a task to the queue and returns the id that was assigned to it.
func (q *TaskQueue) AddTask(commands string) (id string) {
	if q.tasks == nil {
		q.tasks = make([]Task, 0)
	}

	task := Task{strconv.Itoa(q.nextId), commands}
	q.tasks = append(q.tasks, task)
	q.nextId++
	return task.Id
}

// RemoveTask attempts to remove a task from the queue with matching id.
func (q *TaskQueue) RemoveTask(id string) error {
	index := -1
	for i, t := range q.tasks {
		if t.Id == id {
			index = i
			break
		}
	}

	if index > -1 {
		// if task index found, remove it
		q.tasks = append(q.tasks[:index], q.tasks[index+1:]...)
		return nil
	} else {
		return errors.New("remove failed: id not found")
	}
}

func (q TaskQueue) Count() int {
	return len(q.tasks)
}

// GetNextTask returns the next task and removes it from the queue.
func (q *TaskQueue) GetNextTask() Task {
	if len(q.tasks) > 0 {
		nextTask := q.tasks[0]
		q.tasks = append(q.tasks[:0], q.tasks[1:]...) // Remove task at index=0
		return nextTask
	} else {
		return Task{}
	}
}

// ALL DONE.
