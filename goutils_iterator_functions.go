package goutils

// Iterator to browse with index
type IteratorIndex interface {
  Next() bool
  Value() interface{}
  Index() int
  Begin()
  First() bool
}

// ReverseIterator to browse with index
type ReverseIteratorIndex interface {
  Prev() bool
  End()
  Last() bool
  IteratorIndex
}

// Iterator to browse with key-value pairs.
type IteratorKey interface {
  Next() bool
  Value() interface{}
  Key() interface{}
  Begin()
  First() bool
}

// ReverseIterator to browse with key-value pairs.
type ReverseIteratorKey interface {
  Prev() bool
  End()
  Last() bool
  IteratorKey
}
