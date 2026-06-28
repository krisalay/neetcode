/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    if head == nil {
        return
    }
    // find the beginning of the second half of list
    beginingOfSecondHalf := getSecondHalfOfList(head)
    // reverse the second half of the list
    right := reverse(beginingOfSecondHalf)
    left := head
    
    for right != nil {
        leftNext := left.Next
        rightNext := right.Next
        left.Next = right
        right.Next = leftNext
        left = leftNext
        right = rightNext
    }
    left.Next = nil
    
}

func getSecondHalfOfList(head *ListNode) *ListNode {
    slow, fast := head, head.Next
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow.Next
}

func reverse(head *ListNode) *ListNode {
    curr := head
    var prev *ListNode
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev
}
