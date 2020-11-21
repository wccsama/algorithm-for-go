package list

func part_reverse(n int) {
	head := NewlinkedList()
	// 为空
	if head == nil {
		return
	}
	hc := head

	var count int
	for hc != nil {
		count++
		hc = hc.Next
	}

	hc = head
	for hc != nil {
		gap := count % n

		for gap > 0 {
			hc = hc.Next
		}

	}
}
