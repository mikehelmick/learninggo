package main

import "log"

func main() {

	for i := 0; i < 10; i++ {
		log.Printf("1st For loop i = %v", i)
	}

	done := false
	i := 0
	for !done {
		log.Printf("2nd For loop i = %v", i)
		i++
		done = i >= 10
	}

	i = 0
	for {
		log.Printf("3rd For loop i = %v", i)
		i++
		if i >= 10 {
			break
		}
	}
}
