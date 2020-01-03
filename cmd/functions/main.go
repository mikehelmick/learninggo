package main

import "log"

func simpleFunction() {
	log.Printf("You called a simple function")
}

func add(a, b int) int {
	return a + b
}

func divWithRemainder(a, b int) (int, int) {
	div := a / b
	rem := a % b
	return div, rem
}

func betterDivWithRemainder(a, b int) (div, rem int) {
	div = a / b
	rem = a % b
	return
}

func sum(a ...int) (result int) {
	for _, v := range a {
		result += v
	}
	return
}

type foo int

func (a foo) increment() foo {
	return a + 1
}

func main() {
	simpleFunction()

	log.Printf("Add function: %v", add(2, 5))

	{
		d, r := divWithRemainder(42, 5)
		log.Printf("42 / 5 = %v rem %v", d, r)
	}

	{
		d, r := betterDivWithRemainder(42, 5)
		log.Printf("42 / 5 = %v rem %v", d, r)
	}

	total := sum(21, 15, 9)
	log.Printf("Sum is: %v", total)

	{
		var a foo = 5
		log.Printf("a is %v", a)
		a = a.increment()
		log.Printf("a is %v", a)
	}
}
