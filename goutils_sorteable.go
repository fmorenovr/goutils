package goutils

import "sort"

// Sort sorts values (in-place) with respect to the given comparator.
func Sort(values []interface{}, comparator TypeComparator) {
  sort.Sort(sortable{values, comparator})
}

type sortable struct {
  values     []interface{}
  comparator TypeComparator
}

func (s sortable) Len() int {
  return len(s.values)
}

func (s sortable) Swap(i, j int) {
  s.values[i], s.values[j] = s.values[j], s.values[i]
}

func (s sortable) Less(i, j int) bool {
  return s.comparator(s.values[i], s.values[j]) < 0
}
