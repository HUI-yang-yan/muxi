package main
import "fmt"
func main() {
	// Declare a variable with an initial value
	var x int = 10

	// Use a for loop to iterate from 0 to 9
	for i := 0; i < 10; i++ {
		// Print the current value of x and i
		fmt.Printf("x: %d, i: %d\n", x, i)
		
		// Increment x by 1 in each iteration
		x++
	}
}