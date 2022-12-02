package day_1

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
)

func SolutionOne() (int, error) {
	file, err := os.Open("day-1/input.txt")
	if err != nil {
		return 0, err
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	highestCalories := 0
	current := 0
	for fileScanner.Scan() {
		row := fileScanner.Text()
		if len(row) == 0 {
			if current > highestCalories {
				highestCalories = current
			}
			current = 0
			continue
		}
		lineP, _ := strconv.Atoi(row)
		current = current + lineP
	}
	return highestCalories, nil
}

func SolutionTwo() (int, error) {
	file, err := os.Open("day-1/input.txt")
	if err != nil {
		return 0, err
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	h := &MinHeap{}
	heap.Init(h)
	currentElfTot := 0

	for fileScanner.Scan() {
		row := fileScanner.Text()
		if len(row) == 0 {
			if h.Len() < 3 {
				heap.Push(h, currentElfTot)
			} else if (*h)[0] < currentElfTot {
				heap.Push(h, currentElfTot)
				heap.Pop(h)
			}
			currentElfTot = 0
			continue
		}
		rowP, _ := strconv.Atoi(row)
		currentElfTot = currentElfTot + rowP
	}

	return (*h)[0] + (*h)[1] + (*h)[2], nil
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
