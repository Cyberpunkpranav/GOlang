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
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Write Title of the task:")
	scanner.Scan()
	task.Title = scanner.Text()
	fmt.Println("Write description of the task:")
	scanner.Scan()
	task.Description = scanner.Text()
	tasks = append(tasks, task)
	fmt.Printf("%+v", tasks)
	return tasks
}

// func FizzBuzz(n int) map[int]string {
// 	var i int
// 	result := make(map[int]string)
// 	for i = 1; i <= n; i++ {
// 		if i%3 == 0 && i%5 == 0 {
// 			// fmt.Println("FizzBuzz")
// 			result[i] = "FizzBuzz"
// 			continue
// 		} else if i%3 == 0 {
// 			// fmt.Println("Fizz")
// 			result[i] = "Fizz"
// 			continue
// 		} else if i%5 == 0 {
// 			// fmt.Println("Buzz")
// 			result[i] = "Buzz"
// 			continue
// 		}
// 		// fmt.Println(i)
// 		result[i] = strconv.Itoa(i)
// 	}
// 	return result
// }
