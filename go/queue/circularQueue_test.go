package queue

import (
	"fmt"
	"testing"
)

func TestNewCircularQueue(t *testing.T) {
	sz := 6
	q := NewCircularQueue(sz)

	if q != nil {

		if expect, got := -1, q.front; expect != got {
			t.Errorf("expected %v, got %v\n",expect,got)
		}
		if expect, got := -1, q.rear; expect != got {
			t.Errorf("expected %v, got %v\n",expect,got)
		}
		if expect, got := sz, q.length; expect != got {
			t.Errorf("expected %v, got %v\n",expect,got)
		}
	} else {
		t.Error("Circular Queue is nil")
	}

	
}

func TestIsEmpty2(t *testing.T) {
	sz := 6
	q := NewCircularQueue(sz)

	if expect, got := true, q.IsEmpty(); expect != got {
		t.Errorf("expected %v, got %v\n",expect,got)
	}
	
	q.Enqueue(1)
	if expect, got := false, q.IsEmpty(); expect != got {
		t.Errorf("expected %v, got %v\n",expect,got)
	}

	_, err := q.Dequeue()
	if err != nil {
		t.Error(err)
	}
	if expect, got := true, q.IsEmpty(); expect != got {
		t.Errorf("expected %v, got %v\n",expect,got)
	}

	for i := 0; i < sz + 1; i++ {
		q.Enqueue(i + 1)
	}

	if expect, got := false, q.IsEmpty(); expect != got {
		t.Errorf("expected %v, got %v\n",expect,got)
	}
}

func TestSize2(t *testing.T) {
	sz := 6
	q := NewCircularQueue(sz)

	// newly created queue size should be 0=empty
	if expect, got := 0, q.Size(); expect != got {
		t.Errorf("expected %v, got %v\n",expect,got)
	}

	
	for i := 0; i < sz; i++ {
		q.Enqueue(i + 1)
	}

	// should provide length when queue is full
	if expect, got := sz, q.Size(); expect != got {
		t.Errorf("expected %v, got %v\n",expect,got)
	}

	q.Enqueue(7)
	q.Enqueue(8)

	if expect, got := sz, q.Size(); expect != got {
		t.Errorf("expected %v, got %v\n",expect,got)
	}
}

func TestIsFull(t *testing.T) {
	sz := 6
	q := NewCircularQueue(sz)

	if expect, got := false, q.IsFull(); expect != got {
		t.Errorf("expected %v, got %v\n",expect,got)
	}

	for i:=0; i < sz; i++ {
		q.Enqueue(i + 1)
	}

	if expect, got := true, q.IsFull(); expect != got {
		t.Errorf("expected %v, got %v\n",expect,got)
	}
}

func TestToArray(t *testing.T) {
	sz := 6
	q := NewCircularQueue(sz)
	xq := q.ToArray()

	// length should be size
	if expect, got := sz, len(xq); expect != got {
		t.Errorf("expected %v, got %v\n",expect,got)
	}

	for i := 0; i < sz; i++ {
		q.Enqueue(i + 1)
	}

	xq = q.ToArray()

	fmt.Println(xq...)
}

func TestPeek2(t *testing.T) {
	q := NewCircularQueue(6)

	if _, err := q.Peek(); err == nil {
		t.Error("expected error when peek on empty q")
	}

	q.Enqueue(1)
	if v, _ := q.Peek(); v != 1 {
		t.Errorf("expected 1, got %v\n",v)
	}
	
	for i := 0; i < 6; i++ {
		q.Enqueue(i + 2)
	}
	q.Print()
	fmt.Println(q.front,q.rear)
	
	if v, _ := q.Peek(); v != 3 {
		t.Errorf("expected 1, got %v\n",v)
	}
}

func TestEnqueue2(t *testing.T) {
	// Create a queue with capacity 3
	cq := NewCircularQueue(3)

	// Enqueue elements
	cq.Enqueue(1)
	cq.Enqueue(2)
	cq.Enqueue(3)

	// Check if the values are in the correct positions
	expected := []any{1, 2, 3}
	actual := cq.ToArray()

	for i, v := range expected {
		if actual[i] != v {
			t.Errorf("Expected %v at index %d, got %v", v, i, actual[i])
		}
	}

	// Test circular behavior (overwrite oldest element)
	cq.Enqueue(4) // Overwrites `1`

	// Expected: [4, 2, 3] since `4` overwrites `1`
	expectedAfterOverwrite := []any{4, 2, 3}
	actualAfterOverwrite := cq.ToArray()

	for i, v := range expectedAfterOverwrite {
		if actualAfterOverwrite[i] != v {
			t.Errorf("Expected %v at index %d, got %v after overwriting", v, i, actualAfterOverwrite[i])
		}
	}

	// Test Enqueue on a full queue (should overwrite oldest)
	cq.Enqueue(5) // Overwrites `2`
	expectedAfterSecondOverwrite := []any{4, 5, 3}
	actualAfterSecondOverwrite := cq.ToArray()

	for i, v := range expectedAfterSecondOverwrite {
		if actualAfterSecondOverwrite[i] != v {
			t.Errorf("Expected %v at index %d, got %v after second overwrite", v, i, actualAfterSecondOverwrite[i])
		}
	}
}

func TestDequeue2(t *testing.T) {
	// Create a queue with capacity 3
	cq := NewCircularQueue(3)

	// Enqueue elements
	cq.Enqueue(1)
	fmt.Printf("front:%d=%v\trear:%d=%v\n",cq.front,cq.data[cq.front],cq.rear,cq.data[cq.rear])
	cq.Enqueue(2)
	fmt.Printf("front:%d=%v\trear:%d=%v\n",cq.front,cq.data[cq.front],cq.rear,cq.data[cq.rear])
	cq.Enqueue(3)
	fmt.Printf("front:%d=%v\trear:%d=%v\n",cq.front,cq.data[cq.front],cq.rear,cq.data[cq.rear])

	cq.Print()


	// Dequeue first element (should be 1)
	dequeued, err := cq.Dequeue()
	if err != nil {
		t.Errorf("Unexpected error on Dequeue: %v", err)
	}
	if dequeued != 1 {
		t.Errorf("Expected dequeued value 1, got %v", dequeued)
	}

	// Dequeue second element (should be 2)
	dequeued, err = cq.Dequeue()
	if err != nil {
		t.Errorf("Unexpected error on Dequeue: %v", err)
	}
	if dequeued != 2 {
		t.Errorf("Expected dequeued value 2, got %v", dequeued)
	}

	// Dequeue third element (should be 3)
	dequeued, err = cq.Dequeue()
	if err != nil {
		t.Errorf("Unexpected error on Dequeue: %v", err)
	}
	if dequeued != 3 {
		t.Errorf("Expected dequeued value 3, got %v", dequeued)
	}

	// Try to dequeue from empty queue (should return error)
	_, err = cq.Dequeue()
	if err == nil {
		t.Errorf("Expected error when dequeuing from empty queue, but got nil")
	}
}

