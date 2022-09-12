package main

/*
	1. Go supports constants of character, string, boolean, and numeric values.
	2. const declares a constant value.
	3. A const statement can appear anywhere a var statement can.
	4. Constants cannot be declared using the := syntax.
	5. Numeric constants are high-precision values.
	6. Constant expressions perform arithmetic with arbitrary precision.
	7. A numeric constant has no type until it’s given one, such as by an explicit conversion.
	8. A number can be given a type by using it in a context that requires one, such as a variable
		assignment or function call. For example, here math.Sin expects a float64.
*/

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

	// Press enter to continue
	fmt.Println("Press 'Enter' to continue...")
	fmt.Scanln()
}
