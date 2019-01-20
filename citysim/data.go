package main

import (
	"container/list"
)

var topLevelList list.List

type _LineData struct {
	line   [][]float64
	actors int
}

func getLineData(line [][]float64) int {
	return 0
}

func deliverLineData(line [][]float64) {

	fList, fElem := findListAndElement(line)

	newLineData := _LineData{
		line:   line,
		actors: 1}

	if fList == nil {
		newList := list.New()
		topLevelList.PushFront(newList)
		newList.PushFront(newLineData)
	} else {
		if fElem == nil {
			fList.PushFront(newLineData)
		}
		fList.InsertAfter(newLineData, fElem)
	}
}

func findListAndElement(line [][]float64) (*list.List, *list.Element) {

	for subListE := topLevelList.Front(); subListE != nil; subListE = subListE.Next() {

		list := subListE.Value.(*list.List)
		elem := findElement(list, line)

		if elem != nil {
			return list, elem
		}
	}

	return nil, nil
}

func findElement(list *list.List, line [][]float64) *list.Element {

	for lineDataE := list.Front(); lineDataE != nil; lineDataE = lineDataE.Next() {

		lineData := lineDataE.Value.(_LineData)
		linesResult := computeLinesRelation(line, lineData.line)

		switch linesResult {

		case lineOneIsSubline:
			nextLineDataE := lineDataE.Next()
			if nextLineDataE == nil {
				continue
			}
			nextLineData := nextLineDataE.Value.(_LineData)
			nextLinesResult := computeLinesRelation(line, nextLineData.line)
			if nextLinesResult == lineTwoIsSubline {
				return lineDataE
			}

		case lineTwoIsSubline:

		case identicalLines:

		case noIntersect:
		}
	}

	return nil
}
