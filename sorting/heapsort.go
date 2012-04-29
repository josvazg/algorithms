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

func nheapsortInts(list []int) {
	nheapify(list)
	end := len(list) - 1
	for done := end; done > 0; done-- {
		list[0], list[done] = list[done], list[0]
		nheapify(list[:done])
	}
}

func nheapify(list []int) {
	end := len(list) - 1
	start := (end - 1) / 2 // last parent node
	for ; start >= 0; start-- {
		root := start
		for (root*2)+1 <= end { // make parent (root) the biggest/top of the heap
			child := (root * 2) + 1
			swap := root
			if list[swap] < list[child] {
				swap = child
			}
			if child+1 <= end && list[swap] < list[child+1] {
				swap = child + 1
			}
			if swap != root {
				list[root], list[swap] = list[swap], list[root]
				root = swap
			} else {
				break // next (previous) parent...
			}
		}
	}
}

