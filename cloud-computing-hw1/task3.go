package main

import ("fmt")

func main() {       //main function
	// Attribute
	// variable
	var num int

	//input from user
	fmt.Print("Enter a number: ") //print line
	fmt.Scan(&num) //scanner input

	// Check if the number is odd or even
	// If Else Statement is being used
	if num%2 == 0 {                          // == 0 because even can divide by 2 so 2 divided 2 = 0, can change it into num%2 == 1 to make it as even
		fmt.Println("The number is even.")  // if the number is even it will print this line
	} else {
		fmt.Println("The number is odd.")   // if the number is odd it will print this line
	}

	// For Loop to print numbers from 1 to the input number
	// I use i++ so that the number will increase up to the number put by the user.

	for i := 1; i <= num; i++ {
		fmt.Println(i)
	}
}