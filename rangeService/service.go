package rangeService

type RangeService interface {
	GetPlayerRange(string) map[string]struct{}
	GetRandomCordinates(player string) (int, int, error)
}

type service RangeService

func NewRangeServiceFactory(n, playerSize int) service {
	two := NewTwoPlayerRangeService(n)
	switch playerSize {
	case 2:
		return two
	}
	return nil
}
