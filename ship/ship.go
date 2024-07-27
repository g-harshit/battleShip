package ship

import (
	"github.com/battleShip/lib"
)

type Ship struct {
	name        string
	size        int
	x, y        int
	coordinates map[string]struct{}
	alive       bool
}

func NewShip(name string, size, x, y int) *Ship {
	coordinates := make(map[string]struct{})
	rowStart := getStartingPoint(size, x)
	colStart := getStartingPoint(size, y)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			coord := lib.MakeCoordinateString(rowStart+i, colStart+j)
			coordinates[coord] = struct{}{}
		}
	}
	return &Ship{name, size, x, y, coordinates, true}
}

func getStartingPoint(size, x int) int {
	start := x - (size / 2)
	if start < 0 {
		return 0
	}
	return start
}

func (s *Ship) GetCoordinates() map[string]struct{} {
	return s.coordinates
}

func (s *Ship) IsAlive() bool {
	return s.alive
}

func (s *Ship) SetAlive(alive bool) {
	s.alive = alive
}

func (s *Ship) GetName() string {
	return s.name
}
