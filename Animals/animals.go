package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal Animal) Eat() string {
	return animal.food
}

func (animal Animal) Move() string {
	return animal.locomotion
}

func (animal Animal) Speak() string {
	return animal.noise
}

func getOption(animal Animal, option string) string {
	var result string
	switch option {
	case "eat":
		result = animal.Eat()
	case "move":
		result = animal.Move()
	case "speak":
		result = animal.Speak()
	}
	return result
}

func main() {
	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		var result string
		fmt.Printf("Enter animal (cow, bird, snake) and option (eat, move, speak): ")
		scanner.Scan()
		lineInput := scanner.Text()

		words := strings.Fields(lineInput)
		animal := words[0]
		option := words[1]

		switch animal {
		case "cow":
			result = getOption(cow, option)
		case "bird":
			result = getOption(bird, option)
		case "snake":
			result = getOption(snake, option)
		}

		fmt.Println(result)
	}
}
