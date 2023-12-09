package stack

type Stack[E any] []E

func NewStack[E any]() *Stack[E] {
	return &Stack[E]{}
}

func (s *Stack[E]) Push(value E) {
	*s = append(*s, value)
}

func (s *Stack[E]) Pop() E {
	retval := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return retval
}

func (s *Stack[E]) Length() int {
	return len(*s)
}
