/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    curr := dummy
    p1, p2 := l1, l2
    carry := 0
    for p1 != nil || p2 != nil || carry > 0 {
        val1, val2 := 0, 0
        if p1 != nil {
            val1 = p1.Val
        }
        if p2 != nil {
            val2 = p2.Val
        }
        sum := val1 + val2 + carry
        carry = sum / 10
        curr.Next = &ListNode{Val: sum%10}
        curr = curr.Next
        if p1 != nil {
            p1 = p1.Next
        }
        if p2 != nil {
            p2 = p2.Next
        }
    }
    return dummy.Next
}
