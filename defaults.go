package main

var itemDiscount int = 10 // global variable

// The function "changeDiscount" updates the value of the "itemDiscount" variable and returns the
// updated value. Is important use the convention of the name of the function, in this case the name of the function is "DefaultsChangeDiscount"
// because Default is the name of the file and ChangeDiscount is the name of the function.
func DefaultsChangeDiscount(discount int) int {
	itemDiscount = discount
	return itemDiscount
}
