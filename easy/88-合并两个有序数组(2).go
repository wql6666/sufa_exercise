package easy

import "sort"

func mergeSelf(nums1 []int, m int, nums2 []int, n int) {
	for p1, p2, tail := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; tail-- {
		if p1 == -1 {
			nums1[tail] = nums2[p2]
			p2--
		} else if p2 == -1 {
			nums1[tail] = nums1[p1]
			p1--
		} else if nums1[p1] < nums2[p2] {
			nums1[tail] = nums2[p2]
			p2--
		} else {
			nums1[tail] = nums1[p1]
			p1--
		}
	}

}

// 1
func merge(nums1 []int, m int, nums2 []int, n int) {

	copy(nums1[m:], nums2)
	sort.Ints(nums1)

}

// 双指针法
func merge1(nums1 []int, m int, nums2 []int, n int) {

	sorted := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for {
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}
		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}
		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++

		}

	}
	copy(nums1, sorted)

}

// 避免额外空间的双指针，先选出最大的
func merge2(nums1 []int, m int, nums2 []int, n int) {
	for p1, p2, tail := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; tail-- {
		var cur int
		if p1 == -1 {
			cur = nums2[p2]
			p2--

		} else if p2 == -1 {
			cur = nums1[p1]
			p1--

		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--

		} else {
			cur = nums2[p2]
			p2--

		}
		nums1[tail] = cur
	}

}
