package main

import( 
	"fmt"
	"strconv"
	"strings"		
)

func main(){
	var n int
	var word,l,out string
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&word)
		length := len(word)
		if length <= 10 {
			fmt.Println(word)
		} else {
			var output []string
			l = strconv.FormatInt(int64(length-2), 10)
			output = append(output,word[:1])
			output = append(output,l)
			output = append(output,word[length-1:length])
			out = strings.Join(output, "")
			fmt.Println(out)
		 }
	}
}