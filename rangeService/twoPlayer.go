package rangeService

import (
	"errors"

	"github.com/battleShip/lib"
)

type TwoPlayer struct {
	playerRange map[string]map[string]struct{}
}

func NewTwoPlayerRangeService(n int) *TwoPlayer {
	playerRange := make(map[string]map[string]struct{})
	playerRange[lib.PlayerA] = addRageForPlayerA(n)
	playerRange[lib.PlayerB] = addRageForPlayerB(n)
	return &TwoPlayer{
		playerRange: playerRange,
	}
}

func (d *TwoPlayer) GetPlayerRange() map[string]map[string]struct{} {
	return d.playerRange
}

func addRageForPlayerA(n int) map[string]struct{} {
	resp := make(map[string]struct{})
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			resp[lib.MakeCoordinateString(i, j)] = struct{}{}
		}
	}
	return resp
}

func addRageForPlayerB(n int) map[string]struct{} {
	resp := make(map[string]struct{})
	for i := 0; i < n; i++ {
		for j := n / 2; j < n; j++ {
			resp[lib.MakeCoordinateString(i, j)] = struct{}{}
		}
	}
	return resp
}

func (d *TwoPlayer) GetRandomCordinates(player string) (int, int, error) {
	ranges := d.getSearchRange(player)
	if len(ranges) == 0 {
		return 0, 0, errors.New("no rang left")
	}
	coords := ""
	for key := range ranges {
		coords = key
		break
	}
	delete(ranges, coords)
	x, y := lib.GetCordinates(coords)
	return x, y, nil
}

func (d *TwoPlayer) getSearchRange(player string) map[string]struct{} {
	if player == lib.PlayerA {
		return d.playerRange[lib.PlayerB]
	}
	return d.playerRange[lib.PlayerA]
}
