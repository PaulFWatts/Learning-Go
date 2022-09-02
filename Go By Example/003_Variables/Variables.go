package main

/*
	1. The var statement declares a list of variables; as in function argument lists, the type is last.
	2. A var statement can be at package or function level. We see both in this example.
	3. Variables declared without an explicit initial value are given their zero value.
	4. The zero value is:
		0 for numeric types,
		false for the boolean type, and
		"" (the empty string) for strings.
	5. The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case.
	6. Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
*/

import "fmt"

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	// Press enter to continue
	fmt.Println("Press 'Enter' to continue...")
	fmt.Scanln()

}