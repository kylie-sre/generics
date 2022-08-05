package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

const (
	numberInt   int   = 0
	numberInt8  int8  = 0
	numberInt16 int16 = 0
	numberInt32 int32 = 0
	numberInt64 int64 = 0
)

type Foo[T fmt.Stringer] struct {
	Foo *T
}

type Number interface {
	int | int8 | int16 | int32 | int64
}

type Set[T Number] struct {
	values map[T]bool
}

func NewSet[T Number](_ T) *Set[T] {
	return &Set[T]{
		values: map[T]bool{},
	}
}

func (s *Set[T]) Add(value T) {
	s.values[value] = true
}

func (s *Set[T]) Get(value T) (ok bool) {
	_, ok = s.values[value]
	return
}

func (s *Set[T]) Values() (results []T) {
	results = make([]T, 0)
	for k := range s.values {
		results = append(results, k)
	}
	return
}

func (s Set[T]) String() string {
	var buf bytes.Buffer
	buf.WriteString("Set{")
	values := s.Values()
	for idx, val := range values {
		buf.WriteString(strconv.FormatInt(int64(val), 10))
		if idx+1 != len(values) {
			buf.WriteString(", ")
		}
	}
	buf.WriteString("}")
	return buf.String()
}

func main() {
	s := NewSet(numberInt)
	s.Add(4)
	s.Add(5)
	s.Add(4)
	fmt.Println(s)
	fmt.Printf("%T\n", s.Values())

	f := Foo[Set[int]]{Foo: s}
	fmt.Println(f)
	fmt.Printf("%T\n", f)
}
