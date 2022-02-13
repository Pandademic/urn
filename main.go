package main

import (
	f"fmt"
	s "strings"
	i "io/ioutil"
	o "os"
)
func main() {
	 if len(os.Args) <= 1 {
 		fmt.Printf("USAGE : %s <target_filename> \n", os.Args[0])
 		os.Exit(1)
 	}
	fileName := os.Args[1]
	fileBytes, err := i.ReadFile(fileName)
 	if err != nil {
 		fmt.Println(err)
 		os.Exit(1)
 	}
	sliceData := strings.Split(string(fileBytes), "\n") 
	code := `PRINT hewo`
	ss := s.Fields(code)
	fmt.Println(sliceData)
	keyword := ss[0]
	if(keyword == "PRINT"){
		arg := ss[1]
		f.Println(arg)
	}
}

