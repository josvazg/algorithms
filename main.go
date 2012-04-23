package main

import (
	"fmt"
	"sort"
	"time"
)

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
		for _, asort := range sorts {
			copy(list, u)
			asort.fn(list)
			fmt.Println(asort.name, " sorted list:", list)
			if !sort.IntsAreSorted(list) {
				fmt.Println("FAILED!")
				break
			}
		}
	}
	times := 500000
	for _, asort := range sorts {
		t := time.Now()
		for i := 0; i < times; i++ {
			for _, u := range tests {
				list := make([]int, len(u))
				copy(list, u)
				asort.fn(list)
			}
		}
		fmt.Println(asort.name, " sorted", times, "times in", time.Since(t))
	}
}

