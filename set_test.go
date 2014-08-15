package set

import "testing"

func Test_New(t *testing.T) {
  s := New()

  if s.Size() != 0 {
    t.Error("New set should be of size zero")
  }
}

func Test_Add(t *testing.T) {
  s := New()
  s.Add(1)
  s.Add("two")

  if s.Size() != 2 {
    t.Errorf("Expected size to be equals 2, but actual size is %d", s.Size())
  }
}

func Test_AddMultiple(t *testing.T) {
  s := New()
  s.Add(1, 2)

  if s.Size() != 2 {
    t.Errorf("Expected size to be equals 2, but actual size is %d", s.Size())
  }
}

func Test_AddDuplicates(t *testing.T) {
  s := New()
  s.Add(1, 2, 2, 3, 3)

  if s.Size() != 3 {
    t.Errorf("Expected size to be equals 3, but actual size is %d", s.Size())
  }
}

func Test_Remove(t *testing.T) {
  s := New()
  s.Add(1, 2, 3)
  s.Remove(2)

  if s.Size() != 2 {
    t.Fail()
  }

  if s.Includes(2) {
    t.Errorf("Expected %v to not include 2", s)
  }
}

func Test_Includes(t *testing.T) {
  s := New()
  s.Add(1, 2)

  if !s.Includes(1) {
    t.Errorf("Expected %v to include 1", s)
  }

  if !s.Includes(2) {
    t.Errorf("Expected %v to include 2", s)
  }
}

func Test_Size(t *testing.T) {
  s := New()
  s.Add(1)

  if s.Size() != 1 {
    t.Errorf("Expected size is 1, but acual is %d", s.Size())
  }

  s.Add(2)

  if s.Size() != 2 {
    t.Errorf("Expected size is 2, but acual is %d", s.Size())
  }

  s.Remove(2)

  if s.Size() != 1 {
    t.Errorf("Expected size is 1, but acual is %d", s.Size())
  }
}

func Test_Clear(t *testing.T) {
  s := New()
  s.Add(1, 2)

  s.Clear()

  if s.Size() != 0 {
    t.Errorf("Size should be 0, but actual size is %d", s.Size())
  }
}

func Test_String(t *testing.T) {
  s := New()

  if s.String() != "{}" {
    t.Errorf("String representation should be \"{}\", but actual is %s", s.String())
  }

  s.Add(1, "two")

  str := s.String()
  // set is unordered, so need to check both cases
  if str != "{1, two}" && str != "{two, 1}" {
    t.Errorf("String representation should be \"{1, two}\", but actual is %s", str)
  }
}

func Test_Equals(t *testing.T) {
  a := New()
  a.Add(1, 2)

  b := New()
  b.Add(1, 2)

  if !a.Equals(b) {
    t.Error("Expected sets a and b to be equals, but they're not")
  }

  if !b.Equals(a) {
    t.Error("Expected sets a and b to be equals, but they're not")
  }

  b.Add(3)

  if a.Equals(b) {
    t.Error("Expected sets a and b not to be equals, but they are")
  }

  if b.Equals(a) {
    t.Error("Expected sets a and b not to be equals, but they are")
  }

  b.Remove(2)

  if a.Equals(b) {
    t.Error("Expected sets a and b not to be equals, but they are")
  }

  if b.Equals(a) {
    t.Error("Expected sets a and b not to be equals, but they are")
  }
}

func Test_Dup(t *testing.T) {
  s := New()
  s.Add(1)
  clone := s.Dup()

  if !s.Equals(clone) {
    t.Errorf("Expected %v to equals its clone %v", s, clone)
  }
}

func Test_Union(t *testing.T) {
  a := New()
  b := New()

  a.Add(1)
  b.Add(2)

  union := a.Union(b)

  result := New()
  result.Add(1, 2)

  if !union.Equals(result) {
    t.Errorf("Expected union of %v and %v equals %v", a, b, result)
  }
}

func Test_UnionSame(t *testing.T) {
  a := New()
  b := New()

  a.Add(1)
  b.Add(1)

  union := a.Union(b)

  result := New()
  result.Add(1)

  if !union.Equals(result) {
    t.Errorf("Expected union of %v and %v equals %v", a, b, result)
  }
}

func Test_UnionWithEmpty(t *testing.T) {
  a := New()
  b := New()

  a.Add(1)

  if !a.Union(b).Equals(a) {
    t.Errorf("Expected union of %v and %v equals %v", a, b, a)
  }
}

func Test_Intersection(t *testing.T) {
  a := New()
  b := New()

  a.Add(1, 2)

  b.Add(2, 3)

  result := New()
  result.Add(2)

  if !a.Intersection(b).Equals(result) {
    t.Errorf("Expected intersection of %v and %v equals %v", a, b, result)
  }
}

func Test_IntersectionWithEmpty(t *testing.T) {
  a := New()
  b := New()

  a.Add(1, 2)

  if !a.Intersection(b).Equals(New()) {
    t.Errorf("Expected intersection of %v and %v to be empty", a, b)
  }
}

func Test_Diff(t *testing.T) {
  a := New()
  b := New()

  a.Add(1, 2)
  b.Add(2, 3)

  result := New()
  result.Add(1)

  if !a.Diff(b).Equals(result) {
    t.Errorf("Expected diff of %v and %v equals %v", a, b, result)
  }
}

func Test_SymDiff(t *testing.T) {
  a := New()
  b := New()

  a.Add(1, 2)
  b.Add(2, 3)

  result := New()
  result.Add(1, 3)

  if !a.SymDiff(b).Equals(result) {
    t.Errorf("Expected symdiff of %v and %v equals %v", a, b, result)
  }
}

func Test_SymDiffWithEmpty(t *testing.T) {
  a := New()
  b := New()

  a.Add(1)

  if !a.SymDiff(b).Equals(a) {
    t.Errorf("Expected symdiff of %v and %v equals %v", a, b, a)
  }
}

func Test_IsSubsetOf(t *testing.T) {
  a := New()
  b := New()

  a.Add(1, 2)
  b.Add(1, 2, 3, 4)

  if !a.IsSubsetOf(b) {
    t.Errorf("Expected %v to be a subset of %v", a, b)
  }

  b.Remove(1)

  if a.IsSubsetOf(b) {
    t.Errorf("Expected %v to not be a subset of %v", a, b)
  }
}

func Test_IsSupersetOf(t *testing.T) {
  a := New()
  b := New()

  a.Add(1, 2)
  b.Add(1, 2, 3, 4)

  if !b.IsSupersetOf(a) {
    t.Errorf("Expected %v to be a superset of %v", b, a)
  }

  b.Remove(1)

  if b.IsSupersetOf(a) {
    t.Errorf("Expected %v to not be a superset of %v", b, a)
  }
}

func Test_Iter(t *testing.T) {
  s := New()
  s.Add(1, 2)

  seenTable := make(map[int]bool)
  seenTable[1] = false
  seenTable[2] = false

  times := 0

  for item := range s.Iter() {
    intValue := item.(int)
    seenTable[intValue] = true
    times++
  }

  if times != s.Size() {
    t.Errorf("Expected to iterate %d times, but iterated %d times", s.Size(), times)
  }

  if !(seenTable[1] && seenTable[2]) {
    t.Error("Iterator did not iterated through all values")
  }
}
