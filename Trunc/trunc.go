package main

import (
	"fmt"
)

func main() {
	var number float64
	fmt.Scanf("%f", &number)

	var tmp int64 = int64(number)
	fmt.Printf("%d", tmp)
}
