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
	allRanges           map[string]map[string]struct{}
	getRandonRanges     func(player string) (int, int, error)
	validateStatergy    func(grid [][]string, cordinatesA, validRange map[string]struct{}) error
	validateOverlapping func(cordinatesA, cordinatesB map[string]struct{}) error
}

func NewBattleShip(numerOfPlayer int) *battleShip {
	return &battleShip{
		numerOfPlayer: 2,
		gameStarted:   false,
		shipNameMap:   make(map[string]*player.Player),
	}
}

func (b *battleShip) InitGame(n int) error {
	err := b.validateGridSize(n)
	if err != nil {
		return err
	}
	b.n = n
	b.initDependency()
	b.battleField = make([][]string, n)
	for i := range b.battleField {
		b.battleField[i] = make([]string, n)
		for j := range b.battleField[i] {
			b.battleField[i][j] = "."
		}
	}

	for i := 0; i < b.numerOfPlayer; i++ {
		player := player.NewPlayer(lib.Names[i])
		b.shipNameMap[lib.Names[i]] = player
		b.players = append(b.players, player)
	}
	b.gameStarted = false
	return nil
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

	if err = b.validateStatergy(b.battleField, shipA.GetCoordinates(), b.allRanges[lib.PlayerA]); err != nil {
		return err
	}
	if err = b.validateStatergy(b.battleField, shipB.GetCoordinates(), b.allRanges[lib.PlayerB]); err != nil {
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
			b.removeShipFromGrid(b.battleField[x][y])
			if hitPlayer.IsAllShipDead() {
				gameOver = true
				fmt.Println("*******************", b.getWinningStatement(player), "*******************")
			}
		} else {
			fmt.Println(player.GetName(), "has Missed target at", x, y)
		}
		b.ViewBattleField()
	}
	return nil
}
