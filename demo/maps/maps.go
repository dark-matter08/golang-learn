package main

import "fmt"

func main() {

	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 11
	shoppingList["milk"] = 3
	shoppingList["bread"] += 1

	shoppingList["eggs"] += 1

	fmt.Println(shoppingList)

	delete(shoppingList, "milk")
	fmt.Println("Milk deleted, new list: ", shoppingList)

	fmt.Println("need", shoppingList["eggs"], "eggs")

	cereal, found := shoppingList["cereal"]

	if !found {
		println("Nope")
	} else {
		println("Yup", cereal, "boxes")
	}

	var totalItems int

	for _, amount := range shoppingList {
		totalItems += amount
	}

	fmt.Println("Total Items:", totalItems)

}
