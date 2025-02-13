package hashtable

import (
	"testing"
)

func TestNewHashTable(t *testing.T) {
	h := NewHashTable()

	if expect, got := 0, h.count; expect != got {
		t.Errorf("Expected %v, got %v\n",expect,got)
	}
	if expect, got := 6, h.size; expect != got {
		t.Errorf("Expected %v, got %v\n",expect,got)
	}
	
	// empty slice
	for _, v := range h.buckets {
		if v != nil {
			t.Error("expected nil value for init bucket")
		}
	}
}

func TestHash(t *testing.T) {
	h := NewHashTable()
	hashKeys := []string {"santi", "james","anthony","smith"}
	hashCodes := make([]int,len(hashKeys))
	expectedCodes := []int{3,0,1,3}
	for i, name := range hashKeys {
		hashCodes[i] = h.hash(name)
	}
	for i, code := range hashCodes {
		if expectedCodes[i] != code {
			t.Errorf("Expected %v, got %v\n",expectedCodes[i], code)
		}
	}
	
}

func TestInsert(t *testing.T) {
	// []int{3,0,1,3}
	h := NewHashTable()
	h.Insert("santi",1997)
	h.Insert("james",1998)
	h.Insert("anthony",1997)
	h.Insert("smith",1941)

	// 4 entries
	if expect, got := 4, h.Size(); expect != got {
		t.Errorf("Expected %v, got %v\n",expect,got)
	}

	// smith --> santi
	if h.buckets[3].next.key != "santi" {
		t.Errorf("Expected %v, got %v\n","santi", h.buckets[3].next.key)
	}

	// update entry santi
	h.Insert("santi",1998)
	if h.buckets[3].next.value != 1998 {
		t.Errorf("Expected %v, got %v\n",1998, h.buckets[3].next.value)
	}
}

func TestSearch(t *testing.T) {
	h := NewHashTable()
	h.Insert("santi",1997)
	h.Insert("james",1998)
	h.Insert("anthony",1997)
	h.Insert("smith",1941)

	// exists
	v, err := h.Search("santi")
	if err != nil {
		t.Error(err)
	}

	if expect, got := 1997, v; expect != got {
		t.Errorf("Expected %v, got %v\n",expect,got)
	}

	// does not exist
	_, err = h.Search("tomatoes")
	if err == nil {
		t.Error("expected error to be thrown for search on nonexistant value")
	}
}

func TestRemove(t *testing.T) {
	// []int{3,0,1,3}
	h := NewHashTable()
	h.Insert("santi",1997)
	h.Insert("james",1998)
	h.Insert("anthony",1997)
	h.Insert("smith",1941)

	// remove at head
	h.Remove("james")
	if _, err := h.Search("james"); err == nil {
		t.Error("expected key james to be nonexistant after removal")
	}

	if expect, got := 3, h.Size(); expect != got {
		t.Errorf("Expected %v, got %v\n",expect,got)
	}

	// remove at tail
	h.Remove("santi")
	if _, err := h.Search("santi"); err == nil {
		t.Error("expected key santi to be nonexistant after removal")
	}

	h.Insert("santi", 1997)
	h.Insert("!", 0)
	// remove middle
	// ! --> santi --> smith
	h.Remove("santi")
	if h.buckets[3].next.key != "smith" {
		t.Errorf("Expected %v, got %v\n","smith", h.buckets[3].next.key)
	}
}
func TestPrint(t *testing.T) {
	h := NewHashTable()
	h.Insert("santi",1997)
	h.Insert("james",1998)
	h.Insert("anthony",1997)
	h.Insert("smith",1941)
	h.Insert("!", 0)

	h.Print()
}

func TestContainsKey(t *testing.T) {
	h := NewHashTable()
	h.Insert("santi",1997)
	h.Insert("james",1998)
	h.Insert("anthony",1997)
	h.Insert("smith",1941)
	h.Insert("!", 0)

	// h.Print()

	if expect, got := true,h.ContainsKey("santi"); expect != got {
		t.Errorf("Expected %v, got %v\n",expect,got)
	}

	h.Remove("!")
	if expect, got := false,h.ContainsKey("!"); expect != got {
		t.Errorf("Expected %v, got %v\n",expect,got)
	}

	h.Print()
}

func TestSize(t *testing.T) {
	h := NewHashTable()
	h.Insert("santi",1997)
	h.Insert("james",1998)
	h.Insert("anthony",1997)
	h.Insert("smith",1941)
	h.Insert("!", 0)

	if expect, got := 5,h.Size(); expect != got {
		t.Errorf("Expected %v, got %v\n",expect,got)
	}

	h.Remove("anthony")
	h.Remove("smith")
	h.Remove("!")	

	if expect, got := 2,h.Size(); expect != got {
		t.Errorf("Expected %v, got %v\n",expect,got)
	}
}