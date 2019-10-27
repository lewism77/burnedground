package terrain

import "math/rand"

var Grid [][]int //TODO: Use pointer

// Init initialises an empty terrain grid based on a fixed size
func Init(width int, height int) {
	Grid = make([][]int, width)
	for i := range Grid {
		(Grid)[i] = make([]int, height)
	}
}

// Generate populates the terrain grid using an 'algorithm'
// First column is between 40% and 70% of screen height
// Subsequent columns are within 30% of the previous column
func Generate() {
	//width := len(Grid)
	height := len((Grid)[0])

	nextColumn := height * ((rand.Intn(30) + 40) / 100)

	for i := range Grid {
		for x := 0; x < nextColumn; x++ {
			(Grid)[i][x] = 1
		}
		nextColumn = nextColumn * ((rand.Intn(30) + 40) / 100) //TODO: Fix numbers
	}
}
