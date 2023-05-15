/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *ListNode, k int) *ListNode {
    n := 0
    curr := head
    for curr != nil {
        n++
        curr = curr.Next
    }
    if k > n {
        return head
    }
    i := 1
    curr = head
    for i < k {
        i++
        curr = curr.Next
    }
    j := n - k + 1
    curr2 := head
    for j > 1 {
        j--
        curr2 = curr2.Next
    }
    curr.Val, curr2.Val = curr2.Val, curr.Val
    return head
}