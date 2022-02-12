package main

import (
	"fmt"
	s "strings"
)
p = fmt.Println
func main() {
	code := "PRINT hewo"
	ss := strings.Fields(s)
	keyword := ss[0]
	if(keyword == "PRINT"){
		arg :== ss[1]
		P(arg)
	}
}

