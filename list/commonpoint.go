package list

import "math"

func commonPoint() *LinkedList {
	head1 := NewlinkedList()
	head2 := NewlinkedList()

	i := 0
	for {
		if head1 == nil {
			break
		}
		head1 = head1.Next
		i++
	}

	j := 0
	for {
		if head2 == nil {
			break
		}
		head2 = head2.Next
		j++
	}

	step := math.Abs(float64(j - i))

	head := NewlinkedList()
	for {
		if head == nil {
			return nil
		}

		if step == 0 {
			break
		}

		head = head.Next
	}

	head3 := NewlinkedList()
	for {
		if head == nil {
			return nil
		}

		if head3 == head {
			return head3
		}

		head = head.Next
		head3 = head3.Next
	}

	return nil
}
