package stack 

import (
	"testing"
)

func TestNew(t *testing.T) {
	st := New()

	if st.dl == nil {
		t.Fatalf("Stack must be created with double-linked list")
	}
}

func TestPush(t *testing.T) {
	st := New()

	st.Push(3)

	if st.dl.Back() == nil {
		t.Fatalf("Item was not pushed")
	}
}

func TestPopFromNotEmpty(t *testing.T) {
	st := New()
	st.Push(3)
	
	v, err := st.Pop()

	if err != nil {
		t.Fatalf("Error with popping item")
	}
	if v != 3 {
		t.Fatalf("Expected %v, got %v", 3, v)
	}
}

func TestPopFromEmpty(t *testing.T) {
	st := New()

	v, err := st.Pop()

	if err == nil {
		t.Fatalf("Expected error")
	}
	if v != 0 {
		t.Fatalf("Expected %v, got %v", 0, v)
	}
}