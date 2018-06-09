package main 
import (
	"fmt"
)
func main() {
	var input, output int = 0, 0
	fmt.Scan(&input)
	if input == 1 || input == 2 || input == 3 || input == 4 || input == 5 {
		fmt.Println(1)
	} else {
		for input > 0 {
			if input - 5 >= 0 {
				input = input - 5
				output = output + 1
			} else if input - 4 >= 0 {
				input = input - 4
				output = output + 1
			} else if input - 3 >= 0 {
				input = input - 3
				output = output + 1
			} else if input - 2 >= 0 {
				input = input - 2
				output = output + 1
			} else if input - 1 >= 0 {
				input = input - 1
				output = output + 1
			} 
		}
	fmt.Println(output)
	}
}	







