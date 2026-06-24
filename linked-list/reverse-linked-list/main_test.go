package reverselinkedlist

import (
	"testing"
)

func toList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{value: vals[0]}
	curr := head
	for _, v := range vals[1:] {
		curr.next = &ListNode{value: v}
		curr = curr.next
	}
	return head
}

func toSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.value)
		head = head.next
	}
	return result
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestReverseList(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{name: "базовый случай", nums: []int{1, 2, 3, 4, 5}, want: []int{5, 4, 3, 2, 1}},
		{name: "два элемента", nums: []int{1, 2}, want: []int{2, 1}},
		{name: "один элемент", nums: []int{1}, want: []int{1}},
		{name: "пустой список", nums: []int{}, want: nil},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := toSlice(reverseList(toList(tc.nums)))
			if !equal(got, tc.want) {
				t.Errorf("reverseList(%v) = %v, хотели %v", tc.nums, got, tc.want)
			}
		})
	}
}

func TestReverseListRecursive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{name: "базовый случай", nums: []int{1, 2, 3, 4, 5}, want: []int{5, 4, 3, 2, 1}},
		{name: "два элемента", nums: []int{1, 2}, want: []int{2, 1}},
		{name: "один элемент", nums: []int{1}, want: []int{1}},
		{name: "пустой список", nums: []int{}, want: nil},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := toSlice(reverseListRecursive(toList(tc.nums)))
			if !equal(got, tc.want) {
				t.Errorf("reverseListRecursive(%v) = %v, хотели %v", tc.nums, got, tc.want)
			}
		})
	}
}

func BenchmarkReverseListIterative(b *testing.B) {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverseList(toList(nums))
	}
}

func BenchmarkReverseListRecursive(b *testing.B) {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverseListRecursive(toList(nums))
	}
}