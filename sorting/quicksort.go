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

func nquicksortInts(list []int) {
	if len(list) <= 1 {
		return
	}
	p := npivot(list)
	done := npartition(list, p)
	nquicksortInts(list[0:done])
	nquicksortInts(list[done+1:])
}

func npartition(list []int, p int) int {
	end := len(list) - 1
	if p != end { // pivot out of the way (to the end)
		list[end], list[p] = list[p], list[end]
	}
	done := 0
	for i := 0; i < end; i++ {
		if list[i] < list[end] {
			list[i], list[done] = list[done], list[i]
			done++
		}
	}
	list[end], list[done] = list[done], list[end]
	return done
}

func npivot(list []int) int {
	end := len(list) - 1
	min := 0
	p := end
	if list[p] < list[0] {
		p, min = 0, p
	}
	middle := (end / 2)
	if list[middle] < list[p] {
		p = middle
	}
	if list[p] < list[min] {
		return min
	}
	return p
}

