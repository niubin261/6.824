package main

import (
	"math"
	"fmt"
)

//switch
// statistics the prime number between2-100
func sw(score int)(string) {
	var level string
	switch score / 10 {
	case 10:
		level = "A"// no need for break
	case 9:
		level = "A"
	case 8:
		level = "B"
	case 7:
		level = "C"
	case 6:
		level = "D"
	default:
		level = "F"


	}
	return level

}
func count(i int)(bool)  {
	if i == 2 {
		return true
	}
	if i % 2 == 0 {
		return false
	}
	l := int(math.Sqrt(float64(i)))
	for j := 2;j < l + 1; {
		if ((i % j)) == 0 {

			return false
		}
		j = j + 1
	}

	return true
}
func main() {
	ch := []chan int{}
	for i := range ch {
		ch[i] = make(chan int)
	}
	cnt := 0
	for i :=2;i < 100;i ++ {
		if count(i) {
			cnt ++
		}
		fmt.Printf("isNumber prime: %t\n",count(i))
	}
	fmt.Printf("cnout isNumber %d\n",cnt)
	a := [][]int{{1,2,3,4},{1,2,3,4}}
	a = append(a, []int{7})
	for i := range a{
		for j := range a[i] {
			fmt.Printf("%d\n",a[i][j])
		}
	}

	b := [5][5]int{{1,2,3,4,5},{1,2,3,4,5}}
	for i := range b {
		for j := range b[i] {
			fmt.Printf("%d\n",b[i][j])
		}
	}
	m := map[string]map[string]string{}
	m ["a"] = map[string]string{"hello":"world"}
	for k,_:=range m {
		for kl,vl:=range m[k]{
			fmt.Printf("%s,%s\n",kl,vl)
		}

	}

}