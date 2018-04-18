package main

import (
	"fmt"

)
func main() {
	sum := 0
	a := []int{1,2,3,4,5}
	b := []string{"hello","world"}
	fmt.Printf("slice:%s\n",b[0][1:3])
	for k,v := range a {
		fmt.Printf("k:%d\n",k)
		sum += v
	}
	fmt.Printf("sum : %d\n",sum)
	sum = 0
	for i:=0;i<10;i++ {
		sum += i
	}
	fmt.Printf("sum : %d\n",sum)
}

