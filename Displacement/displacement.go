package main

import (
	"fmt"
)

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	return func(ti float64) float64 {
		return 0.5*a*ti*ti + vo*ti + so
	}
}

func main() {
	var a, vo, so, t float64
	fmt.Printf("Enter acceleration, a: ")
	fmt.Scanf("%f\n", &a)
	fmt.Printf("Enter initial velocity, vo: ")
	fmt.Scanf("%f\n", &vo)
	fmt.Printf("Enter initial diplacement, so: ")
	fmt.Scanf("%f\n", &so)
	fmt.Printf("Enter time, t: ")
	fmt.Scanf("%f\n", &t)

	fnc := GenDisplaceFn(a, vo, so)
	fmt.Println(fnc(t))
}
