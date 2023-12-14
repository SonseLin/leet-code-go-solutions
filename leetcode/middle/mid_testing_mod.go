package middle

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
			case 3:
				TestLSS()
			case 11:
				TestCWMW()
			case 2482:
				TestOMZ()
			default:
				fmt.Println("I didn`t complete this task")
			}
		} else {
			panic(err)
		}
	}
}
