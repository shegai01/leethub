/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    new := &ListNode{
        Next: head,
    }

    cur := new

    for cur.Next != nil{
        if cur.Next.Val == val{
            cur.Next = cur.Next.Next
        } else {
            cur = cur.Next
        }
    }
    return new.Next

}