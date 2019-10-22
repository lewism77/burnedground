package vector

// Vector is a simple xy vector
type Vector struct{ X, Y float64 }

// Add adds together two vectors and returns the result
func Add(v1 Vector, v2 Vector) Vector {
	v1.X = v1.X + v2.X
	v1.Y = v1.Y + v2.Y
	return v1
}
