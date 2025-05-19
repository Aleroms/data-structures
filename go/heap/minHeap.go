package heap

import (
	"errors"
)

type MinHeap struct {
	data []int
}

func NewMinHeap() *MinHeap {

	return &MinHeap{
		data: []int{-1},
	}
}

func (h *MinHeap) GetMin() (int, error) {
	if len(h.data) > 1 {
		return h.data[1], nil
	}

	return -1, errors.New("cannot get minimum from empty heap")
}

func (h *MinHeap) Push(val int) {
	h.data = append(h.data, val)
	i := len(h.data) - 1
	

	// if child is smaller than parent
	for i > 1 && h.data[i] < h.data[i/2] {
		swp := h.data[i/2]
		h.data[i/2] = h.data[i]
		h.data[i] = swp

		i = i / 2
	}

}

func (h *MinHeap) Pop() (int, error) {
	if len(h.data) == 1 {
		return -1, errors.New("cannot pop on empty heap")
	}
	if len(h.data) == 2 {
		return h.data[1], nil
	}

	res := h.data[1]

	// replace root with rightmost leaf and pop
	i := len(h.data) - 1
	h.data[1] = h.data[i]
	h.data = h.data[:i]
	i = 1

	// structure property violation fix
	// check if left child exists first
	for 2*i < len(h.data) {
		
		// check if right child exists and is 
		// minimum between children and i > it
		if 2*i + 1 < len(h.data) && h.data[2*i+1] < h.data[2*i] && h.data[i] > h.data[2*i+1]{
			swp := h.data[i]
			h.data[2*i+1] = swp
			h.data[i] = swp
			i = 2*i + 1

			// swap left child
		} else if h.data[i] > h.data[2*i] {
			swp := h.data[2*i]
			h.data[2*i] = h.data[i]
			h.data[i] = swp
			i = 2 * i

			// break bcz values are the same; spot found
		} else {
			break
		}
	}

	return res, nil
}