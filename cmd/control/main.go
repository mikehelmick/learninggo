package main

import (
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func basicIf() {
	x := rand.Intn(2)
	if x == 0 {
		log.Printf("X was zero")
	} else {
		log.Printf("X was not zero: %v", x)
	}
}

func assignIf() {
	if x := rand.Intn(2); x == 0 {
		log.Printf("X was zero")
	} else {
		log.Printf("X was not zero: %v", x)
	}
}

func main() {
	basicIf()
	assignIf()
}
