package main

import (
	f"fmt"
	s "strings"
	i "io/ioutil"
	o "os"
)
func main() {
	 if len(o.Args) <= 1 {
 		f.Printf("USAGE : %s <target_filename> \n", o.Args[0])
 		o.Exit(1)
 	}
	fileName := o.Args[1]
	fileBytes, err := i.ReadFile(fileName)
 	if err != nil {
 		f.Println(err)
 		o.Exit(1)
 	}
	sliceData := s.Split(string(fileBytes), "\n") 
	code := `PRINT hewo`
	ss := s.Fields(code)
	f.Println(sliceData)
	keyword := ss[0]
	if(keyword == "PRINT"){
		arg := ss[1]
		f.Println(arg)
	}
}

