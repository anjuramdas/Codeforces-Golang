package main
import (
	"fmt"
	"strings"
)
func main() {
	var n int //the number of games played
	var input string //the outcome of each of the games
	fmt.Scan(&n, &input)
	A := strings.Count(input,"A")
	D := strings.Count(input,"D")
	if A > D {
		fmt.Println("Anton")
	} else if A < D {
		fmt.Println("Danik")
	} else {
		fmt.Println("Friendship")
	}
}	
