package easy

import "fmt"

func TestingModule() {
	ok := true
	for ok {
		fmt.Println("Input number of task from leetcode")
		var numberOfTask int
		_, err := fmt.Scanf("%d", &numberOfTask)
		if err == nil {
			switch numberOfTask {
			case 0:
				ok = false
			case 867:
				TestTP()
			default:
				fmt.Println("I didn`t complete this task")
			}
		} else {
			panic(err)
		}
	}
}
