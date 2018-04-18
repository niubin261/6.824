package main

import (
	"regexp"
	"fmt"
	"container/list"
	"strconv"
)

func isNumber(word string) bool {
	re,err := regexp.Compile(`[0-9]+`);

	if err != nil {
		return false
	}
	ok := re.MatchString(word)
	return ok
}
func isParentheses(word string) (bool,bool) {
	return word == "(" , word == ")"
}
func isOperator(word string) bool {
	re,err := regexp.Compile(`(\+|\-|\*|\/){1}`)
	if err != nil {
		return false
	}
	ok := re.MatchString(word)
	return ok
}
func priority(word1,word2 string) bool {
	if (word1 == "*" || word1 == "/") && (word2 == "+" || word2 == "-") {
		return true
	}
	return false

}
func suffixExp(words []string) ([]string) {

	l1 := list.New()
	l2 := list.New()

	for _,word := range words{

		if isNumber(word){
			l2.PushBack(word)
		}
		if isOperator(word){
			if l1.Back() == nil {
				l1.PushBack(word)
			} else if e := l1.Back();e != nil && isOperator(e.Value.(string)) {
				if  !priority(word,e.Value.(string)) {
					l2.PushBack(e.Value.(string))
					l1.Remove(e)
					l1.PushBack(word)
				}
			} else {
				l1.PushBack(word)
			}
		}

		left,right := isParentheses(word)

		if left {
			l1.PushBack(word)
		}
		if right {
			var e *list.Element
			for e = l1.Back(); e != nil && e.Value.(string) != "("; e = e.Prev(){
				l2.PushBack(e.Value.(string))
				l1.Remove(e)
			}
			e = l1.Back()
			if e != nil && e.Value.(string) == "("  {
				l1.Remove(e)
			}
		}
	}
	for l := l1.Back();l != nil; l = l.Prev() {
		l2.PushBack(l.Value.(string))
	}
	var ret []string
	for l := l2.Back();l != nil; l = l.Prev() {
		ret = append(ret, l.Value.(string))
	}
	return ret
}
func cal(words []string) int {
	l := list.New()
	var result int
	le := len(words)
	for i := le;i > 0; i -- {
		word := words[i - 1]

		if isNumber(word) {
			var v int
			v,_ = strconv.Atoi(word)
			l.PushBack(v)
		}
		if isOperator(word) {
			f := l.Back()
			fv := l.Back().Value.(int)
			l.Remove(f)
			t := l.Back()
			tv := t.Value.(int)
			l.Remove(t)
			switch word {
			case "+":
				result = fv + tv
			case "-":
				result = tv - fv
			case "*":
				result = fv * tv
			case "/":
				result = tv / fv
			}
			l.PushBack(result)
		}
	}
	return l.Back().Value.(int)
}
func main() {
	s := "1+((2+3)*4)-5"
	words := []string{}
	re,err := regexp.Compile(`([0-9]+|\+|\-|\*|\/|\(|\))`)
	if err != nil {
		return
	}
	words = re.FindAllString(s,-1)

	for _,word := range words {
		fmt.Printf("%s\n", word)
	}
	words = suffixExp(words)
	fmt.Printf("%d\n",cal(words))
}