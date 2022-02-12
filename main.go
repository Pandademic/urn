package main

import (
	"fmt"
	s "strings"
)
func main() {
	code := "PRINT hewo"
	ss := strings.Fields(s)
	keyword := ss[0]
	if(keyword == "PRINT"){
		arg := ss[1]
		p = fmt.Println
		p(arg)
	}
}

