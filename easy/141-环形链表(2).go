package easy

// 方法一：哈希表
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	nodeMap := make(map[*ListNode]interface{})
	for head.Next != nil {
		if _, ok := nodeMap[head]; ok {
			return true
		}
		nodeMap[head] = struct{}{}
		head = head.Next
	}
	return false

}

// 这个更快。空间更少。
// 方法二：快慢指针，如何退出循环，没有环则肯定会走到尾部，走到尾部的时候退出循环即可
// 每次追一步，至多n步能够追上。
func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil { // 若是会到尾部，那么是线fast会到达尾部为空的地方。
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next

	}
	return true

}
