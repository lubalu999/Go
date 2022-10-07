package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}
type Bird struct{}
type Snake struct{}

func (Cow) Eat() {
	fmt.Println("grass")
}
func (Cow) Move() {
	fmt.Println("walk")
}
func (Cow) Speak() {
	fmt.Println("moo")
}

func (Bird) Eat() {
	fmt.Println("worms")
}
func (Bird) Move() {
	fmt.Println("fly")
}
func (Bird) Speak() {
	fmt.Println("peep")
}

func (Snake) Eat() {
	fmt.Println("mice")
}
func (Snake) Move() {
	fmt.Println("slither")
}
func (Snake) Speak() {
	fmt.Println("hsss")
}

func createAnimal(array map[string]Animal, name string, kind string) {
	switch kind {
	case "cow":
		array[name] = Cow{}
	case "bird":
		array[name] = Bird{}
	case "snake":
		array[name] = Snake{}
	}
}

func getQuery(array map[string]Animal, name string, option string) {
	switch option {
	case "eat":
		array[name].Eat()
	case "move":
		array[name].Move()
	case "speak":
		array[name].Speak()
	}
}

func main() {
	data := make(map[string]Animal)
	var opt1 string
	var opt2 string
	var opt3 string

	for {
		fmt.Printf("> ")
		len, err := fmt.Scanf("%s %s %s\n", &opt1, &opt2, &opt3)

		if err != nil {
			fmt.Println(err)
			return
		}

		if len == 3 {
			switch opt1 {
			case "newanimal":
				createAnimal(data, opt2, opt3)
				fmt.Println("Created it!")
			case "query":
				if _, ok := data[opt2]; ok { //if the name is contained in the map
					getQuery(data, opt2, opt3)
				}
			}
		}
	}
}
