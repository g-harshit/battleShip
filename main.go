package main

import (
	"fmt"

	"github.com/battleShip/game"
)

func main() {

	bs, err := game.InitGame(6)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = bs.AddShip("SH1", 2, 1, 1, 1, 4); err != nil {
		fmt.Println(err)
		return
	}

	bs.ViewBattleField()

	if err = bs.AddShip("SH2", 2, 4, 0, 5, 5); err != nil {
		fmt.Println(err)
		return
	}

	bs.ViewBattleField()

	if err = bs.StartGame(); err != nil {
		fmt.Println(err)
		return
	}
}
