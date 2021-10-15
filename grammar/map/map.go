// map.go

package mapt

// Set
type Set struct {
	Container map[interface{}]bool
}

// FuncValueMap ...
func FuncValueMap() (m map[int]func(op int) int) {
	m = make(map[int]func(op int) int, 4)
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }

	return
}

func (s *Set) Add(key interface{}) bool {
	s.Container[key] = true
	return true
}

func (s *Set) Delete(key interface{}) bool {
	c := s.Container
	if c[key] {
		delete(c, key)
	}

	return true
}

func (s *Set) Contains(key interface{}) bool {
	c := s.Container

	return c[key]
}

func (s *Set) Size() int {

	return len(s.Container)
}
