package ingredients

// I'm going to essentially mimic the set object I have in
// my utils directory, but for integers specifically. Also
// I never exported it so instead of mucking with it too much
// I'll just recreate it here.

// Copilot said this code is used in a bunch of references elsewhere!
// apparently a bunch of people have similar implementations for leetcode problems.
// Cool I guess? I dunno maybe that's where I got this from initially.

var exists = struct{}{}

type set struct {
	m map[int]struct{}
}

func NewSet() *set {
	s := &set{}
	s.m = make(map[int]struct{})
	return s
}

func (s *set) Add(value int) {
	s.m[value] = exists
}

func (s *set) Remove(value int) {
	delete(s.m, value)
}

func (s *set) Contains(value int) bool {
	_, c := s.m[value]
	return c
}

func (s *set) Size() int {
	return len(s.m)
}
