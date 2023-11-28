package calc

// The Plus function in Go takes two integers as input and returns their sum.
func Plus(a, b int) int {
	return a + b
}

// The Minus function subtracts two integers and returns the result.
func Minus(a, b int) int {
	return a - b
}

// The Multi function in Go multiplies two integers and returns the result.
func Multi(a, b int) int {
	return a * b
}

// The function "Devide" divides two integers and returns the result.
func Devide(a, b int) int {
	return a / b
}

// The Discount function calculates the discount amount based on the given percentage and value.
func Discount(a, b int) float64 {
	return float64(a * b) / 100
}
