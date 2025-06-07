package todo

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	Title       string `json:"title"`
	Description string `json:"descrition"`
	Paragraph   string `json:"paragraph"`
	CreateDate  string `json:"create_date"`
	FinishDate  string `json:"finish_date"`
}

func Todo() []Task {
	var tasks []Task
	var task Task
	scanner:=bufio.NewScanner(os.Stdin)
	fmt.Println("Write Title of the task:")
	scanner.Scan()
	task.Title=scanner.Text()		
	fmt.Println("Write description of the task:")
	scanner.Scan()
	task.Description=scanner.Text()		
	tasks = append(tasks, task)
	fmt.Printf("%+v", tasks)
	return tasks
}
