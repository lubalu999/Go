package main

import (
	"fmt"
)

func Swap(arr []int, idx int) {
	tmp := arr[idx]
	arr[idx] = arr[idx+1]
	arr[idx+1] = tmp
}

func BubbleSort(sli []int) {
	for i := 0; i < len(sli); i++ {
		flag_swap := false
		for j := 1; j < len(sli); j++ {
			if sli[j] < sli[j-1] {
				Swap(sli, j-1)
				flag_swap = true
			}
		}
		if flag_swap == false {
			break
		}
	}
}

func main() {
	var length int
	mass := make([]int, 0)

	fmt.Printf("Enter a length of array ( <= 10): ")
	fmt.Scanf("%d\n", &length)

	if length <= 10 && length >= 0 {
		for i := 0; i < length; i++ {
			var num int
			fmt.Scanf("%d\n", &num)
			mass = append(mass, num)
		}

		BubbleSort(mass)
		fmt.Printf("%v", mass)
	} else {
		fmt.Printf("You entered the wrong array length!")
	}
}
