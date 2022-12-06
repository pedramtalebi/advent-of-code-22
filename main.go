package main

import (
	"fmt"
	"github.com/pedramtalebi/advent-of-code-22/day-1"
	day_2 "github.com/pedramtalebi/advent-of-code-22/day-2"
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
	totalScore, err3 := day_2.SolutionOne()
	if err3 != nil {
		fmt.Errorf("error from day one task two: %s", err2)
	}
	totalScoreSecond, err4 := day_2.SolutionTwo()
	if err4 != nil {
		fmt.Errorf("error from day one task two: %s", err2)
	}

	fmt.Println(highestCalories)
	fmt.Println(totalCalories)
	fmt.Println(totalScore)
	fmt.Println(totalScoreSecond)
}
