package goutils_test

import(
  "github.com/jenazads/goutils";
  "fmt";
)

func main() {
  strings := []interface{}{}                  // []
  strings = append(strings, "d")              // ["d"]
  strings = append(strings, "a")              // ["d","a"]
  strings = append(strings, "b")              // ["d","a",b"
  strings = append(strings, "c")              // ["d","a",b","c"]
  fmt.Println("Is Ordered ? ",strings)  
  goutils.Sort(strings, goutils.StringComparator) // ["a","b","c","d"]
  fmt.Println("Is Ordered ? ",strings)  
}
