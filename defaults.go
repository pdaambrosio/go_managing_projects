package main

var itemDiscount int = 10 // global variable

// The function "changeDiscount" updates the value of the "itemDiscount" variable and returns the
// updated value.
func changeDiscount(discount int) int {
	itemDiscount = discount
	return itemDiscount
}
