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

  // This file intentionally doesn't compile.
  if a < b
  {
    doSomething()
  }
}
