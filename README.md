## go-set

Simple library that implements Set data structure. Supports basic set operations like
union, intersection, difference, symetric difference.

### Usage
    import (
      "github.com/uzzz/go-set"
      "fmt"
    )

    func main() {
      s := set.New()

      s.Add(1)
      s.Add(2)
      s.Add(3, 4)

      fmt.Println(s.Includes(1)) // true
      fmt.Println(s.Includes(5)) // false

      other := set.New()
      other.Add(4, 5, 6)
      fmt.Println(s.Union(other)) // {1,2,3,4,5,6}
      fmt.Println(s.Intersection(other)) // {4}

      for item := range s.Iter() {
        fmt.Println(item)
      }
    }

