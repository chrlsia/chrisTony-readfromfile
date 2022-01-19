package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fileHandler, _ := os.Open("quotes.txt")
	defer fileHandler.Close()

	lineScanner := bufio.NewScanner(fileHandler)

	counter := 0
	for lineScanner.Scan() {
		fmt.Println(lineScanner.Text())		
		counter++
		if counter == 3 {
			break
		}
	} 
}