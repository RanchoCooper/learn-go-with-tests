package integers

import (
	"fmt"
)

// Add takes two integers and returns the sum of them
func Add(x, y int) int {
	return x + y
}

func ExampleAdd() {
	sum := Add(2, 8)
	fmt.Println(sum)
	// Output: 10
}
