//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price int
}

func displayDetails(shoppingList [4]Product) {
	fmt.Println("===> Full Shopping List ===>:", shoppingList)

	var sum, totalItems int

	for i := 0; i < len(shoppingList); i++ {
		item := shoppingList[i]
		sum += item.price
		if item.name != "" {
			totalItems += 1
		}
	}

	fmt.Println("Sum total price of items:", sum)
	fmt.Println("The total number of items:", totalItems)
	fmt.Println("The last item is:", shoppingList[totalItems-1])

}

func main() {

	var shoppingList [4]Product

	shoppingList[0] = Product{"Shoes", 5}
	shoppingList[1] = Product{"TV", 45}
	shoppingList[2] = Product{"Umbrella", 2}

	displayDetails(shoppingList)

	shoppingList[3] = Product{"Phone", 67}
	displayDetails(shoppingList)

}
