package main

import (
	"sort"
)

func heapsortInts(list []int) {
	heapsort(sort.IntSlice(list))
}

func heapsort(list sort.Interface) {
	heapify(list, list.Len())
	for done := list.Len() - 1; done > 0; done-- {
		list.Swap(0, done)
		heapify(list, done)
	}
}

func heapify(list sort.Interface, count int) {
	end := count - 1
	start := (count - 2) / 2 // last parent node
	for ; start >= 0; start-- {
		root := start
		for (root*2)+1 <= end { // make parent (root) the biggest/top of the heap
			child := (root * 2) + 1
			swap := root
			if list.Less(swap, child) {
				swap = child
			}
			if child+1 <= end && list.Less(swap, child+1) {
				swap = child + 1
			}
			if swap != root {
				list.Swap(root, swap)
				root = swap
			} else {
				break // next (previous) parent...
			}
		}
	}
}

