package main

import (
	"container/list"
	"fmt"
)

// Multiset is a sorted multiset data structure
type Multiset struct {
	elements map[interface{}]*list.Element
	list     *list.List
}

// NewMultiset creates a new multiset
func NewMultiset() *Multiset {
	return &Multiset{
		elements: make(map[interface{}]*list.Element),
		list:     list.New(),
	}
}

// Add adds an element to the multiset in sorted order
func (m *Multiset) Add(element interface{}) {
	if el, exists := m.elements[element]; exists {
		el.Value.(*item).count++
	} else {
		newItem := &item{element, 1}
		var inserted *list.Element
		for e := m.list.Front(); e != nil; e = e.Next() {
			if compare(e.Value.(*item).value, element) > 0 {
				inserted = m.list.InsertBefore(newItem, e)
				break
			}
		}
		if inserted == nil {
			inserted = m.list.PushBack(newItem)
		}
		m.elements[element] = inserted
	}
}

// Remove removes one occurrence of an element from the multiset
func (m *Multiset) Remove(element interface{}) {
	if el, exists := m.elements[element]; exists {
		el.Value.(*item).count--
		if el.Value.(*item).count <= 0 {
			m.list.Remove(el)
			delete(m.elements, element)
		}
	}
}

// Count returns the count of an element in the multiset
func (m *Multiset) Count(element interface{}) int {
	if el, exists := m.elements[element]; exists {
		return el.Value.(*item).count
	}
	return 0
}

// Contains checks if the element is in the multiset
func (m *Multiset) Contains(element interface{}) bool {
	_, exists := m.elements[element]
	return exists
}

// Size returns the total number of elements in the multiset (including duplicates)
func (m *Multiset) Size() int {
	total := 0
	for el := m.list.Front(); el != nil; el = el.Next() {
		total += el.Value.(*item).count
	}
	return total
}

// UniqueSize returns the number of unique elements in the multiset
func (m *Multiset) UniqueSize() int {
	return len(m.elements)
}

// Begin returns the smallest element in the multiset
func (m *Multiset) Begin() interface{} {
	if m.list.Len() == 0 {
		return nil
	}
	return m.list.Front().Value.(*item).value
}

// End returns the largest element in the multiset
func (m *Multiset) End() interface{} {
	if m.list.Len() == 0 {
		return nil
	}
	return m.list.Back().Value.(*item).value
}

type item struct {
	value interface{}
	count int
}

// compare is a helper function for comparing elements
func compare(a, b interface{}) int {
	switch a.(type) {
	case int:
		ai, bi := a.(int), b.(int)
		switch {
		case ai < bi:
			return -1
		case ai > bi:
			return 1
		default:
			return 0
		}
	case string:
		as, bs := a.(string), b.(string)
		switch {
		case as < bs:
			return -1
		case as > bs:
			return 1
		default:
			return 0
		}
	default:
		panic("unsupported type")
	}
}

func longestSubarray(nums []int, limit int) int {
	j := 0
	ms := NewMultiset()
	maxi := 0
	for ind, val := range nums {
		for ms.Begin() != nil && val-int(ms.Begin().(int)) > limit {
			ms.Remove(nums[j])
			j++
		}
		for ms.End() != nil && int(ms.End().(int))-val > limit {
			ms.Remove(nums[j])
			j++
		}
		ms.Add(val)
		if ind-j+1 > maxi {
			maxi = ind-j+1
		}
	}
	return maxi
}

func main() {
	nums := []int{8, 2, 4, 7}
	limit := 4
	fmt.Println("Longest subarray length:", longestSubarray(nums, limit)) // Output: 2

	nums = []int{10, 1, 2, 4, 7, 2}
	limit = 5
	fmt.Println("Longest subarray length:", longestSubarray(nums, limit)) // Output: 4

	nums = []int{4, 2, 2, 2, 4, 4, 2, 2}
	limit = 0
	fmt.Println("Longest subarray length:", longestSubarray(nums, limit)) // Output: 3
}
