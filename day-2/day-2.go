package day_2

import (
	"bufio"
	"os"
)

func SolutionOne() (int, error) {
	file, err := os.Open("day-2/input.txt")
	if err != nil {
		return 0, err
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	totalScore := 0
	for fileScanner.Scan() {
		row := fileScanner.Text()
		score := calculateScore(row)
		totalScore += score
	}

	return totalScore, nil
}

func SolutionTwo() (int, error) {
	file, err := os.Open("day-2/input.txt")
	if err != nil {
		return 0, err
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	totalScore := 0
	for fileScanner.Scan() {
		row := fileScanner.Text()
		score := calculateScoreSecond(row)
		totalScore += score
	}
	return totalScore, nil
}

// A for Rock, B for Paper, and C for Scissors.
// X for Rock, Y for Paper, and Z for Scissors
// The score for a single round is the score for the shape you selected
// (1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for the outcome of the round
// (0 if you lost, 3 if the round was a draw, and 6 if you won)
func calculateScore(row string) int {
	scores := map[string]int{
		"A X": 1 + 3,
		"A Y": 2 + 6,
		"A Z": 3 + 0,
		"B X": 1 + 0,
		"B Y": 2 + 3,
		"B Z": 3 + 6,
		"C X": 1 + 6,
		"C Y": 2 + 0,
		"C Z": 3 + 3,
	}

	return scores[row]
}

// A for Rock, B for Paper, and C for Scissors.
// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win."
func calculateScoreSecond(row string) int {
	scores := map[string]int{
		"A X": 3 + 0,
		"A Y": 1 + 3,
		"A Z": 2 + 6,
		"B X": 1 + 0,
		"B Y": 2 + 3,
		"B Z": 3 + 6,
		"C X": 2 + 0,
		"C Y": 3 + 3,
		"C Z": 1 + 6,
	}

	return scores[row]
}
