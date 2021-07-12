package linkedlist

import (
	"testing"
)

func TestAddIntoEmpty(t *testing.T) {
	ll := &LinkedList{}

	ll.Add(10)

	if ll.Head == nil {
		t.Fatalf("Head is empty")
	}
	if ll.Head.Value != 10 {
		t.Fatalf("Expected %v .Head.Value, got %v", 10, ll.Head.Value)
	}
	if ll.Head.Next != nil {
		t.Fatalf("Expected nil for .Head.Next")
	}
}

func TestAddIntoNotEmpty(t *testing.T) {
	ll := &LinkedList{}
	ll.Add(10)

	ll.Add(11)

	if ll.Head == nil {
		t.Fatalf("Head became nil")
	}
	if ll.Head.Value != 10 {
		t.Fatalf("Changed head value")
	}
	if ll.Head.Next == nil {
		t.Fatalf("New item was not created")
	}
	if ll.Head.Next.Value != 11 {
		t.Fatalf("Expected %v value for new node, got %v", 11, ll.Head.Next.Value)
	}
	if ll.Head.Next.Next != nil {
		t.Fatalf("Expected nil after new node")
	}
}

func TestReverseWithZeroItems(t *testing.T) {
	ll := &LinkedList{}

	ll.Reverse()

	if ll.Head != nil {
		t.Fatalf("Linked list's head was changed")
	}
}

func TestReverseWithOneItem(t *testing.T) {
	ll := &LinkedList{}
	ll.Add(1)

	ll.Reverse()

	if ll.Head == nil || ll.Head.Value != 1 || ll.Head.Next != nil {
		t.Fatalf("Linked list's head was changed")
	}
}

func TestReverseWithTwoItems(t *testing.T) {
	ll := &LinkedList{}
	ll.Add(1)
	ll.Add(2)

	ll.Reverse()

	if ll.Head == nil {
		t.Fatalf("Linked list's head became nil")
	}
	if ll.Head.Value != 2 {
		t.Fatalf("Expected first item's value %v, got %v", 2, ll.Head.Value)
	}
	secondItem := ll.Head.Next
	if secondItem == nil {
		t.Fatalf("Only one item in list")
	}
	if secondItem.Value != 1 {
		t.Fatalf("Expected second item's value %v, got %v", 1, secondItem.Value)
	}
	if secondItem.Next != nil {
		t.Fatalf("Expected nil after second item")
	}
}

func TestReverseWithManyItems(t *testing.T) {
	ll := &LinkedList{}
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)

	ll.Reverse()

	if ll.Head == nil {
		t.Fatalf("Linked list's head became nil")
	}
	if ll.Head.Value != 3 {
		t.Fatalf("Expected first item's value %v, got %v", 2, ll.Head.Value)
	}
	secondItem := ll.Head.Next
	if secondItem == nil {
		t.Fatalf("Only one item in list")
	}
	if secondItem.Value != 2 {
		t.Fatalf("Expected second item's value %v, got %v", 2, secondItem.Value)
	}
	thirdItem := secondItem.Next
	if thirdItem == nil {
		t.Fatalf("Only two items in list")
	}
	if thirdItem.Value != 1 {
		t.Fatalf("Expected third item's value %v, got %v", 1, thirdItem.Value)
	}
	if thirdItem.Next != nil {
		t.Fatalf("Expected nil after third item")
	}
}

func TestReversed(t *testing.T) {
	ll := &LinkedList{}
	ll.Add(1)
	ll.Add(2)

	res := Reversed(ll)

	if res.Head == nil {
		t.Fatalf("Empty list created")
	}
	if res.Head.Value != 2 {
		t.Fatalf("Expected first item's value %v, got %v", 2, res.Head.Value)
	}
	secondItem := res.Head.Next
	if secondItem == nil {
		t.Fatalf("Only one item in list")
	}
	if secondItem.Value != 1 {
		t.Fatalf("Expected second item's value %v, got %v", 1, secondItem.Value)
	}
	if secondItem.Next != nil {
		t.Fatalf("Expected nil after second item")
	}
}
