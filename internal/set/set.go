package set

type Set map[interface{}]struct{}

func (s *Set) Insert(value interface{}) {
	_, ok := (*s)[value]

	if !ok {
		(*s)[value] = struct{}{}
	}
}

func (s *Set) Intersection(otherSet *Set) *Set {
	resultSet := make(Set)
	for k := range *s {
		_, ok := (*otherSet)[k]

		// Put value in result set
		if ok {
			resultSet[k] = struct{}{}
		}
	}
	return &resultSet
}

func (s *Set) Size() int {
	return len(*s)
}
