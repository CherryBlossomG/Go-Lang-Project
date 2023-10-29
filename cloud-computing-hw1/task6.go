package main

import ("fmt")

func main() {
	var num1, num2 float64
	// Float64 for decimal numbers
	// Ask the user to input the first number
	// Scanner is being used
	// If the user input is not equal to a number it will print an Error Line
	// nil = null
	// err = error
	fmt.Print("Enter first number: ")
	_, err1 := fmt.Scanln(&num1)
	if err1 != nil {
		fmt.Println("Error: Strings or special characters are not allowed!")
		return
	}

	// Ask the user to input the second number
	fmt.Print("Enter second number: ")
	_, err2 := fmt.Scanln(&num2)
	if err2 != nil {
		fmt.Println("Error: Strings or Special Characters are not allowed!")
		return
	}

	// Check for division by zero
	// If the number is divided into zero it will print an Error Line
	if num2 == 0 {
		fmt.Println("Error: Division by zero is not allowed!")
		return
	}

	// Print line that displaying the details and result of the division numbers
	result := num1 / num2
	fmt.Printf("%.2f divided by %.2f is %.2f\n", num1, num2, result)
}
