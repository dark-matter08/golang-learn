//--Summary:
//  Create a program that can perform dice rolls. The program should
//  report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must handle any number of dice, rolls, and sides
//
//--Notes:
//* Use packages from the standard library to complete the project

package main

import (
	"fmt"
	"math/rand"
)

func roll(sides int) int {
	return rand.Intn(sides) + 1
}

func main() {
	dice, sides := 2, 6
	rolls := 2

	for r := 1; r <= rolls; r++ {
		var sum int

		fmt.Println("========: Roll No", r, ":========")
		fmt.Println("Total Die:", dice)
		for d := 1; d <= dice; d++ {
			rolled := roll(sides)
			fmt.Println("Roll Die#", d, ":", rolled)

			sum += rolled
		}

		fmt.Println("Total rolled:", sum)

		// if sum == 2 && dice == 2 {
		// 	fmt.Println("Snake eyes")
		// } else if sum == 7 {
		// 	fmt.Println("Lucky 7")
		// } else if sum%2 == 0 {
		// 	fmt.Println("Even")
		// } else if sum%2 == 1{
		// 	fmt.Println("Odd")
		// }

		switch sum := sum; {
		case sum == 2 && dice == 2:
			fmt.Println("==> Snake eyes")
		case sum == 7:
			fmt.Println("==> Lucky 7")
		case sum%2 == 0:
			fmt.Println("==> Even")
		case sum%2 == 1:
			fmt.Println("==> Odd")
		}
	}
}
