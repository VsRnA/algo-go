package twosum

func TwoSum(nums []int, target int) []int {
	numSet := make(map[int]int)

	for i, value := range nums {
		needValue := target - value;
		if j, ok := numSet[needValue]; ok {
			return []int{j, i}
		}
		numSet[value] = i;
	}
	return nil;
}
