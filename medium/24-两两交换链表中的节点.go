package medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/swap-nodes-in-pairs/solution/shou-hua-tu-jie-24-liang-liang-jiao-huan-lian-biao/
func swapPairs(head *ListNode) *ListNode {
	// 辅助节点 画图
	dummy := &ListNode{}
	dummy.Next = head
	prev := dummy

	for head != nil && head.Next != nil {
		next := head.Next
		head.Next = next.Next
		next.Next = head
		prev.Next = next

		prev = head
		head = head.Next
	}
	return dummy.Next

}

// / 递归版本，不太明白
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}
