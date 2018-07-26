package goutils

// Functions for Ordered Containers that index doesnt work in these values
type EnumerableIndex interface {
  Each(func(index int, value interface{}))

  // Map invokes the given function once for each element and returns a
  // container containing the values returned by the given function.
  // TODO would appreciate help on how to enforce this in containers (don't want to type assert when chaining)
  // Map(func(index int, value interface{}) interface{}) Container

  // Select returns a new container containing all elements for which the given function returns a true value.
  // TODO need help on how to enforce this in containers (don't want to type assert when chaining)
  // Select(func(index int, value interface{}) bool) Container

  Any(func(index int, value interface{}) bool) bool
  All(func(index int, value interface{}) bool) bool
  Find(func(index int, value interface{}) bool) (int, interface{})
}

// Functions for Ordered Containers that index doesnt work in these key/values pairs
type EnumerableKey interface {
  Each(func(key interface{}, value interface{}))

  // Map invokes the given function once for each element and returns a container
  // containing the values returned by the given function as key/value pairs.
  // TODO need help on how to enforce this in containers (don't want to type assert when chaining)
  // Map(func(key interface{}, value interface{}) (interface{}, interface{})) Container

  // Select returns a new container containing all elements for which the given function returns a true value.
  // TODO need help on how to enforce this in containers (don't want to type assert when chaining)
  // Select(func(key interface{}, value interface{}) bool) Container

  Any(func(key interface{}, value interface{}) bool) bool
  All(func(key interface{}, value interface{}) bool) bool
  Find(func(key interface{}, value interface{}) bool) (interface{}, interface{})
}
