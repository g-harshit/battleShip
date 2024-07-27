package stratergy

import (
	"errors"

	"github.com/battleShip/lib"
)

type basic struct {
	n int
}

func NewBasicStratergy(n int) *basic {
	return &basic{
		n: n,
	}
}

func (b *basic) Validate(grid [][]string, cordinates, validRange map[string]struct{}) error {
	for coord := range cordinates {
		x, y := lib.GetCordinates(coord)
		if _, exists := validRange[coord]; !exists {
			return errors.New("not in valid range")
		}
		if x >= b.n || y >= b.n {
			return errors.New("ship is out of battle field")
		}
		if grid[x][y] != "." {
			return errors.New("ship overlapping grid")
		}
	}
	return nil
}

func (b *basic) ValidateOverlapping(cordinatesA, cordinatesB map[string]struct{}) error {
	for coord := range cordinatesA {
		if _, exists := cordinatesB[coord]; exists {
			return errors.New("overlapping ships")
		}
	}
	return nil
}
