// package casing doesn't matter i think
package stack

// Capital first letter means Stack is accessible from other packages
type Stack[T int | string] struct {
	// lowercase first letter means value is private to this package
	value  []T
}

// this little * means we're adding a method to Stack
func (s *Stack[T]) Push(element T) {
	s.value = append(s.value, element)
}

func (s *Stack[T]) Pop() T {
	lastElementIndex := len(s.value) - 1
	element := s.value[lastElementIndex]
	s.value = s.value[:lastElementIndex]
	return element
}