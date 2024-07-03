//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	var favoriteColor = "Orange"
	fmt.Println("My favorite color is", favoriteColor)

	birthYear, ageInYears := 1999, 25
	fmt.Println("Born in", birthYear, "I am", ageInYears, "Years old")

	var (
		firstInitial = "N"
		lastInitial  = "N"
	)

	fmt.Println("My first and last initials are", firstInitial, "and", lastInitial)

	var ageInDays int

	ageInDays = ageInYears * 365

	fmt.Println("My total days around the sun is", ageInDays, "Days")

}
