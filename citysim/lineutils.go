package main

const (
	lineOneIsSubline int = iota

	lineTwoIsSubline = iota

	identicalLines = iota

	noIntersect = iota
)

func computeLinesRelation(lineOne [][]float64, lineTwo [][]float64) int {

	// first check the lines lenght
	if (len(lineOne) == 0) || (len(lineTwo) == 0) {
		return noIntersect
	}

	// first step is to assess if lines are entirely equal
	if areSequencesEqual(lineOne, lineTwo) {
		return identicalLines
	}

	if checkLineIsSubline(lineOne, lineTwo) {
		return lineOneIsSubline
	}

	if checkLineIsSubline(lineTwo, lineOne) {
		return lineTwoIsSubline
	}

	return noIntersect
}

func checkLineIsSubline(subLine [][]float64, line [][]float64) bool {

	if len(subLine) >= len(line) {
		return false
	}

	for i := 0; i < len(line); i++ {
		seqAreEq := areSequencesEqual(subLine, line[i:i+len(subLine)])
		if seqAreEq {
			return true
		}
	}

	return false
}

func areSequencesEqual(s1 [][]float64, s2 [][]float64) bool {

	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {

		pointsAreEq := arePointsEqual(s1[i], s2[i])

		if !pointsAreEq {
			return false
		}
	}

	return true
}

func arePointsEqual(p1 []float64, p2 []float64) bool {
	if len(p1) != len(p2) {
		return false
	}

	for i := 0; i < len(p1); i++ {
		if p1[i] != p2[i] {
			return false
		}
	}

	return true
}
