package main

import ("fmt")

func factorial(n int) int { 
	//function for factorial of a given positive integer by the user input
	//if statement
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

//main function
func main() {

	//attributes using scanner
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	
	//if statement to avoid negative integers.
	//if the integer that the user put is negative it will print the line below

	if num < 0 {
		fmt.Println("Please enter a non-negative number.")
		return
	}

	fmt.Printf("Factorial of %d is: %d\n", num, factorial(num))
}
