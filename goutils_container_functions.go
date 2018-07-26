package goutils

// Principal container description for Data Structures
type Container interface {
  IsEmpty() bool
  Size() int
  Clear()
  Values() []interface{}
}
