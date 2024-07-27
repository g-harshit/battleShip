package lib

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	PlayerA = "A"
	PlayerB = "B"
)

var Names = []string{PlayerA, PlayerB}

func GetCordinates(coord string) (int, int) {
	parts := strings.Split(coord, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return x, y
}

func MakeCoordinateString(x, y int) string {
	return fmt.Sprintf("%v,%v", x, y)
}
