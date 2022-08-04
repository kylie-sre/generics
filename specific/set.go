package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type Set map[int]bool

func (s Set) Set(value int) {
	s[value] = true
}

func (s Set) Get(value int) (ok bool) {
	_, ok = s[value]
	return
}

func (s Set) Values() (results []int) {
	results = make([]int, 0)
	for k := range s {
		results = append(results, k)
	}
	return
}

func (s Set) String() string {
	var buf bytes.Buffer
	buf.WriteString("Set{")
	values := s.Values()
	for idx, val := range values {
		buf.WriteString(strconv.Itoa(val))
		if idx+1 != len(values) {
			buf.WriteString(", ")
		}
	}
	buf.WriteString("}")
	return buf.String()
}

func main() {
	s := make(Set)
	s.Set(4)
	s.Set(5)
	s.Set(4)
	fmt.Println(s)
}
