package binarysearch

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{name: "базовый случай", nums: []int{1, 3, 5, 7, 9, 11, 13, 15}, target: 9, want: 4},
		{name: "первый элемент", nums: []int{1, 3, 5, 7, 9, 11, 13, 15}, target: 1, want: 0},
		{name: "последний элемент", nums: []int{1, 3, 5, 7, 9, 11, 13, 15}, target: 15, want: 7},
		{name: "нет в массиве", nums: []int{1, 3, 5, 7, 9, 11, 13, 15}, target: 6, want: -1},
		{name: "меньше минимума", nums: []int{1, 3, 5, 7}, target: 0, want: -1},
		{name: "больше максимума", nums: []int{1, 3, 5, 7}, target: 100, want: -1},
		{name: "один элемент найден", nums: []int{5}, target: 5, want: 0},
		{name: "один элемент не найден", nums: []int{5}, target: 3, want: -1},
		{name: "пустой массив", nums: []int{}, target: 1, want: -1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := search(tc.nums, tc.target)
			if got != tc.want {
				t.Errorf("search(%v, %d) = %d, хотели %d", tc.nums, tc.target, got, tc.want)
			}
		})
	}
}

func BenchmarkSearch(b *testing.B) {
	nums := make([]int, 100_000)
	for i := range nums {
		nums[i] = i * 2
	}
	target := 99_998
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		search(nums, target)
	}
}