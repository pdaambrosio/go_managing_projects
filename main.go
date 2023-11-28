package main

import (
	"fmt"
)

/*
 The function "Discount" isn't imported from the calc package, but it's available to the main function.
 It's because both the functions are in the same package "main" and the same directory.
*/

// The main function calculates the total discount based on the item price and discount percentage.
func main() {
	var itemPrice, itemDiscount int = 50, 10

	totalDiscount := Discount(itemPrice, itemDiscount)
	fmt.Println("Item Price:      ", itemPrice)
	fmt.Println("Item Discount %: ", itemDiscount)
	fmt.Println("Total Discount:  ", totalDiscount)
}
