package main

import (
	"fmt"

	"github.com/kamlendu1982/sre-gig-tasks/task1/objectives"
)

func main() {
	fmt.Println("Result obj1")
	//Obj1()
	for i, v := range objectives.Obj1() {
		fmt.Println(i, v)
	}
	//fmt.Println(objectives.Obj1())
	fmt.Println("Result obj2")
	//objectives.Obj2()
	for i, v := range objectives.Obj2() {
		fmt.Println(i, v)
	}
}
