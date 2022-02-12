package main

import (
	f"fmt"
	s "strings"
)
func main() {
	code := "PRINT hewo"
	ss := s.Fields(code)
	keyword := ss[0]
	if(keyword == "PRINT"){
		arg := ss[1]
		f.Println(arg)
	}
}

