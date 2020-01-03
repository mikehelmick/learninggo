package main

import "log"

func sayHello() {
	log.Printf("Hello")
}

func main() {
	log.Printf("I'm going to say hello")
	go sayHello()
	log.Printf("I said hello")
}
