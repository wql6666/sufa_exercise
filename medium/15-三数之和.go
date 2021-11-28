package medium

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// 方法一：会超时
func threeSum(nums []int) [][]int {
	tmpResult := make([][]int, 0)
	result := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					tmpResult = append(tmpResult, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}

	// 去重
	tmp := make(map[string]bool)
	for _, one := range tmpResult {
		sort.Ints(one)
		key := fmt.Sprintf("%v_%v_%v", one[0], one[1], one[2])
		tmp[key] = true
	}
	for key, _ := range tmp {
		strSlice := strings.Split(key, "_")
		i, _ := strconv.Atoi(strSlice[0])
		j, _ := strconv.Atoi(strSlice[1])
		k, _ := strconv.Atoi(strSlice[2])
		result = append(result, []int{i, j, k})
	}
	return result
}

// 排序后使用双指针
func threeSum2(nums []int) [][]int {
	result := make([][]int, 0)
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums)
	for i, v := range nums {
		if v > 0 {
			return result
		}
		// 得执行过，假如还相等就跳过，若改成nums[i] == nums[i+1]会可能漏掉
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for left, right := i+1, len(nums)-1; left < right; {
			sum := nums[left] + nums[right] + v
			if sum < 0 {
				left++
			}
			if sum > 0 {
				right--
			}
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				// 多个重复值的判断，还需要一个循环去重
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for right > left && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--

			}

		}

	}
	return result
}
