package main

import (
	leetez "GOleet/leetcode/easy"
	leethard "GOleet/leetcode/hard"
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
				leetez.TestingModule()
			case "middle":
				leetmid.TestingModule()
			case "hard":
				leethard.TestingModule()
			default:
				ok = false
			}
		}
	}
}
