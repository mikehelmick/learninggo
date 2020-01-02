package semicolons

import "log"

func doSomething() {
	log.Printf("I did a thing!")
}

func test() {
	var a, b int

	if a < b {
		doSomething()
	}

  if a < b
  {
    doSomething()
  }
}
