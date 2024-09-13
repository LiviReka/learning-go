package logicalops

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

func inverse(b bool) bool {
	return !b
}

func and(x, y bool) bool {
	return x && y
}

func deMorgan(a, b bool) bool {
	//return !(a || b) == (!a && !b)
	return and(inverse(a), inverse(b))
}
