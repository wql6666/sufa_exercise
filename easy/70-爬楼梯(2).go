package easy

// 方法一 时间复杂度n的平方，空间n
func climbStairs(n int) int {
	cache := make(map[int]int)
	return climbStairsWithCache(n, cache)
}

func climbStairsWithCache(n int, cache map[int]int) int {
	if n <= 3 {
		return n
	}
	if v, ok := cache[n]; ok {
		return v
	}
	resp := climbStairsWithCache(n-1, cache) + climbStairsWithCache(n-2, cache)
	cache[n] = resp
	return resp
}

// 方法二 时间复杂度n，空间1
func climbStairs2(n int) int {
	if n < 3 {
		return n
	}
	p, q := 1, 2
	for i := 3; i <= n; i++ {
		p, q = q, p+q
	}
	return q

}
