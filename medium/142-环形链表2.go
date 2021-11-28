package medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 方法一：hash
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	nodeMap := make(map[*ListNode]interface{})
	for head.Next != nil {
		if _, ok := nodeMap[head]; ok {
			return head
		}
		nodeMap[head] = struct{}{}
		head = head.Next
	}
	return nil
}

// 方法二:数学推到出来的
// https://leetcode-cn.com/problems/linked-list-cycle-ii/solution/huan-xing-lian-biao-ii-by-leetcode-solution/

func detectCycle2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil {
		if fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil

}

func detectCycle2_1(head *ListNode) *ListNode {
	slow := head
	fast := head

	for {
		if fast == nil || fast.Next == nil {
			return nil
		}

		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}

	result := head

	for ; result != slow; {
		result = result.Next
		slow = slow.Next
	}
	return result
}
