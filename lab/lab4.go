package main

import (
	"fmt"
	"regexp"
	"strings"


)

func split2(s string) ([] string)  {
	words := []string{}
	i := 0
	j := 0
	l :=len(s)
	for i < l && j < l {

		for i < l && (s[i] == ' ' || s[i] == '\n') {
			i ++
		}
		j = i
		for j < l && s[j] != ' ' && s[j] != '\n' {
			j ++

		}
		fmt.Printf("i : %d j : %d slice %s\n",i,j,s[i:j])
		if j != i {

			words = append(words, s[i:j])
		}
		i = j


	}
	return words

}
func split(s string)([] string){
	words := []string{}
	i := 0
	j := 0
	l := len(s)
	for i< l {
		if s[i] == ' ' {
			i ++
		} else {
			j = i
			for s[j] != ' ' {
				j ++
			}

			i = j
		}

		words = append(words, s[i:j])
	}
	return words
}
func main() {
	const val string = " hello   world\n helloworld   "
	s := "hello, hello world. niubin"
	m := map[string]int{}
	re := regexp.MustCompile(`[a-zA-Z]+`)
	words := re.FindAllString(s,-1)

	for _,val := range words{
		fmt.Printf("%s\n",val)
	}

	for _,key := range words {

		_,found := m[key]
		if !found {
			m[key] = 1
		} else {
			m[key] = m[key] + 1
		}

	}
	for k,v := range m {
		fmt.Printf("%s = %d\n",k,v)
	}

	words = strings.Split(s," ")

	for _,val := range words{
		fmt.Printf("%s\n",val)
	}



	words = split2(val)
	for _,val := range words{
		fmt.Printf("%s\n",val)
	}




}