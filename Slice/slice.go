package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sli := make([]int, 0, 3)
	for {
		var date string
		fmt.Scanf("%s", &date)

		if date == "X" {
			break
		}
		num, err := strconv.Atoi(date)
		if err == nil {
			sli = append(sli, num)
			sort.Ints(sli)
			fmt.Println(sli)
		}
	}
}
