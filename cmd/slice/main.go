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

	// Show array and slice types are distinct
	// arr := [2]int{1, 2}
	// printAll(arr)
}
