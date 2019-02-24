package main

import (
	"container/list"
)

var topLevelList *list.List

type _LineData struct {
	line   [][]float64
	actors int
}

func getLineData(line [][]float64) int {
	tll := findSubList(line)

	if tll == nil {
		return 0
	}

	result := lineOneIsSubline
	aggregateValue := 0

	for elem := tll.Front(); result == lineOneIsSubline && elem != nil; elem = elem.Next() {
		lineFromList := elem.Value.(_LineData)
		result = computeLinesRelation(line, lineFromList.line)
		aggregateValue += lineFromList.actors
	}

	return aggregateValue
}

func deliverLineData(line [][]float64) {

	if topLevelList == nil {
		topLevelList = list.New()
	}

	subList := findSubList(line)
	if subList == nil {
		// 1. create new sub-list and push it into the top level list
		newLineData := _LineData{
			line:   line,
			actors: 1}
		newList := list.New()
		topLevelList.PushFront(newList)
		// 2. push new line element into the new sub-list
		newList.PushFront(newLineData)

	} else {
		// there is a sub-list where the new line element fits into
		// so the element is inserted into its correct position
		addLineInSublist(subList, line)
	}
}

func findSubList(line [][]float64) *list.List {

	if topLevelList == nil {
		return nil
	}

	for subListElem := topLevelList.Front(); subListElem != nil; subListElem = subListElem.Next() {
		sublist := subListElem.Value.(*list.List)
		firstLine := sublist.Front().Value.(_LineData)
		linesResult := computeLinesRelation(line, firstLine.line)
		if linesResult == lineOneIsSubline || linesResult == identicalLines {
			return sublist
		}
	}
	return nil
}

func addLineInSublist(list *list.List, line [][]float64) {

	for elem := list.Front(); elem != nil; elem = elem.Next() {
		lineFromList := elem.Value.(_LineData)
		relationResult := computeLinesRelation(line, lineFromList.line)

		switch relationResult {
		case lineTwoIsSubline:
			{
				newLineData := _LineData{
					line:   line,
					actors: 1}
				list.InsertBefore(newLineData, elem)
				return
			}
		case lineOneIsSubline:
			{
				newLineData := _LineData{
					line:   line,
					actors: 1}
				list.InsertAfter(newLineData, elem)
				return
			}
		case identicalLines:
			{
				lineFromList.actors++
			}
		}
	}
}
