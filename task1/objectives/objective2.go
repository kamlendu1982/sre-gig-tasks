package objectives

import "fmt"

func Obj2() []string {
	Menu := [5]string{"apple", "kiwi", "banana", "mango", "avacado"}
	var result []string
	for i := 0; i < len(Menu); i++ {
		str_to_print := fmt.Sprintf("%s%s%s%d", "This is ", Menu[i], " and its index in the array is ", i)
		result = append(result, str_to_print)
		//fmt.Println(str_to_print)
	}
	return result
}
