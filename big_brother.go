package main

import (
	"fmt"	
)
func main() {
	var a, b int
	fmt.Scan(&a,&b)
	for i := 1; i < 1000; i++ {
		a = 3 * a
		b = 2 * b
		if a > b {
			fmt.Println(i)
			break
		}
	}
}	 