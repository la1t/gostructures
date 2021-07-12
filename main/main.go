package main

import (
	"github.com/la1t/gostructures/linkedlist"
)

func main() {
	list := &linkedlist.LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Print()
	list.Reverse()
	list.Print()
	linkedlist.Reversed(list).Print()
	list = &linkedlist.LinkedList{}
	list.Add(10)
	list.Add(15)
	linkedlist.Reversed(list).Print()
}
