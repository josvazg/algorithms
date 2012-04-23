package main

func mergesort(list []int) {
	if len(list) == 1 {
		return
	}
	for size := 1; size < len(list); size = 2 * size {
		for i := 0; i < len(list)-size; i += 2 * size {
			limit := i + (2 * size)
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

func smergesort(list []int) {
	if len(list) == 1 {
		return
	}
	for size := 1; size < len(list); size += size {
		for i := 0; i < len(list)-size; i += (size + size) {
			limit := i + (2 * size)
			if limit > len(list) {
				limit = len(list)
			}
			swapOnlyMerge(list[i:limit], size)
		}
	}
}

func swapOnlyMerge(list []int, size int) {
	o := 0
	a := size
	for ; o < size; o++ {
		if o == a {
			a++
		}
		if list[a] < list[o] {
			list[a], list[o] = list[o], list[a]
		}
		for b := a; b < len(list)-1 && list[b+1] < list[b]; b++ {
			list[b], list[b+1] = list[b+1], list[b]
		}
	}
}

