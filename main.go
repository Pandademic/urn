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
        	f.Println(scanner.Text())
    	}
    	if err := scanner.Err(); err != nil {
        	log.Fatal(err)
    	}
	code := `PRINT hewo`
	ss := s.Fields(code)
	keyword := ss[0]
	if(keyword == "PRINT"){
		arg := ss[1]
		f.Println(arg)
	}
}

