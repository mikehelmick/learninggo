package main

import "log"

func main() {

	var a = 0
	var b int
	log.Printf("a = %v, b = %v", a, b)

	c := 0
	log.Printf("c = %v", c)

	d := int64(24)
	var e int64 = 24
	log.Printf("d = %v, e = %v", d, e)

	// Incompatible assignments.
	// a = d
	// d = a

	s := "This is a string."
	log.Printf("s = %s", s)
}
