package easy


// 方法一：
func fabonacci(n int64) int64 {
	if n <= 2 {
		return n
	}
	return fabonacci(n-1) + fabonacci(n-2)

}

// 方法二
func fabonacci1(n int64) int64 {
	if n <= 2 {
		return n
	}
	a, b, c := int64(0), int64(1), int64(1)
	for i := int64(2); i <= n; i++ {
		a = b
		b = c
		c = a + b
	}
	return c
}

