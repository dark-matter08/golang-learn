//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func printSlice(title string, slice []Part) {
	fmt.Println()
	fmt.Println("---", title, "---")
	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println("=>", element)

	}
}

func main() {
	line := []Part{"start", "middle", "half-middle"}

	printSlice("Starting Assembly", line)

	line = append(line, "second-half-middle", "third-half-middle")
	printSlice("Extended Assembly", line)

	line = line[3:]

	printSlice("Sliced Assembly", line)

}
