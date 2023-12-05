package set

type Set[E comparable] map[E]struct{}

func NewSet[E comparable]() *Set[E] {
	return &Set[E]{}
}

func (s *Set[E]) Insert(value E) {
	_, ok := (*s)[value]

	if !ok {
		(*s)[value] = struct{}{}
	}
}

func (s *Set[E]) Intersection(otherSet *Set[E]) *Set[E] {
	resultSet := make(Set[E])
	for k := range *s {
		_, ok := (*otherSet)[k]

		// Put value in result set
		if ok {
			resultSet[k] = struct{}{}
		}
	}
	return &resultSet
}

func (s *Set[E]) Size() int {
	return len(*s)
}
