package easy

// 别忘了边界条件
func removeDuplicates(nums []int) int {
	// 快慢指针，快指针负责去探测不重复的值，慢指针负责保存不重复的值。
	length := len(nums)
	if length <= 1 {
		return length
	}
	slow := 1
	for i := 1; i < length; i++ {
		if nums[i] != nums[i-1] {
			nums[slow] = nums[i]
			slow++
		}
	}
	return slow
}
