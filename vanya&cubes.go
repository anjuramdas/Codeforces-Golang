package main 
import (
	"fmt"
)
func main() {
	var cubes, height, sum, i int = 0, 0, 0, 1
	fmt.Scan(&cubes)
	for true {
		sum = sum + i
		cubes = cubes - sum
		i = i + 1
		if cubes >= 0 {
			height = height + 1
		} else {
			break
		}
	}
	fmt.Println(height)
}	