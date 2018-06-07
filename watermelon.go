package main
import (
	"fmt"
)
func main() {
	var weight int
	fmt.Scan(&weight)
	if weight % 2 == 0 {
		if weight == 2 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	} else {
		fmt.Println("NO")
	}
}
