package objectives

import (
	"strings"
	"testing"
)

func TestObj2(t *testing.T) {
	t.Run("Check apple in plate", func(t *testing.T) {
		res := false
		list_furits := Obj2()
		for i := 0; i < len(list_furits); i++ {
			if strings.Contains(list_furits[i], "apple") {
				res = true
				break
			}
		}
		if !res {
			t.Errorf("Apple is missing from the plate")
		}
	})
}
