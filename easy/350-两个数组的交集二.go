package easy

func intersect(nums1 []int, nums2 []int) []int {
	var result []int
	nums1Map := make(map[int]int)
	for _, v := range nums1 {
		nums1Map[v]++
	}

	nums2Map := make(map[int]int)
	for _, v := range nums2 {
		nums2Map[v]++
	}

	for k, c1 := range nums1Map {
		if c2, ok := nums2Map[k]; ok {
			for i := 0; i < getSmallNum(c1, c2); i++ {
				result = append(result, k)
			}
		}
	}
	return result

}

func getSmallNum(a, b int) int {
	if a < b {
		return a
	}
	return b

}
