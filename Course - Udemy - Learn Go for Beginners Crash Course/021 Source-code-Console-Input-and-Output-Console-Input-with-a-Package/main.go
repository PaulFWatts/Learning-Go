package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

func main() {
	// Open the keyboard for reading.
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	// Make sure to close the keyboard in the main function before the program exits!
	defer func() {
		_ = keyboard.Close() // Ignore any errors.
	}()

	fmt.Println("Press any key on the keyboard. Press ESC to quit.")

	for {
		// Read a single key from the keyboard.
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		// Print the key that was pressed.
		if key != 0 {
			fmt.Println("You pressed", char, key)
		} else {
			fmt.Println("You pressed", char)
		}

		// Check if the pressed key was ESC.
		if key == keyboard.KeyEsc {
			break
		}
	}

	fmt.Println("Program exiting...")
}
