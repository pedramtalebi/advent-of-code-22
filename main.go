package main

import (
	"fmt"
	"github.com/pedramtalebi/advent-of-code-22/day-1"
)

func main() {
	highestCalories, err1 := day_1.SolutionOne()
	if err1 != nil {
		fmt.Errorf("error from day one task one: %s", err1)
	}
	totalCalories, err2 := day_1.SolutionTwo()
	if err2 != nil {
		fmt.Errorf("error from day one task two: %s", err2)
	}
	fmt.Println(highestCalories)
	fmt.Println(totalCalories)
}
