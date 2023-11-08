package objective

import (
	"fmt"
)

func Gen_random(ran_num int) string {

	//var ran_num int = rand.Intn(100)
	fmt.Println(ran_num)
	var result string = "It 50!"
	if ran_num < 50 {
		result = "It's closer to 0"
	} else if ran_num > 50 {
		result = "It's closer to 100"
		if ran_num%2 == 0 {
			result = "It's closer to 100, and t's even"
		}
	}
	return result
}
