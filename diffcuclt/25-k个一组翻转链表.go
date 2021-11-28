package diffcuclt

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/solution/tu-jie-kge-yi-zu-fan-zhuan-lian-biao-by-user7208t/
func reverseKGroup(head *ListNode, k int) *ListNode {

	dummy := &ListNode{
		Next: head,
	}
	pre, end := dummy, dummy

	for end.Next != nil {

		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		// 反转
		start := pre.Next
		next := end.Next
		end.Next = nil // 先断开
		pre.Next = reverse(start)
		start.Next = next // 接上去

		// 重新赋值
		pre = start
		end = pre
	}

	return dummy.Next

}

func reverse(head *ListNode) (*ListNode) {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
