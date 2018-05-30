package main

import(
	
	"fmt"
	"strconv"
	
)
func main(){
	var input_year, output_year int
	var j string
	fmt.Scan(&input_year)
	output_year = input_year + 1
	for output_year != 0 {
		j = strconv.FormatInt(int64(output_year), 10)
		if j[0] != j[1] && j[0] != j[2] && j[0] != j[3] && j[1] != j[2] && j[2] != j[3] && j[1] != j[3]{
			fmt.Println(output_year)
			break			
		} else {
		  	output_year = output_year + 1
		  	
		  }	
	}
}	
