package main

import "fmt"

func obj1() {
	var Menu []string
	Menu = append(Menu, "Hamburger")
	Menu = append(Menu, "Salad")
	for i := 0; i < len(Menu); i++ {
		fmt.Println("Food: " + Menu[i])
	}
}

func obj2() {
	Menu := [5]string{"apple", "kiwi", "banana", "mango", "avacado"}
	for i := 0; i < len(Menu); i++ {
		str_to_print := fmt.Sprintf("%s%s%s%d", "This is ", Menu[i], " and its index in the array is ", i)
		fmt.Println(str_to_print)

	}
}

func main() {
	fmt.Println("Result obj1")
	obj1()
	fmt.Println("\n")
	fmt.Println("Result obj2")
	obj2()
}
