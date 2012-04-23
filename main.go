package main

import (
	"fmt"
	"sort"
	"time"
)

func sorted(list []int) bool {
	if len(list) <= 1 {
		return true
	}
	last := list[0]
	for i := 1; i < len(list); i++ {
		if list[i] < last {
			return false
		}
		last = list[i]
	}
	return true
}

var tests = [][]int{
	{9, 6, 5, 3, 1, 8, 7, 2, 4},
	{9, 6, 5, 3, 1, 8, 7, 2, 4, 0},
	{9, 8, 7, 6, 5, 4, 3, 2, 1},
	{9, 9, 9, 2, 3, 4, 2, 4, 5, 7, 6, 8, 3, 5},
}

var sorts = []struct {
	name string
	fn   func([]int)
}{
	{"mergesort", mergesort},
	{"swap mergesort", smergesort},
	{"sort.Ints", sort.Ints},
}

func main() {
	for _, u := range tests {
		list := make([]int, len(u))
		fmt.Println("Unsorted:", u)
		for _, sort := range sorts {
			copy(list, u)
			sort.fn(list)
			fmt.Println(sort.name, " sorted list:", list)
			if !sorted(list) {
				fmt.Println("FAILED!")
				break
			}
		}
	}
	times := 500000
	for _, sort := range sorts {
		t := time.Now()
		for i := 0; i < times; i++ {
			for _, u := range tests {
				list := make([]int, len(u))
				copy(list, u)
				sort.fn(list)
			}
		}
		fmt.Println(sort.name, " sorted", times, "times in", time.Since(t))
	}
}

