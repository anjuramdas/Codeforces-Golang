package main
import (
	"fmt"
)
func main() {
	var s int
	fmt.Scan(&s)
    array1 := []string{"zero","one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	array2 := []string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	if 0 <= s && s <= 19 {
		fmt.Println(array1[s])
	} else if s == 20 || s == 30 || s == 40 || s == 50 || s == 60 || s == 70 || s == 80 || s == 90 {
		fmt.Println(array2[(s/10)-2])
	} else {
		fmt.Print(array2[(s/10)-2])
		fmt.Print("-")
		fmt.Println(array1[s%10])
	} 
}