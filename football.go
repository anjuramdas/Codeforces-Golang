/* 
# Codeforces
# Problem : A. Football
# URL : http://codeforces.com/contest/96/problem/A
*/
package main
import (
	"fmt"
)
func main () {
	var input string
	fmt.Scan(&input)
	var i, j, count, flag int = 0, 1, 0, 0
	for i < len(input) && j < len(input) {
		if input[i] == input[j] {
			j = j + 1
			count = count + 1
			if count == 6 {
				flag = 1
				break
			} 
		} else {
			i = j
			j = j + 1
			count = 0
		}
	}
	if flag == 1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
