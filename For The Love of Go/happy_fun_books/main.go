package main

import "fmt"

func main() {
	var title string
	var copies int
	var name string
	var specialOffer bool
	var discount float64

	title = "The Happiness Project"
	name = "John Harper"
	copies = 99
	discount = 0.10
	specialOffer = true
	print(title, name, copies, discount, specialOffer)
}

func print(title string, name string, copies int, discount float64, specialOffer bool) {
	fmt.Println(title)
	fmt.Println(name)
	fmt.Println(copies)
	fmt.Println(discount)
	fmt.Println(specialOffer)
}
