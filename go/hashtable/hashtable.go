package hashtable

import (
	"errors"
	"fmt"
)

type Entry struct {
	key   string
	value any
	next  *Entry
}

type Hashtable struct {
	buckets []*Entry
	count    int
	size     int
}

func NewHashTable() *Hashtable {
	return &Hashtable{
		buckets: make([]*Entry, 6),
		count:    0,
		size: 6,
	}
}

func (h *Hashtable) hash(key string) int {
	hash := 0
	prime := 31
	for _, rn := range key {
		hash = prime*hash + int(rn)
	}
	return hash % h.size
}

func (h *Hashtable) rehash(key string, newSize int) int {
	hash := 0
	prime := 31
	for _, rn := range key {
		hash = prime*hash + int(rn)
	}
	return hash % newSize
}

func (h *Hashtable) Insert(key string, value any) {

	idx := h.hash(key)

	// search if it already exists
	for t := h.buckets[idx]; t != nil; t = t.next {
		
		// update value if exists
		if t.key == key {
			t.value = value
			return
		}
	}

	// inserting new entry
	entry := &Entry{key: key,value: value, next: h.buckets[idx]}
	h.buckets[idx] = entry
	h.count++

	h.resize()
}

func (h *Hashtable) Search(key string) (value any, err error) {
	idx := h.hash(key)
	if h.buckets[idx] != nil {
		for t := h.buckets[idx]; t != nil; t = t.next {
			if t.key == key {
				return t.value, nil
			}
		}
	}
	return nil, errors.New("key does not exist in hashtable")
}

func (h *Hashtable) Remove(key string) {

	idx := h.hash(key)
	if h.buckets[idx] != nil {
		current := h.buckets[idx]

		// if head of list
		if current.key == key {
			h.buckets[idx] = h.buckets[idx].next
			h.count--
			return
		}

		for current != nil {
			next := current.next
			if next != nil && next.key == key {
				current.next = next.next
				h.count--
				return
			}
			current = next
		}
	}
}

func (h *Hashtable) ContainsKey(key string) bool {
	idx := h.hash(key)
	for t := h.buckets[idx]; t != nil; t = t.next {
		if t.key == key {
			return true
		}
	}

	return false
}

func (h *Hashtable) Size() int {
	return h.count
}

func (h *Hashtable) resize() {
	load_factor := float32(h.count) / float32(h.size)

	if load_factor >= 0.7 {
		newSize := h.size * 2
		newBuckets := make([]*Entry,newSize)

		for _, bucket := range h.buckets {
			for t := bucket; t != nil; t = t.next {
				newIdx := h.rehash(t.key, newSize) 
				newEntry := &Entry{key: t.key, value: t.value, next: newBuckets[newIdx]}
				newBuckets[newIdx] = newEntry
			}
		}
		h.buckets = newBuckets
		h.size = newSize
	}
}

func (h *Hashtable) Print() {
	for idx, bucket := range h.buckets {
		if bucket != nil {
			fmt.Printf("[%d]:",idx)
			for t := bucket; t != nil; t = t.next {
				fmt.Printf(" %v=%v ",t.key,t.value)
				if t.next != nil {
					fmt.Print("-->")
				}
			}
			fmt.Println()
		}
		
		
	}
}