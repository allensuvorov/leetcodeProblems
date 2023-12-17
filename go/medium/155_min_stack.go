type MinStack struct {
	head *node
}

type node struct {
	val  int
	min  int
	next *node
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(val int) {
	if s.head == nil {
		s.head = &node{val: val, min: val, next: nil}
		return
	}

	s.head = &node{val: val, min: min(s.head.min, val), next: s.head}
}

func (s *MinStack) Pop() {
	s.head = s.head.next
}

func (s *MinStack) Top() int {
	return s.head.val
}

func (s *MinStack) GetMin() int {
	return s.head.min
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
