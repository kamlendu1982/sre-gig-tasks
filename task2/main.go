package main

import (
	"fmt"
	"math/rand"

	"github.com/kamlendu1982/sre-gig-tasks/task2/objective"
)

func main() {
	ran_num := rand.Intn(100)
	fmt.Println(objective.Gen_random(ran_num))
}
