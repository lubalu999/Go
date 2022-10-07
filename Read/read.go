package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const maxLength = 20

type Person struct {
	fname string
	lname string
}

func main() {
	var fileName string
	sli := make([]Person, 0, 1)
	fmt.Printf("Enter file name: ")
	fmt.Scanf("%s", &fileName)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text()) //split fname and lname from one line
		person := Person{words[0], words[1]}

		if len(person.fname) > maxLength { //if the length of fname or lname > 20
			person.fname = words[0][:maxLength]
		}
		if len(person.lname) > maxLength {
			person.lname = words[1][:maxLength]
		}
		sli = append(sli, person)
	}

	for _, name := range sli {
		fmt.Println(name.fname, name.lname)
	}
	file.Close()
}
