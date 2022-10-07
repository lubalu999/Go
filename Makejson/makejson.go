package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	person := make(map[string]string, 2)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter a name: ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Printf("Enter an adress: ")
	scanner.Scan()
	adress := scanner.Text()

	person["name"] = name
	person["adress"] = adress

	date, _ := json.Marshal(person)
	fmt.Printf(string(date))
}
