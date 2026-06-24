package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "базовый случай",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "ответ не в начале",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "одинаковые числа",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "отрицательные числа",
			nums:   []int{-3, 4, 3, 90},
			target: 0,
			want:   []int{0, 2},
		},
		{
			name:   "нет решения",
			nums:   []int{1, 2, 3},
			target: 100,
			want:   nil,
		},
	}
 
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := TwoSum(tc.nums, tc.target)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("TwoSum(%v, %d) = %v, хотели %v",
					tc.nums, tc.target, got, tc.want)
			}
		})
	}
}

func BenchmarkTwoSumHashMap(b *testing.B) {
	nums := make([]int, 10_000)
	for i := range nums {
		nums[i] = i
	}
	target := 19_997
 
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum(nums, target)
	}
}