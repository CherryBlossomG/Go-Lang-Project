package main

// Packages
import (
	"fmt"
	"strings"
)

func main() {
	// Create a slice of integers
	// List of integers
	numbers := []int{1, 2, 3, 4}

	// Calculate sum of slice elements using Loop
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("Sum:", sum)

	// Creating a dictionary map string keys with an integer value
	dict := map[string]int{
		"Phone":    100,
		"Laptop":   200,
		"LCD":     250,
		"Mouse":   300,
	}
	// For Adding item/s to the map
     dict["Keyboard"] = 350
	

	for {
		// Ask the user for input
		fmt.Println("\nEnter an item name:")
		var userInput string
		fmt.Scanln(&userInput)
		userInput = strings.TrimSpace(userInput) // Trim or removing any leading or extra spaces

		// Check if the user input exists in the dictionary then it will print the corresponding value else it will print we dont have
		// If Else Statement
		value, exists := dict[userInput]
		if exists {
			fmt.Println(userInput, "value is:", value)
		} else {
			fmt.Println("We don't have", userInput)
		}

		// Ask user if they want to check more items
		// if "yes" it will ask what items the user wants to check.
		// if "no" it will just print "Thank you".

		fmt.Println("Do you want to check more? (yes/no)")
		var decision string
		fmt.Scanln(&decision)
		decision = strings.ToLower(strings.TrimSpace(decision))

		if decision != "yes" {
			break
		}
	}

	fmt.Println("Thank you!")
}