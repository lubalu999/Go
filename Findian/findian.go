package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	myscanner := bufio.NewScanner(os.Stdin) //need for a line with spaces
	myscanner.Scan()
	str := myscanner.Text()

	var new_str string = strings.ToLower(str)
	var leap bool = len(new_str) >= 3 && new_str[0] == 'i' && new_str[len(new_str)-1] == 'n' && strings.Contains(new_str, "a")

	if leap {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}
}
