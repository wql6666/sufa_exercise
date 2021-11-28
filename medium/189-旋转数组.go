package medium

// 自己写的
func rotate(nums []int, k int) {
	k = k % len(nums)
	moveToHead := nums[len(nums)-k:]
	stay := nums[:len(nums)-k]
	result := make([]int, len(nums))
	result = append(moveToHead, stay...)
	copy(nums, result)
}

func rotate2(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}

// 原地旋转
func rotate1(nums []int, k int) {
	k = k % len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])

}

func reverse(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]
	}
}
