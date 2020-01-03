package main

import (
	"log"
	"time"
)

type message struct {
	sender string
	value  int
}

func ping(name string, in, out chan message) {
	for {
		m := <-in
		log.Printf("%v received %4d", name, m.value)
		time.Sleep(10 * time.Millisecond)
		out <- message{name, m.value + 1}
	}
}

func main() {
	a := make(chan message)
	b := make(chan message)
	defer close(a)
	defer close(b)

	go ping("ping", a, b)
	go ping("pong", b, a)

	a <- message{"start", 0}

	select {
	case <-time.After(1 * time.Second):
		log.Printf("done")
	}
}
