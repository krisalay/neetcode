/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    curr1, curr2 := list1, list2
    head := &ListNode{}
    curr3 := head

    for curr1 != nil && curr2 != nil {
        if curr1.Val <= curr2.Val {
            curr3.Next = curr1
            curr1 = curr1.Next
        } else {
            curr3.Next = curr2
            curr2 = curr2.Next
        }
        curr3 = curr3.Next
    }

    if curr1 != nil {
        curr3.Next = curr1
    }
    if curr2 != nil {
        curr3.Next = curr2
    }
    return head.Next
}
