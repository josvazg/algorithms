package main

import (
	"fmt"
)

func mergesort(list []int) []int {
	if(len(list)==1) {
		return list
	}
	fmt.Println("unordered list=",list)
	for size:=1;size<len(list);size=2*size {
		for i:=0;i<len(list)-size;i+=2*size {
			limit:=i+(2*size)
			if limit>len(list) {
				limit=len(list)
			}
			merge(list[i:limit],size)
		}
	}
	return list
}

func merge(list []int, size int) {
	if size==1 {
		if list[1]<list[0] {
			list[0],list[1]=list[1],list[0]
		}
		return
	}
	tmp:=[]int{}
	a:=list[0:size]
	b:=list[size:]
	for ;len(a)>0 || len(b)>0; {
		if len(a)>0 && len(b)>0 {
			if a[0]<b[0] {
				tmp=append(tmp,a[0])
				a=a[1:]
			} else {
				tmp=append(tmp,b[0])
				b=b[1:]
			}
		} else if len(a)>0 {
			tmp=append(tmp,a...)
			a=[]int{}
		} else if len(b)>0 {
			tmp=append(tmp,b...)
			b=[]int{}
		}
	}
	copy(list,tmp)
}

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
	s:=mergesort([]int{9,6,5,3,1,8,7,2,4})
	fmt.Println("Sorted list:",s)
	fmt.Println("Sorted?",sorted(s))
}