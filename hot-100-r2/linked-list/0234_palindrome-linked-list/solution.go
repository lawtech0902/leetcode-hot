/*
__author__ = 'robin-luo'
__date__ = '2024/01/22 16:35'
*/

package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		if nums[i] != nums[j] {
			return false
		}
	}

	return true
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}

	return prev
}

func endOfFirstHalf(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func isPalindrome1(head *ListNode) bool {
	if head == nil {
		return true
	}

	firstHalfEnd := endOfFirstHalf(head)
	secondHalfFirst := reverseList(firstHalfEnd.Next)
	p1, p2 := head, secondHalfFirst
	res := true
	for res && p2 != nil {
		if p1.Val != p2.Val {
			res = false
		}

		p1 = p1.Next
		p2 = p2.Next
	}

	firstHalfEnd.Next = reverseList(secondHalfFirst)
	return res
}
