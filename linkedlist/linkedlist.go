package linkedlist

import (
	"fmt"
	"github.com/la1t/gostructures/stack"
)

type Node struct {
	Next  *Node
	Value int
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) Add(value int) {
	if ll.Head == nil {
		ll.Head = &Node{Value: value}
		return
	}
	lastItem := ll.Head
	for lastItem.Next != nil {
		lastItem = lastItem.Next
	}
	lastItem.Next = &Node{Value: value}
}

func (ll *LinkedList) Print() {
	for item := ll.Head; item != nil; item = item.Next {
		if item.Next != nil {
			fmt.Printf("%v -> ", item.Value)
		} else {
			fmt.Printf("%v\n", item.Value)
		}
	}
}

func (ll *LinkedList) Reverse() {
	if ll.Head == nil || ll.Head.Next == nil {
		return
	}
	n1 := ll.Head
	n2 := ll.Head.Next
	n1.Next = nil
	for {
		next := n2.Next
		n2.Next = n1
		if next == nil {
			break
		}
		n1 = n2
		n2 = next
	}
	ll.Head = n2
}

func Reversed(ll *LinkedList) *LinkedList {
	st := stack.New()
	for e := ll.Head; e != nil; e = e.Next {
		st.Push(e.Value)
	}
	newLl := &LinkedList{}
	for v, err := st.Pop(); err == nil; v, err = st.Pop() {
		newLl.Add(v)
	}
	return newLl
}
