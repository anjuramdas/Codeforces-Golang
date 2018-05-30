package main

import (
	"fmt"
)
func main() {
	var n, k, part, output int
	fmt.Scan(&n)
	fmt.Scan(&k)
	if n%2 == 0 {
	    part = n/2
	} else {
	    part = (n/2) +  1
	  }
 
	if (k <= part) {
    	     output = (2*k) - 1
	} else {
    	     output = (k-part) * 2
      }
 
        fmt.Println(output)
}


