package main
import (
	"fmt"
)

func main() {
	var a, b, k int
	var answer int = 0
	fmt.Scan(&a, &b)
	answer = a
	for a >= b {
		k = a % b
		a = a / b
		answer = answer + a 
		a = a + k
		
	}
	fmt.Println(answer)
}

