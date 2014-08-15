package set

import (
  "bytes"
  "fmt"
)

type Set struct {
  table map[interface{}]bool
}

func New() *Set {
  set := new(Set)
  set.table = make(map[interface{}]bool)
  return set
}

func (s *Set) Dup() *Set {
  dup := New()

  for value := range s.Iter() {
    dup.Add(value)
  }

  return dup
}

func (s *Set) Equals(other *Set) bool {
  if s.Size() != other.Size() {
    return false
  }

  for item := range s.Iter() {
    if !other.Includes(item) {
      return false
    }
  }

  return true
}

func (s *Set) Size() int {
  return len(s.table)
}

func (s *Set) Add(values ...interface{}) {
  for _, value := range values {
    s.table[value] = true
  }
}

func (s *Set) Remove(value interface{}) {
  delete(s.table, value)
}

func (s *Set) Includes(value interface{}) bool {
  return s.table[value] == true
}

func (s *Set) Clear() {
  s.table = make(map[interface{}]bool)
}

func (s *Set) Union(other *Set) *Set {
  result := s.Dup()
  for item := range other.Iter() {
    result.Add(item)
  }

  return result
}

func (s *Set) Intersection(other *Set) *Set {
  result := New()
  for item := range s.Iter() {
    if other.Includes(item) {
      result.Add(item)
    }
  }

  return result
}

func (s *Set) Diff(other *Set) *Set {
  result := New()

  for item := range s.Iter() {
    if !other.Includes(item) {
      result.Add(item)
    }
  }

  return result
}

func (s *Set) SymDiff(other *Set) *Set {
  result := New()

  for item := range s.Iter() {
    if !other.Includes(item) {
      result.Add(item)
    }
  }

  for item := range other.Iter() {
    if !s.Includes(item) {
      result.Add(item)
    }
  }

  return result
}

func (s *Set) IsSubsetOf(other *Set) bool {
  for item := range s.Iter() {
    if !other.Includes(item) {
      return false
    }
  }

  return true
}

func (s *Set) IsSupersetOf(other *Set) bool {
  return other.IsSubsetOf(s)
}

func (s *Set) String() string {
  var buffer bytes.Buffer
  size := s.Size()
  i := 0

  buffer.WriteString("{")

  for value := range s.Iter() {
    buffer.WriteString(fmt.Sprintf("%v", value))
    if i < size - 1 {
      buffer.WriteString(", ")
    }
    i++
  }

  buffer.WriteString("}")

  return buffer.String()
}

func (s *Set) Iter() <-chan interface{} {
  ch := make(chan interface{})

  go func(ch chan interface{}) {
    for item, _ := range s.table {
      ch <- item
    }
    close(ch)
  }(ch)

  return ch
}
