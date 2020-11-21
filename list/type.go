package list

type LinkedList struct {
	Next  *LinkedList
	Value int
}

func NewlinkedList() *LinkedList {
	head := &LinkedList{
		Value: 0,
	}

	temp := head
	for i := 1; i < 10; i++ {
		newstruct := &LinkedList{
			Value: i,
		}
		temp.Next = newstruct
		temp = newstruct
	}

	return head
}
