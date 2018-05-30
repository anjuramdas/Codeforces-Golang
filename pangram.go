package main
import (
	"fmt"
	"strings"
)
func main() {
	var n int
	var alpha string
	a := "abcdefghijklmnopqrstuvwxyz"
	f := 0
	fmt.Scan(&n)
	fmt.Scan(&alpha)
	word := strings.ToLower(alpha)
	if n<26 {
		fmt.Println("NO")
	} else {
		for i := 0; i < 26; i++ {
			count := 0
			for j := 0; j < len(word); j++ {
				if a[i] == word[j] {
					count = count + 1
				}
			}
			if count == 0 {
				f = 1
				
			}
		}
		if f == 1 {
		    fmt.Println("NO")
	        } else {
		    fmt.Println("YES")
	          }

	 }


}
