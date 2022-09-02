package main

/*
	1. Go has various value types including strings, integers, floats, booleans, etc. Here are a few basic examples.
	2. Strings, which can be added together with +.
	3. Integers and floats.
	4. Booleans, with boolean operators as you’d expect.
*/

import "fmt"

func main() {

	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// Press enter to continue
	fmt.Println("Press 'Enter' to continue...")
	fmt.Scanln()
}
