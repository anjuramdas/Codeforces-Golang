/* 
# Codeforces
# Problem : B. T-primes
# URL : http://codeforces.com/contest/230/problem/B
# A number is a T prime if it is a perfect square and its square root is a prime number.
*/
package main

import (
	"fmt"
	"math"
)
func perfect_square(x float64) bool { // To check if the input is a perfect square or not 
	var y, z float64
	y = math.Sqrt(x)
	z = math.Floor(y)
	if y-z == 0 {
		return true
	} else {
		return false
	}
}
func main() {
	var n, intput, i, j float64
	var psq bool
	var flag int = 0
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&intput)
		if intput == 1 || intput == 2 {
			fmt.Println("NO")
		} else {
			psq = perfect_square(intput)
			if psq == true {
				sq := math.Sqrt(intput)
				for j = 2; j <= (sq / 2); j++ { // Check if the square root of the input is a prime or not 
					if int(intput)%int(j) == 0 {
						flag = 1
						break
					}
				}
				if flag == 1 {
					fmt.Println("NO")
				} else {
					fmt.Println("YES")
				}
			} else {
				fmt.Println("NO")
			}
		}
	}
}
