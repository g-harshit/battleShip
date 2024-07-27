package game

import (
	"errors"
	"strings"

	"github.com/battleShip/lib"
	"github.com/battleShip/player"
	"github.com/battleShip/rangeService"
	"github.com/battleShip/ship"
	"github.com/battleShip/stratergy"
)

func (b *battleShip) validateGridSize(n int) error {
	if n <= 0 {
		return errors.New("n should be greater than zero")
	}
	if n%2 != 0 {
		return errors.New("grid size should be even")
	}
	return nil
}

func (b *battleShip) setGridForNewShip(ship *ship.Ship, playerName, id string) {
	for coord := range ship.GetCoordinates() {
		x, y := lib.GetCordinates(coord)
		b.battleField[x][y] = playerName + "-" + id
	}
}

func (b *battleShip) getPlayerAndShipName(val string) (string, string) {
	parts := strings.Split(val, "-")
	if len(parts) != 2 {
		return "", ""
	}
	return parts[0], parts[1]
}

func (b *battleShip) getPlayingPlayer(i *int) *player.Player {
	player := b.players[*i]
	*i++
	if *i == b.numerOfPlayer {
		*i = 0
	}
	return player
}

func (b *battleShip) getWinningStatement(player *player.Player) string {
	return player.GetName() + " Has Won"
}

func (b *battleShip) initDependency() {
	st := stratergy.NewBasicStratergy(b.n)
	b.validateStatergy = st.Validate
	b.validateOverlapping = st.ValidateOverlapping
	rs := rangeService.NewRangeServiceFactory(b.n, 2)
	b.getRandonRanges = rs.GetRandomCordinates
	b.allRanges = rs.GetPlayerRange()
}

func (b *battleShip) removeShipFromGrid(cood string) {
	if cood == "." {
		return
	}
	for i := 0; i < b.n; i++ {
		for j := 0; j < b.n; j++ {
			if b.battleField[i][j] == cood {
				b.battleField[i][j] = "."
			}
		}
	}
}
