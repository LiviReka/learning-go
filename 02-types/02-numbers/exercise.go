package calculator

import (
	"math"
	"strconv"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

func amean(x, y float64) int {
	return int(math.Round((x + y) / 2))
}

func ameanString(x, y string) (int, error) {
	newX, err := strconv.ParseFloat(x, 64)
	if err != nil {
		return 0, err
	}
	newY, err := strconv.ParseFloat(y, 64)
	if err != nil {
		return 0, err
	}
	return amean(newX, newY), nil
}
