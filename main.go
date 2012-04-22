package main

import (
	"fmt"
)

func sorted(list []int) bool {
	if len(list)<=1 {
		return true
	}
	last:=list[0]
	for i:=1;i<len(list);i++ {
		if list[i]<last {
			return false
		}
		last=list[i]
	}
	return true
}

func main() {
	u:=[]int{9,6,5,3,1,8,7,2,4}
	fmt.Println("Unsorted:",u)
	s:=mergesort(u)
	fmt.Println("Sorted list:",s)
	fmt.Println("Sorted?",sorted(s))
}