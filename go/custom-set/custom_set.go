package stringset

import (
	"fmt"
	"strings"
)

const testVersion = 3

type Set map[string]struct{}

func New() Set {
	return make(Set)
}

func NewFromSlice(keys []string) Set {
	ans := New()
	for _, k := range keys {
		ans.Add(k)
	}
	return ans
}

func (s Set) Add(e string) {
	if _, found := s[e]; !found {
		s[e] = struct{}{}
	}
}

func (s Set) Delete(e string) {
	delete(s, e)
}

func (s Set) Has(e string) bool {
	_, found := s[e]
	return found
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Len() int {
	return len(s)
}

func (s Set) Slice() []string {
	keys := make([]string, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}

func (s Set) String() string {
	keys := make([]string, 0, len(s))
	for k := range s {
		keys = append(keys, fmt.Sprintf("%q", k))
	}
	return fmt.Sprintf("{%s}", strings.Join(keys, ", "))
}

func Equal(s1, s2 Set) bool {
	if s1.Len() != s2.Len() {
		return false
	}
	for k := range s1 {
		if !s2.Has(k) {
			return false
		}
	}
	return true
}

func Subset(s1, s2 Set) bool {
	if s1.Len() > s2.Len() {
		return false
	}
	for k := range s1 {
		if !s2.Has(k) {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	if s1.Len() > s2.Len() {
		s1, s2 = s2, s1
	}
	for k := range s1 {
		if s2.Has(k) {
			return false
		}
	}
	return true
}

func Intersection(s1, s2 Set) Set {
	ans := New()
	if s1.Len() > s2.Len() {
		s1, s2 = s2, s1
	}
	for k := range s1 {
		if s2.Has(k) {
			ans.Add(k)
		}
	}
	return ans
}

func Union(s1, s2 Set) Set {
	ans := New()
	for k := range s1 {
		ans.Add(k)
	}
	for k := range s2 {
		ans.Add(k)
	}
	return ans
}

func Difference(s1, s2 Set) Set {
	ans := New()
	for k := range s1 {
		if !s2.Has(k) {
			ans.Add(k)
		}
	}
	return ans
}

func SymmetricDifference(s1, s2 Set) Set {
	return Union(Difference(s1, s2), Difference(s2, s1))
}
