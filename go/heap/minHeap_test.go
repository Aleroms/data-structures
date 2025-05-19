package heap

import (
	"fmt"
	"testing"
)

func TestNewMinHeap(t *testing.T){
	h := NewMinHeap()
	fmt.Printf("data: %v", h.data)

	if len(h.data) != 1 {
		t.Errorf("expected data to be size 1 on initialization")
	}
}

func TestMinHeapPush(t *testing.T){
	h := NewMinHeap()
	values := []int{100,200,350,90,57,75,25,18,18,20,1}

	for _, v := range values {
		h.Push(v)
	}

	expect := []int{-1,1,18,57,25,18,350,75,200,90,100,20}

	if len(h.data) != len(expect) {
		t.Errorf("incorrect size for test values")
	}

	for i, result := range h.data {
		if result != expect[i] {
			t.Errorf("expected %d, got %d\n", result, expect[i])
		} 
	}

}

func TestGetMin(t *testing.T){
	h := NewMinHeap()

	if _, err := h.GetMin(); err == nil {
		t.Error("expected GetMin to throw error on getting minimum on empty heap.")
	}

	val := 3
	val2 := 1
	h.Push(val)
	if v, _ := h.GetMin(); v != val {
		t.Errorf("expected minimum value to be %d\n", val)
	}
	h.Push(val2)
	if v, _ := h.GetMin(); v != val2 {
		t.Errorf("expected %d but got %d\n", val2, v)
	}
}

func TestPopMin(t *testing.T) {
	h := NewMinHeap()
	test := []int{14,19,16,21,26,19,68,65,30}

	// pop on empty heap
	if _, err := h.Pop(); err == nil {
		t.Error("expected error to be thrown on pop on empty heap")
	}

	for _, v := range test {
		h.Push(v)
	}

	// pop on heap with values
	expect := test[0]
	if result, _ := h.Pop(); result != expect {
		t.Errorf("expected %d but got %d\n", expect, result)
	}

	// last value should be 65
	if h.data[len(h.data)-1] != 65 {
		t.Errorf("expected %d but got %d\n", 65, h.data[len(h.data)-1])
	}

}