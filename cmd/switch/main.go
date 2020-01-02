package main

import "log"

func main() {
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			log.Printf("%d is divisible by 2", i)
		case i%3 == 0:
			log.Printf("%d is divisible by 3", i)
		default:
			log.Printf("%d is not divisible by 2 or 3", i)
		}
	}
}
