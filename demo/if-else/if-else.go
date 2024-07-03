package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	quiz1, quiz2, quiz3 := 9, 7, 8

	if quiz1 > quiz2 {
		fmt.Println("Quiz one is scored higher that quiz two")
	} else if quiz1 < quiz2 {
		fmt.Println("Quiz two is scored higher than quiz one")
	} else {
		fmt.Println("Quiz one and quiz two have the same score")
	}

	if average(quiz1, quiz2, quiz3) > 7 {
		fmt.Println("Acceptable grades")
	} else {
		fmt.Println("Instructor did a bad job!")

	}

}
