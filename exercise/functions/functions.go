//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func great(name string) {
	fmt.Println("Hello there", name)
}

func hiThere() string {
	return "Yoooo, dude!"
}

func multipleAddition(x, y, z int) int {
	return x + y + z
}

func returningFive() int {
	return 5
}

func returningTwoNumbers() (int, int) {
	return 20, 21
}

func main() {
	great("Nde Lucien")
	fmt.Println(hiThere())
	number1, number2 := returningTwoNumbers()
	finalSumation := multipleAddition(returningFive(), number1, number2)

	fmt.Println("The final sumation is:", finalSumation)

}
