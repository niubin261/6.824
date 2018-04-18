package main
import "fmt"
func  sum(a []int,result chan <- int)(){
	sum := 0
	for _,v := range a{
		sum += v
	}
	result <- sum
}
func swap(s1 string,s2 string)(string,string)  {

	return s2,s1
}
func  main(){
	a :=[]int{1,2,3,4,5,6}
	result := make(chan int)
	go sum(a[:3],result)
	go sum(a[3:],result)
	x,y := <-result,<-result
	fmt.Printf("x:%d\n",x)
	fmt.Printf("sum:%d\n",x+y)
	s1 := "hello"
	s2 := "world"
	s1,s2 = swap(s1,s2)
	fmt.Printf("%s + %s \n",s1,s2)
}
