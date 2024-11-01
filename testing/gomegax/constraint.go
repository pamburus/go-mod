package gomegax

import "testing"

// HierarchicalTest is a testing.TB with a Run method to run hierarchical tests.
type HierarchicalTest[T testing.TB] interface {
	testing.TB
	Run(name string, f func(T)) bool
}
