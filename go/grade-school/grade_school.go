package school

import (
	"sort"
)

type Grade struct {
	grade   int
	student []string
}

type School struct {
	students map[int][]string
}

func New() *School {
	return &School{map[int][]string{}}
}

func (s *School) Add(stu string, grade int) {
	s.students[grade] = append(s.students[grade], stu)
}

func (s *School) Enrollment() []Grade {
	grades := make([]Grade, len(s.students))
	keys := make([]int, len(s.students))
	c := 0
	for k, _ := range s.students {
		keys[c] = k
		c++
	}
	sort.Ints(keys)
	c = 0
	for _, k := range keys {
		grades[c].grade = k
		sorted := make([]string, len(s.students[k]))
		copy(sorted, s.students[k])
		sort.Strings(sorted)
		grades[c].student = sorted
		c++
	}
	return grades
}

func (s *School) Grade(grade int) []string {
	return s.students[grade]
}
