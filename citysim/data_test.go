package main

import "testing"

func TestTopLevelListWithOneLine(t *testing.T) {
	line := [][]float64{
		{1, 2},
		{3, 4}}

	deliverLineData(line)

	data := getLineData(line)

	if data != 1 {
		t.Fail()
	}
}

func TestTopLevelWithTwoLines(t *testing.T) {
	line := [][]float64{
		{1, 2},
		{3, 4},
		{5, 6}}
	deliverLineData(line)

	line2 := [][]float64{
		{3, 4},
		{5, 6}}

	data := getLineData(line2)

	if data != 1 {
		t.Fail()
	}
}

func TestTopLevelWithTwoUnrelatedLines(t *testing.T) {
	line := [][]float64{
		{1, 2},
		{3, 4},
		{5, 6}}
	deliverLineData(line)

	line2 := [][]float64{
		{7, 8},
		{9, 10}}

	data := getLineData(line2)

	if data != 0 {
		t.Fail()
	}
}

func TestSubListWithTwoRelatedLines(t *testing.T) {
	line := [][]float64{
		{1, 2},
		{3, 4},
		{5, 6}}

	line2 := [][]float64{
		{3, 4},
		{5, 6}}

	deliverLineData(line)
	deliverLineData(line2)

	data := getLineData(line)
	if data != 1 {
		t.Fail()
	}

	data = getLineData(line2)
	if data != 2 {
		t.Fail()
	}
}

func TestSubListWithTwoDifferentLines(t *testing.T) {
	line := [][]float64{
		{1, 2},
		{3, 4},
		{5, 6}}

	line2 := [][]float64{
		{3, 4},
		{5, 6}}

	deliverLineData(line)
	deliverLineData(line2)

	data := getLineData(line)
	if data != 1 {
		t.Fail()
	}

	data = getLineData(line2)
	if data != 2 {
		t.Fail()
	}
}

func TestSubListWithThreeDifferentLines(t *testing.T) {
	line1 := [][]float64{
		{1, 2},
		{3, 4},
		{5, 6}}

	line2 := [][]float64{
		{3, 4},
		{5, 6}}

	line3 := [][]float64{
		{1, 2},
		{3, 4}}

	deliverLineData(line1)
	deliverLineData(line2)
	deliverLineData(line3)

	data := getLineData(line1)
	if data != 1 {
		t.Fail()
	}

	data = getLineData(line2)
	if data != 2 {
		t.Fail()
	}

	data = getLineData(line3)
	if data != 2 {
		t.Fail()
	}
}

func TestSubListWithThreeDifferentLines_2(t *testing.T) {
	line1 := [][]float64{
		{1, 2},
		{3, 4},
		{5, 6}}

	line2 := [][]float64{
		{3, 4},
		{5, 6}}

	line3 := [][]float64{
		{7, 8},
		{9, 10}}

	deliverLineData(line1)
	deliverLineData(line2)
	deliverLineData(line3)

	data := getLineData(line1)
	if data != 1 {
		t.Fail()
	}

	data = getLineData(line2)
	if data != 2 {
		t.Fail()
	}

	data = getLineData(line3)
	if data != 1 {
		t.Fail()
	}
}
