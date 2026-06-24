package reverselinkedlist

type ListNode struct {
	value int
	next *ListNode 
}

func reverseList (head *ListNode) *ListNode {
  var prev *ListNode
	curr := head
	
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	return prev;
}

func reverseListRecursive (head *ListNode) *ListNode {
	if (head == nil || head.next == nil) {
		return head
	}
	nowHead := reverseListRecursive(head.next);
	head.next.next = head;
	head.next = nil;
	return nowHead;
}