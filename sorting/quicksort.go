package main

import (
	"sort"
)

func quicksortInts(list []int) {
	quicksort(sort.IntSlice(list))
}

func quicksort(list sort.Interface) {
	qsort(list, 0, list.Len()-1)
}

func qsort(list sort.Interface, left, right int) {
	if left >= right {
		return
	}
	p := pivot(list, left, right)
	done := partition(list, left, right, p)
	qsort(list, left, done-1)
	qsort(list, done+1, right)
}

func partition(list sort.Interface, left, right, p int) int {
	if p != right { // pivot out of the way
		list.Swap(right, p)
	}
	done := left
	for i := left; i < right; i++ {
		if list.Less(i, right) {
			list.Swap(i, done)
			done++
		}
	}
	list.Swap(right, done)
	return done
}

func pivot(list sort.Interface, start, end int) int {
	p := end
	min := start
	if list.Less(p, start) {
		p, min = start, p
	}
	middle := ((end - start) / 2) + start
	if list.Less(middle, p) {
		p = middle
	}
	if list.Less(p, min) {
		p = min
	}
	return p
}
