package stratergy

type ValidationService interface {
	Validate(grid [][]string, cordinatesA, cordinatesB map[string]struct{}) error
	ValidateOverlapping(cordinatesA, cordinatesB map[string]struct{}) error
}
