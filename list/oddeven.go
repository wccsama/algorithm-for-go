package list

// 0 1 2 3 4 5 6
// 0 2 4 6 1 3 5
func odd() {
	head := NewlinkedList()
	even := head
	odd := head.Next
	temp := odd

	for odd.Next != nil && even.Next != nil {
		if odd.Next != nil {
			even.Next = odd.Next
			even = odd.Next
		}

		if even.Next != nil {
			odd.Next = even.Next
			odd = odd.Next
		}
	}
	even.Next = temp
	return
}
