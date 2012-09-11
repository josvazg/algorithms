package main

import (
	"fmt"
	"math"
)

const (
	second=1000000
	minute=60*second
	hour=60*minute
	day=24*hour
	month=30*day
	year=365*day
	century=100*year
)

type solver func (float64)string

var funcNames=[]string{"log(n)","sqrt(n)","n","nlog(n)","n²","n³","2^n","n!"}
var funcs=[]solver{logn,sqrtn,n,nlogn,n2,n3,_2n,factn}
var durations=[]float64{second,minute,hour,day,month,year,century}

func logn(t float64) string {
	return fmt.Sprintf("2^%.2e",t)
}

func sqrtn(t float64) string {
	exp:=2*math.Log10(t)
	fexp:=math.Floor(exp)
	if fexp<exp {	
		mul:=math.Pow(10,exp-fexp)
		return fmt.Sprintf("%.2fx10^%d",mul,int(fexp))
	}
	return fmt.Sprintf("10^%d",int(fexp))
}

func n(t float64) string {
	return fmt.Sprintf("%.2e",t)	
}

func nlogn(t float64) string {
	i:=0
	n:=t/2
	nlogn:=float64(n*math.Log2(n))
	sign:=t-nlogn
	min:=float64(0)
	diff:=t-n
	for diff!=0 && i<100 {
		n,min,diff=nextTry(n,sign,min)
		nlogn=n*math.Log2(n)
		sign=t-nlogn		
		i++
	}	
	if n>99999999 {
		return fmt.Sprintf("%.2e",n-2)	
	}	
	return fmt.Sprintf("%.0f",n)
}

func nextTry(n, sign, min float64) (next, nmin, diff float64) {
	delta:=math.Floor((n-min)/2)
	if sign>0 {
		return n+delta,n, +delta
	}
	return n-delta,min, -delta
}

func n2(t float64) string {
	return fmt.Sprintf("%.0f",math.Sqrt(t))
}

func n3(t float64) string {
	return fmt.Sprintf("%.0f",math.Cbrt(t))
}

func _2n(t float64) string {
	return fmt.Sprintf("%.0f",math.Log2(t))
}

func factn(t float64) string {
	n:=float64(1)
	fact:=float64(n)
	for ;fact<t;n++ {
		fact=fact*n
	}
	return fmt.Sprintf("%.0f",n-2)
}

func main() {
	fmt.Println("Function   1second     1minute     1hour       1day        1month      1year       1century")
	fmt.Println("________   __________  __________  __________  __________  __________  __________  __________")
	
	for i,f:=range funcs {
		fmt.Printf("%8s   ",funcNames[i])
		for _,d := range durations {
			fmt.Printf("%10s  ",f(d))
		}
		fmt.Println();	
	}
}

