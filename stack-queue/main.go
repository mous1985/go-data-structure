package main

type Stack struct {
	items []int
}

func (s *Stack) push(i int) {
	s.items = append(s.items, i)
}
func (s *Stack) pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}
