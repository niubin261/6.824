package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s1 string
	var s2 string
	s1 = "hello"
	s2 = "world"
	fmt.Printf("%s\n",s1)
	fmt.Printf("%s\n",s2)
	s := s1 + " " + s2
	fmt.Printf("%s\n",s)
	s = "123456"
	result,_ := strconv.Atoi(s)
	fmt.Printf("%d\n",result)
	fmt.Printf("%s\n",strconv.Itoa(123))

}