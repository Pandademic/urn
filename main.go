package main

import (
	"fmt"
	s "strings"
)
func main() {
	p = fmt.Println
	code := 'PRINT hewo'
	ss := strings.Fields(s)
	keyword := ss[0]
	if(keyword == "PRINT"){
		arg := ss[1]
		P(arg)
	}
}

