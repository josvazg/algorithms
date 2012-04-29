package main

import (
	"sort"
)

func mergesort(list []int) {
	if len(list) == 1 {
		return
	}
	for size := 1; size < len(list); size += size {
		for i := 0; i < len(list)-size; i += (size + size) {
			limit := i + size + size
			if limit > len(list) {
				limit = len(list)
			}
			merge(list[i:limit], size)
		}
	}
}

func merge(list []int, size int) {
	if size == 1 {
		if list[1] < list[0] {
			list[0], list[1] = list[1], list[0]
		}
		return
	}
	tmp := []int{}
	a := list[0:size]
	b := list[size:]
	for len(a) > 0 || len(b) > 0 {
		if len(a) > 0 && len(b) > 0 {
			if a[0] < b[0] {
				tmp = append(tmp, a[0])
				a = a[1:]
			} else {
				tmp = append(tmp, b[0])
				b = b[1:]
			}
		} else if len(a) > 0 {
			tmp = append(tmp, a...)
			a = []int{}
		} else if len(b) > 0 {
			tmp = append(tmp, b...)
			b = []int{}
		}
	}
	copy(list, tmp)
}

func smergesortInts(list []int) {
	smergesort(sort.IntSlice(list))
}

func smergesort(list sort.Interface) {
	if list.Len() == 1 {
		return
	}
	for size := 1; size < list.Len(); size += size {
		for i := 0; i < list.Len()-size; i += (size + size) {
			limit := i + size + size
			if limit > list.Len() {
				limit = list.Len()
			}
			swapMerge(list, i, limit, size)
		}
	}
}

func swapMerge(list sort.Interface, start, end, size int) {
	o := start
	a := start + size
	for ; o < (start + size); o++ {
		if o == a {
			a++
		}
		if list.Less(a, o) {
			list.Swap(a, o)
		}
		for b := a; b < end-1 && list.Less(b+1, b); b++ {
			list.Swap(b, b+1)
		}
	}
}
