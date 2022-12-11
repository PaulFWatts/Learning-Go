package main

import (
	"bufio"
	"fmt"
	"myapp/packageone"
	"os"
)

var one = "One"

func main() {
	var somethingElse = "this is a block level variable"
	fmt.Println(somethingElse)
	myFunc()

	newString := packageone.PublicVar
	fmt.Println("From packageone:", newString)

	packageone.Exported()
	pressAnyKey()

}

func myFunc() {
	fmt.Println(one)
}

func pressAnyKey() string {
	fmt.Println("Press Return To Continue...")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return input
}
