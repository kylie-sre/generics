package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type IntSet map[int]bool

func (s IntSet) Set(value int) {
	s[value] = true
}

func (s IntSet) Get(value int) (ok bool) {
	_, ok = s[value]
	return
}

func (s IntSet) Values() (results []int) {
	results = make([]int, 0)
	for k := range s {
		results = append(results, k)
	}
	return
}

func (s IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("IntSet{")
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

type Int8Set map[int8]bool

func (s Int8Set) Set(value int8) {
	s[value] = true
}

func (s Int8Set) Get(value int8) (ok bool) {
	_, ok = s[value]
	return
}

func (s Int8Set) Values() (results []int8) {
	results = make([]int8, 0)
	for k := range s {
		results = append(results, k)
	}
	return
}

func (s Int8Set) String() string {
	var buf bytes.Buffer
	buf.WriteString("IntSet{")
	values := s.Values()
	for idx, val := range values {
		buf.WriteString(strconv.Itoa(int(val)))
		if idx+1 != len(values) {
			buf.WriteString(", ")
		}
	}
	buf.WriteString("}")
	return buf.String()
}

func main() {
	s := make(IntSet)
	s.Set(4)
	s.Set(5)
	s.Set(4)
	fmt.Println(s)

	s8 := make(Int8Set)
	s8.Set(4)
	s8.Set(5)
	s8.Set(4)
	fmt.Println(s8)
}
