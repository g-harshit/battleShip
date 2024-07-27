package game

type Service interface {
	AddShip(id string, size, xa, ya, xb, yb int) error
	ViewBattleField()
	StartGame() error
}

var defaultObject Service

func Init(obj Service) {
	defaultObject = obj
}

func GetService() Service {
	return defaultObject
}
