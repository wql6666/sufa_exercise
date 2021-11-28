package easy

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i, v := range nums {
		if index, ok := numsMap[target-v]; ok {
			return []int{i, index}
		}
		numsMap[v] = i
	}
	return nil
}
