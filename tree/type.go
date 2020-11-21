package tree

type Tree struct {
	Next  *Tree
	F     *Tree
	Left  *Tree
	Right *Tree
	value int
}

func NewTree() *Tree {
	root := &Tree{
		value: 0,
	}

	for i := 1; i < 3; i++ {
		left := &Tree{
			value: i,
		}

		right := &Tree{
			value: i,
		}

		root.Left = left
		root.Right = right
		root = left
	}

	return root
}

type Stack struct {
}

func (s *Stack) pop() interface{} {
	return nil
}

func (s *Stack) push(obj interface{}) {

}

func (s *Stack) top(obj interface{}) {

}

func (s *Stack) empty() bool {
	return true
}

type Queue struct {
}

func (s *Queue) enqueue(root *Tree) {

}

func (s *Queue) dequeue() *Tree {
	return nil
}

func (s *Queue) len() int {
	return 0
}

type helpTree struct {
	root    *Tree
	isFirst bool
}
