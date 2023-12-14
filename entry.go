package main

import (
	leetmid "GOleet/leetcode/middle"
	"fmt"
)

func main() {
	fmt.Println("I`m entry point")
	ok := true
	for ok {
		fmt.Println("Input difficult(easy/middle/hard)")
		var difficultType string
		_, err := fmt.Scanf("%s", &difficultType)
		if err != nil {
			panic(err)
		} else {
			switch difficultType {
			case "easy":
			case "middle":
				leetmid.TestingModule()
			case "hard":
			default:
				ok = false
			}
		}
	}
}
