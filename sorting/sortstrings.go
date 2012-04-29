package main

func SortStrings(data []string) {
	// Switch to heapsort if depth of 2*ceil(lg(n+1)) is reached.
	n := len(data)
	maxDepth := 0
	for i := n; i > 0; i >>= 1 {
		maxDepth++
	}
	maxDepth *= 2
	quickSortStrings(data, 0, n, maxDepth)
}

func quickSortStrings(data []string, a, b, maxDepth int) {
	for b-a > 7 {
		if maxDepth == 0 {
			heapSortStrings(data, a, b)
			return
		}
		maxDepth--
		mlo, mhi := doPivotStrings(data, a, b)
		// Avoiding recursion on the larger subproblem guarantees
		// a stack depth of at most lg(b-a).
		if mlo-a < b-mhi {
			quickSortStrings(data, a, mlo, maxDepth)
			a = mhi // i.e., quickSort(data, mhi, b)
		} else {
			quickSortStrings(data, mhi, b, maxDepth)
			b = mlo // i.e., quickSort(data, a, mlo)
		}
	}
	if b-a > 1 {
		insertionSortStrings(data, a, b)
	}
}

func doPivotStrings(data []string, lo, hi int) (midlo, midhi int) {
   		m := lo + (hi-lo)/2 // Written like this to avoid integer overflow.
   		if hi-lo > 40 {
   			// Tukey's ``Ninther,'' median of three medians of three.
   			s := (hi - lo) / 8
   			medianOfThreeStrings(data, lo, lo+s, lo+2*s)
   			medianOfThreeStrings(data, m, m-s, m+s)
   			medianOfThreeStrings(data, hi-1, hi-1-s, hi-1-2*s)
   		}
   		medianOfThreeStrings(data, lo, m, hi-1)
   	
   		// Invariants are:
   		//	data[lo] = pivot (set up by ChoosePivot)
   		//	data[lo <= i < a] = pivot
   		//	data[a <= i < b] < pivot
   		//	data[b <= i < c] is unexamined
   		//	data[c <= i < d] > pivot
   		//	data[d <= i < hi] = pivot
   		//
   		// Once b meets c, can swap the "= pivot" sections
   		// into the middle of the slice.
   		pivot := lo
   		a, b, c, d := lo+1, lo+1, hi, hi
   		for b < c {
   			if data[b]<data[pivot] { // data[b] < pivot
   				b++
   				continue
   			}
   			if !(data[pivot]<data[b]) { // data[b] = pivot
   				data[a],data[b]=data[b],data[a]
   				a++
   				b++
   				continue
   			}
   			if data[pivot]<data[c-1] { // data[c-1] > pivot
   				c--
   				continue
   			}
   			if !(data[c-1]<data[pivot]) { // data[c-1] = pivot
   				data[c-1],data[d-1]=data[d-1],data[c-1]
   				c--
   				d--
   				continue
   			}
   			// data[b] > pivot; data[c-1] < pivot
   			data[b],data[c-1]=data[c-1],data[b]
   			b++
   			c--
   		}
   	
   		n := min(b-a, a-lo)
   		swapRangeStrings(data, lo, b-n, n)
   	
   		n = min(hi-d, d-c)
   		swapRangeStrings(data, c, hi-n, n)
   	
   		return lo + b - a, hi - (d - c)
   	}


// medianOfThree moves the median of the three values data[a], data[b], data[c] into data[a].
func medianOfThreeStrings(data []string, a, b, c int) {
	m0 := b
	m1 := a
	m2 := c
	// bubble sort on 3 elements
	if data[m1]<data[m0] {
		data[m1],data[m0]=data[m0],data[m1]
	}
	if data[m2]<data[m1] {
		data[m2],data[m1]=data[m1],data[m2]
	}
	if data[m1]<data[m0] {
		data[m1],data[m0]=data[m0],data[m1]
	}
	// now data[m0] <= data[m1] <= data[m2]
}

func swapRangeStrings(data []string, a, b, n int) {
	for i := 0; i < n; i++ {
		data[a+i],data[b+i]=data[b+i],data[a+i]
	}
}

func heapSortStrings(data []string, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDownStrings(data, i, hi, first)
	}
	
	// Pop elements, largest first, into end of data.
	for i := hi - 1; i >= 0; i-- {
		data[first],data[first+i]=data[first+i],data[first]
		siftDownStrings(data, lo, i, first)
	}
}

// siftDownInts implements the heap property on data[lo, hi).
// first is an offset into the array where the root of the heap lies.
func siftDownStrings(data []string, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && data[first+child]<data[first+child+1] {
			child++
		}
		if !(data[first+root]<data[first+child]) {
			return
		}
		data[first+root],data[first+child]=data[first+child],data[first+root]
		root = child
	}
}

// Insertion sort
func insertionSortStrings(data []string, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && data[j]<data[j-1]; j-- {
			data[j],data[j-1]=data[j-1],data[j]
		}
	}
}

