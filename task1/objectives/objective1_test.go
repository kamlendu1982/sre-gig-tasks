package objectives

import (
	"fmt"
	"strings"
	"testing"
)

func TestObj1(t *testing.T) {
	t.Run("Check the length of return string", func(t *testing.T) {
		length := len(Obj1())
		exp_length := 2
		if length != exp_length {
			t.Errorf("Expected %d, but got %d", exp_length, length)
		}
	})

	t.Run("Test the Hamburger in result", func(t *testing.T) {
		food_list := Obj1()
		var is_present bool = false
		for i, v := range food_list {
			fmt.Println(i, v)
			if strings.Contains(v, "Hamburger") {
				is_present = true
				break
			}
		}
		if !is_present {
			t.Errorf("Hamburger is not in item")
		}
	})

	t.Run("Test the Salad in result", func(t *testing.T) {
		food_list := Obj1()
		var is_present bool = false
		for i, v := range food_list {
			fmt.Println(i, v)
			if strings.Contains(v, "Salad") {
				is_present = true
				break
			}
		}
		if !is_present {
			t.Errorf("Salad is not in item")
		}
	})
}
