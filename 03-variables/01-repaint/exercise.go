package repaint

import "fmt"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// repaintColor returns the complementary of the color received as argument or an error for an unknown error.
func repaintColor(color string) (string, error) {
	if color == "teal" {
		return "vermilion", nil
	}
	if color == "vermilion" {
		return "teal", nil
	}
	return "", fmt.Errorf("unknown color: %s", color)
}
