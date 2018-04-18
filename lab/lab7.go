package main

import (
	"fmt"
	"container/list"
)
type Stack struct {
	list *list.List
}

func New() *Stack {
	list := list.New()
	return &Stack{list}
}
func (stack Stack)Pop() interface{}  {
	e := stack.list.Back()
	if e!= nil {
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}
func (stack Stack)Push(value interface{})  {
	stack.list.PushBack(value)
}

type Vehicle struct {
	wheels int
	lights int
}

func (v Vehicle)Setwheels(w int)()  {
	v.wheels = w

}
func (v Vehicle)Getwheels()(int)  {
	return v.wheels
}
func (v Vehicle)Setlights(l int) {
	v.lights = l
}
type Plane struct {
	Vehicle
	fly bool
	run bool
}

func (p *Plane)Canfly()bool {
	return p.fly
}
func (p *Plane)Canrun() bool {
	return p.run
}
func (p *Plane)Pro() () {
	fmt.Printf("%d %d %t\n",p.Vehicle.lights,p.Vehicle.wheels,p.fly)
}

type Car struct {
	Vehicle
	run bool
	fly bool
}

func (c *Car)Canfly() (bool) {
	return c.fly
}
func (c *Car)Canrun() (bool) {
	return c.run
}
func (c *Car)Pro()()  {
	fmt.Printf("%d %d %t\n",c.wheels,c.lights,c.run)
}
func main()  {

	c := Car{Vehicle{4,4,},true,false}
	c.Pro()
	p :=Plane{Vehicle{2,2},true,false}
	p.Pro()
	s := New()
	s.Push(1)
	s.Push(2)
	if e := s.Pop();e != nil {
		fmt.Printf("e :%d\n",e.(int))
	}

}