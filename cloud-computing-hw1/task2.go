package main

import ("fmt")

func main() {

	// Attributes
	// Variables Declaration
	// Baking a Chocolate Cupcakes
	typeOfCupcake := "Chocolate Chip"   // string: Type or flavor of the cupcake
	ovenTemperature := 180.05           // float64: Oven temperature needed for baking
	numberOfCupcakes := 12              // integer: Number of cupcakes being baked


    // Need to put outside the function so when running the program each line of the details will not have this
    // "\n" is use for paragraph spacing
	fmt.Println("\n------ Cupcake Info ------\n")


	// Displaying Baking Details
	displayBakingDetail("Type of Cupcake", typeOfCupcake)	
	displayBakingDetail("Oven Temperature", ovenTemperature)
	displayBakingDetail("Number of Cupcakes", numberOfCupcakes)

	fmt.Println("\n***********************")     // last line

}

// Function for displaying the Details
func displayBakingDetail(detail string, value interface{}) {
	fmt.Printf("%s: %v, Data Type: %T\n", detail, value, value)
 }