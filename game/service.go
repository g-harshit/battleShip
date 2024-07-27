package game

type Service interface {
	InitGame(n int) error
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
