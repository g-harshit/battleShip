package player

import (
	"github.com/battleShip/ship"
)

type Player struct {
	name          string
	ships         []*ship.Ship
	hitCordinates map[int]map[int]struct{}
}

func NewPlayer(name string) *Player {
	return &Player{
		name:          name,
		ships:         []*ship.Ship{},
		hitCordinates: make(map[int]map[int]struct{}),
	}
}

func (p *Player) AddShip(ship *ship.Ship) {
	p.ships = append(p.ships, ship)
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) IsAllShipDead() bool {
	for key := range p.ships {
		if p.ships[key].IsAlive() {
			return false
		}
	}
	return true
}

func (p *Player) MarkShipHit(name string) {
	for key := range p.ships {
		if p.ships[key].GetName() == name {
			p.ships[key].SetAlive(false)
			return
		}
	}
}

func (p *Player) IsAlreadyHit(x, y int) bool {
	_, exists := p.hitCordinates[x][y]
	return exists
}

func (p *Player) AddHitCoordinates(x, y int) {
	if _, exists := p.hitCordinates[x]; !exists {
		p.hitCordinates[x] = make(map[int]struct{})
	}
	p.hitCordinates[x][y] = struct{}{}
}
