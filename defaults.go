package main

import (
	"math/rand"
)

var itemDiscount int = 10 // global variable

/*
The function "changeDiscount" updates the value of the "itemDiscount" variable and returns the
updated value. Is important use the convention of the name of the function, in this case the name of the function is "DefaultsChangeDiscount"
because Default is the name of the file and ChangeDiscount is the name of the function.
*/
func DefaultsChangeDiscount(discount int) int {
	itemDiscount = discount
	return itemDiscount
}

// The function randomNumber generates a random number between 0 and 100.
func randomNumber() int {
	min := 0
	max := 100
	itemDiscount := rand.Intn(max-min+1) + min
	return itemDiscount
}

func DefaultsRandomNumber() int {
	return randomNumber()
}
