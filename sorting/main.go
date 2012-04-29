package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

var tests = [][]int{
	{9, 6, 5, 3, 1, 8, 7, 2, 4},
	{9, 6, 5, 3, 1, 8, 7, 2, 4, 0},
	{9, 8, 7, 6, 5, 4, 3, 2, 1},
	{9, 9, 9, 2, 3, 4, 2, 4, 5, 7, 6, 8, 3, 5},
	{1, 2, 3},
	{30, 20, 1},
	{25, 0, 50},
	{70, 0, 50},
}

var sorts = []struct {
	name string
	fn   func([]int)
}{
	{"mergesort", mergesort},
	{"swap mergesort", imergesortInts},
	{"quicksort", quicksortInts},
	{"heapsort", heapsortInts},
	{"sort.Ints", sort.Ints},
	{"native in-place mergesort", nimergesortInts},
	{"native quicksort", nquicksortInts},
	{"native heapsort", nheapsortInts},
	{"sortInts",SortInts},
}

var sortstr = []struct {
	name string
	fn   func([]string)
}{
	{"sort.Strings",sort.Strings},
	{"sortStrings",SortStrings},
}

func convert(in [][]int) [][]string {
	out:=make([][]string,len(in))
	for i,test:=range in {
		out[i]=make([]string,len(test))
		for j,item :=range test {
			out[i][j]=strconv.Itoa(item)
		}
	}
	return out
}

func main() {
	for _, u := range tests {
		list := make([]int, len(u))
		//fmt.Println("Unsorted:", u)
		for _, asort := range sorts {
			copy(list, u)
			asort.fn(list)
			//fmt.Println(asort.name, " sorted list:", list)
			if !sort.IntsAreSorted(list) {
				fmt.Println("FAILED! list=", list)
				os.Exit(-1)
			}
		}
	}
	fmt.Println("PASSED")
	times := 50000
	for _, asort := range sorts {
		t := time.Now()
		for i := 0; i < times; i++ {
			for _, u := range tests {
				list := make([]int, len(u))
				copy(list, u)
				asort.fn(list)
			}
		}
		fmt.Println(asort.name, "sorted", times, "times in", time.Since(t))
	}
	fmt.Println("Strings tests")
	teststr:=convert(tests)
	for _, u := range teststr {
		list := make([]string, len(u))
		//fmt.Println("Unsorted:", u)
		for _, asort := range sortstr {
			copy(list, u)
			asort.fn(list)
			//fmt.Println(asort.name, " sorted list:", list)
			if !sort.StringsAreSorted(list) {
				fmt.Println("FAILED!")
				os.Exit(-1)
			}
		}
	}
	fmt.Println("PASSED")
	for _, asort := range sortstr {
		t := time.Now()
		for i := 0; i < times; i++ {
			for _, u := range teststr {
				list := make([]string, len(u))
				copy(list, u)
				asort.fn(list)
			}
		}
		fmt.Println(asort.name, " sorted", times, "times in", time.Since(t))
	}
}
