package objective

import (
	"strings"
	"testing"
)

func TestGen_Random(t *testing.T) {
	t.Run("Test number less then 50", func(t *testing.T) {
		num_test := 40
		res_string := Gen_random(num_test)
		exp_substring := "closer to 0"
		if !strings.Contains(res_string, exp_substring) {
			t.Errorf("Something not right, expecting closer to 0")
		}
	})

	t.Run("Test number 50", func(t *testing.T) {
		num_test := 50
		res_string := Gen_random(num_test)
		exp_substring := "It 50"
		if !strings.Contains(res_string, exp_substring) {
			t.Errorf("Something not right, expecting 50")
		}
	})

	t.Run("Test greatter then 50", func(t *testing.T) {
		num_test := 70
		res_string := Gen_random(num_test)
		exp_substring := "closer to 100"
		if !strings.Contains(res_string, exp_substring) {
			t.Errorf("Something not right, expecting closer to 100")
		}
	})

	t.Run("Test if number is even", func(t *testing.T) {
		num_test := 80
		res_string := Gen_random(num_test)
		exp_substring := "even"
		if !strings.Contains(res_string, exp_substring) {
			t.Errorf("Something not right, expecting it's even")
		}
	})

}
