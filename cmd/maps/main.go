package main

import (
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func rollDice(res map[int]int) {
	d := rand.Intn(6) + 1
	// This works because of the default value of something not being present.
	res[d]++
}

func membershipDemo() {
	code := map[string]string{
		"a": "5",
		"b": "t"}

	log.Printf("code[a] = '%v'", code["a"])
	log.Printf("code[b] = '%v'", code["b"])
	log.Printf("code[c] = '%v'", code["c"])

	if _, ok := code["c"]; !ok {
		log.Printf("code does not contain the key c")
	}
}

func main() {
	results := make(map[int]int)
	for i := 0; i < 1000; i++ {
		rollDice(results)
	}

	for k, v := range results {
		log.Printf("%1d = %2d", k, v)
	}
	log.Printf("%v", results)

	membershipDemo()
}
