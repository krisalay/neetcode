/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    nodeArray := []*ListNode{}
    curr := head
    for curr != nil {
        nodeArray = append(nodeArray, curr)
        curr = curr.Next
    }
    
    left, right := 0, len(nodeArray)-1
    for left < right {
        next := left + 1
        nodeArray[left].Next = nodeArray[right]
        nodeArray[right].Next = nodeArray[next]
        left++
        right--
    }
    nodeArray[left].Next = nil
}
