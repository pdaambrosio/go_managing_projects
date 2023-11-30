package main

import (
	"go_managing_projects/src/calc"
	"fmt"
)

func init() {
	println("main.go initialized")
}

/*
 The function "calc.Discount" is imported from the "calc" package. The "calc" package is located in the "src" folder.
 The "src" folder is located in the "go_managing_projects" folder. It's important to note that the "go_managing_projects"
 because it's the root folder of the project and I don't use GOPATH, so I need to specify the root folder of the project.
*/

// The main function calculates the total discount based on the item price and discount percentage.
func main() {
	var itemPrice int = 50
	// If you comment this line, the default value of the "itemDiscount" variable will be used.
	itemDiscount = DefaultsChangeDiscount(20) // The function "DefaultsChangeDiscount" is from the "defaults.go" file.
	itemDiscount = DefaultsRandomNumber() // The function "DefaultsRandomNumber" is from the "defaults.go" file.

	totalDiscount := calc.Discount(itemPrice, itemDiscount)
	fmt.Println("Item Price:      ", itemPrice)
	fmt.Println("Item Discount %: ", itemDiscount)
	fmt.Println("Total Discount:  ", totalDiscount)
}
