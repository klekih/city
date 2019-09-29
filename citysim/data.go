package main

import (
	"container/list"
)

var topLevelList *list.List

type _LineData struct {
	line   [][]float64
	actors int
}

func initData() {
	if topLevelList == nil {
		topLevelList = list.New()
	}
}

func newData() {
	if topLevelList != nil {
		topLevelList.Init()
	}
}

func deliverLineData(line [][]float64) int {

	initData()

	for elem := topLevelList.Front(); elem != nil; elem = elem.Next() {
		elemLine := elem.Value.(_LineData)
		result := computeLinesRelation(line, elemLine.line)
		if result == identicalLines {
			elemLine.actors++
			elem.Value = elemLine
			return elemLine.actors
		}
	}

	newLineData := _LineData{line, 1}
	topLevelList.PushBack(newLineData)
	return 1
}

func getLineData(line [][]float64) int {

	if topLevelList == nil {
		return 0
	}

	totalActors := 0

	for elem := topLevelList.Front(); elem != nil; elem = elem.Next() {
		elemLine := elem.Value.(_LineData)
		result := computeLinesRelation(line, elemLine.line)
		switch result {
		case identicalLines:
			fallthrough
		case lineOneIsSubline:
			totalActors += elemLine.actors
		}
	}

	return totalActors
}

func deleteLineData(line [][]float64) int {

	if topLevelList == nil {
		return 0
	}

	for elem := topLevelList.Front(); elem != nil; elem = elem.Next() {
		elemLine := elem.Value.(_LineData)
		result := computeLinesRelation(line, elemLine.line)
		if result == identicalLines {
			if elemLine.actors == 1 {
				topLevelList.Remove(elem)
				return 0
			} else {
				elemLine.actors--
				return elemLine.actors
			}
		}
	}

	return 0
}
