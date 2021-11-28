package easy

import "sort"

// 1 哈希
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	map1 := make(map[int32]int)
	map2 := make(map[int32]int)
	for _, v := range s {
		map1[v]++
	}

	for _, v := range t {
		map2[v]++
	}
	for k, v := range map1 {
		if v != map2[k] {
			return false
		}
	}
	return true
}

// 2排序后对比
func isAnagramV2(s string, t string) bool {
	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	return string(s1) == string(s2)
}
