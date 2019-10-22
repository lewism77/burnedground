package vector

import "testing"

// Add adds together two vectors and returns the result
func TestAdd(t *testing.T) {
	v1 := Vector{X: 1.25, Y: 2.25}
	v2 := Vector{X: 1.0, Y: 2.5}

	got := Add(v1, v2)

	if got.X != 2.25 || got.Y != 4.75 {
		t.Errorf("Add v1 v2 failed")
	}
}
