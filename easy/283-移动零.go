package easy

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
// 示例:
//
// 输入: [0,1,0,3,12]
// 输出: [1,3,12,0,0]
//
// 说明:
//
//
// 必须在原数组上操作，不能拷贝额外的数组。
// 尽量减少操作次数。

// 记下0的位置，然后再对应位置进行替换。
func moveZeroes(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	slow := 0
	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			nums[slow] = nums[i]
			slow++
		}
	}
	if slow < length {
		for i := slow; i < len(nums); i++ {
			nums[i] = 0
		}
	}

}

func moveZeroes2(nums []int) {
	j := 0 // 记录0在的位置
	for i, v := range nums {
		if v != 0 {
			if i != j { // 去掉自己和自己换的情况
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}

}
