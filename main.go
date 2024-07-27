package main

import (
	"fmt"

	"github.com/battleShip/game"
)

func main() {
	game.Init(game.NewBattleShip(2))

	game := game.GetService()

	err := game.InitGame(6)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = game.AddShip("SH1", 2, 1, 1, 1, 4); err != nil {
		fmt.Println(err)
		return
	}

	if err = game.AddShip("SH2", 2, 5, 0, 5, 5); err != nil {
		fmt.Println(err)
	}

	game.ViewBattleField()

	if err = game.StartGame(); err != nil {
		fmt.Println(err)
		return
	}
}
