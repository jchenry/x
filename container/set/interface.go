package set

var x = struct{}{}

type Set map[any]struct{}

func (s *Set) Init() {
	for k := range *s {
		delete(*s, k)
	}
}

func (s *Set) Add(e any) {
	(*s)[e] = x
}

func (s *Set) Remove(e any) {
	delete(*s, e)
}

func (s *Set) Contains(e any) bool {
	_, c := (*s)[e]
	return c
}

func New() *Set {
	return new(Set)
}
