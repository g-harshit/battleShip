package game

import (
	"errors"
	"fmt"

	"github.com/battleShip/lib"
	"github.com/battleShip/player"
	"github.com/battleShip/ship"
)

type battleShip struct {
	numerOfPlayer       int
	n                   int
	battleField         [][]string
	players             []*player.Player
	gameStarted         bool
	shipNameMap         map[string]*player.Player
	getPlayerRanges     func(player string) map[string]struct{}
	getRandonRanges     func(player string) (int, int, error)
	validateStatergy    func(grid [][]string, cordinatesA, validRange map[string]struct{}) error
	validateOverlapping func(cordinatesA, cordinatesB map[string]struct{}) error
}

func InitGame(n int) (*battleShip, error) {
	b := &battleShip{
		numerOfPlayer: 2,
		shipNameMap:   make(map[string]*player.Player),
		gameStarted:   false,
	}
	err := b.validateGridSize(n)
	if err != nil {
		return nil, err
	}
	b.n = n
	b.initDependency()
	b.initGrid()
	b.initPlayers()
	return b, nil
}

func (b *battleShip) AddShip(id string, size, xa, ya, xb, yb int) error {
	if b.gameStarted {
		return errors.New("game has already started")
	}
	shipA := ship.NewShip(id, size, xa, ya)
	shipB := ship.NewShip(id, size, xb, yb)

	err := b.validateOverlapping(shipA.GetCoordinates(), shipB.GetCoordinates())
	if err != nil {
		return err
	}

	if err = b.validateStatergy(b.battleField, shipA.GetCoordinates(), b.getPlayerRanges(lib.PlayerA)); err != nil {
		return err
	}
	if err = b.validateStatergy(b.battleField, shipB.GetCoordinates(), b.getPlayerRanges(lib.PlayerB)); err != nil {
		return err
	}

	b.players[0].AddShip(shipA)
	b.players[1].AddShip(shipB)
	b.setGridForNewShip(shipA, lib.PlayerA, id)
	b.setGridForNewShip(shipB, lib.PlayerB, id)
	return nil
}

func (b *battleShip) ViewBattleField() {
	fmt.Println("BattleField :")
	for i := 0; i < b.n; i++ {
		for j := 0; j < b.n; j++ {
			fmt.Printf("%v \t", b.battleField[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func (b *battleShip) StartGame() error {
	gameOver := false
	i := 0
	for !gameOver {
		player := b.getPlayingPlayer(&i)
		x, y, err := b.getRandonRanges(player.GetName())
		if err != nil {
			return err
		}
		hitPlayerName, shipName := b.getPlayerAndShipName(b.battleField[x][y])
		if hitPlayerName != "" && shipName != "" {
			hitPlayer := b.shipNameMap[hitPlayerName]
			hitPlayer.MarkShipHit(shipName)
			fmt.Println(player.GetName(), "has HIT target at", x, y)
			fmt.Println()
			b.removeShipFromGrid(b.battleField[x][y])
			if hitPlayer.IsAllShipDead() {
				gameOver = true
				fmt.Println("*******************", b.getWinningStatement(player), "*******************")
				fmt.Println()
			}
		} else {
			fmt.Println(player.GetName(), "has Missed target at", x, y)
			fmt.Println()
		}
		b.ViewBattleField()
	}
	return nil
}
