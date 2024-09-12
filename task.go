package main
import "fmt"

type task struct {
    name    string
    status  string
    urgency int
}

// make new task
func newTask(name string) task {
    t := task{
        name:    name,
        status:  "Todo",
        urgency: 0,
    }
    return t
}

func (t task) formatTask() string {
    formattedTask := fmt.Sprintf("%v: \n Status........%v \n Urgency........%v", t.name, t.status, t.urgency)
    return formattedTask
}

func (t *task) updateTaskUrgency(urgency int) {
    t.urgency = urgency
}

func (t *task) updateTaskStatus(status string) {
    t.status = status
}