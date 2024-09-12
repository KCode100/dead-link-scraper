package main
import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
    fmt.Print(prompt)
    input, err := r.ReadString('\n')
    return strings.TrimSpace(input), err
}

func createTask() task {
    reader := bufio.NewReader(os.Stdin)
    name, _ := getInput("Create a new task: ", reader)
    task := newTask(name)
    fmt.Println("Created the task - ", task.name)
    return task
}

func promptOptions(t task) {
    reader := bufio.NewReader(os.Stdin)
    option, _ := getInput("Choose an option (a - add a task, s - save the task, us - update the status): ", reader)
    switch option {
    case "a":
        fmt.Println("you chose a")
    case "s":
        fmt.Println("you chose s")
    case "us":
        fmt.Println("you chose us")
    default:
        fmt.Println("that was not a valid option")
        promptOptions(t)
    }
}

func main() {
    myTask := createTask()
    promptOptions(myTask)
}