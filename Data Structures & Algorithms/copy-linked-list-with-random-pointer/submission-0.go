/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    nodeMap := make(map[*Node]*Node)
    nodeMap[nil] = nil
    curr := head
    for curr != nil {
        newNode := &Node{Val: curr.Val}
        nodeMap[curr] = newNode
        curr = curr.Next
    }

    curr = head
    for curr != nil {
        node, _ := nodeMap[curr]
        nextNode, _ := nodeMap[curr.Next]
        randomNode, _ := nodeMap[curr.Random]
        node.Next = nextNode
        node.Random = randomNode
        curr = curr.Next
    }
    return nodeMap[head]
}
