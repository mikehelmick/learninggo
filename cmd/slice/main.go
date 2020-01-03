package main

import "log"

func doubleAll(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}

func printAll(s []int) {
	for i, v := range s {
		log.Printf("idx %3d = %3d", i, v)
	}
}

func main() {
	s := make([]int, 0, 2)

	for i := 0; i < 20; i++ {
		s = append(s, i)
		log.Printf("%3d :: len = %3d, cap = %3d", i, len(s), cap(s))
	}

	printAll(s)
	doubleAll(s)
	printAll(s)
	// Create a new slice without the first 2 or last 2 elements of s
	trimEnds := s[2 : len(s)-2]
	printAll(trimEnds)
	// Show that they still point to the same underlying memory
	trimEnds[0] = 50
	printAll(s)

	// Show array and slice types are distinct
	// arr := [2]int{1, 2}
	// printAll(arr)

	{
		// Example of 2D slices
		board := make([][]int, 5)
		for i := 0; i < len(board); i++ {
			board[i] = make([]int, 5)
		}
		i := 0
		for y := 0; y < len(board); y++ {
			for x := 0; x < len(board[y]); x++ {
				board[y][x] = i
				i++
			}
		}
		log.Printf("Board: %v", board)
	}
}
