/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func (l *ListNode) Len() int {
    curr := l
    len := 0
    for curr != nil {
        len++
        curr = curr.Next
    }
    return len
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    len := head.Len()
    if len - n == 0 {
        return head.Next
    }
    curr := head
    ptr := len - n
    for ptr > 1 {
        curr = curr.Next
        ptr--
    }
    next := curr.Next
    curr.Next = next.Next
    next.Next = nil
    return head
}
