package main
import (
	"fmt"
)
func main() {
	var input, output int
	fmt.Scan(&input)
	if input % 2 == 0 {
		output = input / 2
	} else {
		output = ((input + 1) / 2) * (-1)
	}
	fmt.Println(output)
}	