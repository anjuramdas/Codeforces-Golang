package main
import (
	"fmt"
	"strconv"
)	
func main() {
	var n int
	var s string = "" 
	fmt.Scan(&n)
	if n <= 9 {
		fmt.Println(n)
	} else {
		for i := 1; i <= n; i++ {
			s = s + strconv.Itoa(i)
		}
		fmt.Println(string(s[n-1]))
	}
}