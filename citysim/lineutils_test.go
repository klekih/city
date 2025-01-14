package main

import "testing"

func TestFirstLineIsSublineInMiddle(t *testing.T) {
	lineOne := [][]float64{
		{2, 6},
		{3, 7}}

	lineTwo := [][]float64{
		{1, 5},
		{2, 6},
		{3, 7},
		{4, 8},
		{5, 9}}

	result := computeLinesRelation(lineOne, lineTwo)

	if result != lineOneIsSubline {
		t.Fail()
	}
}

func TestFirstLineIsSublineAtBeginning(t *testing.T) {
	lineOne := [][]float64{
		{1, 5},
		{2, 6}}

	lineTwo := [][]float64{
		{1, 5},
		{2, 6},
		{3, 7},
		{4, 8},
		{5, 9}}

	result := computeLinesRelation(lineOne, lineTwo)

	if result != lineOneIsSubline {
		t.Fail()
	}
}

func TestFirstLineIsSublineAtEnd(t *testing.T) {
	lineOne := [][]float64{
		{4, 8},
		{5, 9}}

	lineTwo := [][]float64{
		{1, 5},
		{2, 6},
		{3, 7},
		{4, 8},
		{5, 9}}

	result := computeLinesRelation(lineOne, lineTwo)

	if result != lineOneIsSubline {
		t.Fail()
	}
}

func TestFirstLineIsSublineWithMoreCoordinates(t *testing.T) {
	lineOne := [][]float64{
		{2, 6, 10},
		{3, 7, 11}}

	lineTwo := [][]float64{
		{1, 5, 9},
		{2, 6, 10},
		{3, 7, 11},
		{4, 8, 12},
		{5, 9, 13}}

	result := computeLinesRelation(lineOne, lineTwo)

	if result != lineOneIsSubline {
		t.Fail()
	}
}

func TestSecondLineIsSubline(t *testing.T) {

	lineOne := [][]float64{
		{1, 5},
		{2, 6},
		{3, 7},
		{4, 8},
		{5, 9}}

	lineTwo := [][]float64{
		{2, 6},
		{3, 7}}

	result := computeLinesRelation(lineOne, lineTwo)

	if result != lineTwoIsSubline {
		t.Fail()
	}
}

func TestIdenticalLines(t *testing.T) {

	lineOne := [][]float64{
		{1, 5},
		{2, 6},
		{3, 7},
		{4, 8},
		{5, 9}}

	lineTwo := [][]float64{
		{1, 5},
		{2, 6},
		{3, 7},
		{4, 8},
		{5, 9}}

	result := computeLinesRelation(lineOne, lineTwo)

	if result != identicalLines {
		t.Fail()
	}
}

func TestNoIntersect(t *testing.T) {

	lineOne := [][]float64{
		{11, 51},
		{21, 61},
		{31, 71},
		{41, 81},
		{51, 91}}

	lineTwo := [][]float64{
		{1, 5},
		{2, 6},
		{3, 7},
		{4, 8},
		{5, 9}}

	result := computeLinesRelation(lineOne, lineTwo)

	if result != noIntersect {
		t.Fail()
	}
}

func TestWithIntermitentIntersect(t *testing.T) {

	lineOne := [][]float64{
		{1, 5},
		{2, 6},
		{31, 71},
		{4, 8},
		{5, 9}}

	lineTwo := [][]float64{
		{1, 5},
		{2, 6},
		{3, 7},
		{4, 8},
		{5, 9}}

	result := computeLinesRelation(lineOne, lineTwo)

	if result != noIntersect {
		t.Fail()
	}
}
