package main

import (
	f"fmt"
	s "strings"
	"os"
	"log"
	"bufio"
)
func main() {
	 file, err := os.Open(os.Args[1])
    	if err != nil {
        	log.Fatal(err)
    	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		code := scanner.Text()
		ss := s.Fields(code)
		keyword := ss[0]
		if(keyword == "PRINT"){
			arg := ss[1]
			f.Println(arg)
		}
		if(keyword == "SET"){
			name := ss[1]
			val := ss[5]
			valtype := ss[2]
			if valType == "int"{
				var str(name) int = int(val)
			}
		}
    	}
    	if err := scanner.Err(); err != nil {
        	log.Fatal(err)
    	}
}

