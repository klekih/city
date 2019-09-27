package main

import (
	"fmt"
	"os"
	"time"

	"github.com/tomagb/city/common"
)

var ticker *time.Ticker
var myRoute *Route
var currentInstructionIndex = -1
var currentPosInInstruction float64
var averageSpeed float64 = 10 // meter/second

// WalkerStatus is the answer for reporting walker status
type WalkerStatus struct {
	arrivedAtDestination bool
}

// StartWalker is the main entry point for starting a walker
func StartWalker(config *Config, city *common.CityInterface) chan WalkerStatus {

	// first, get a route to go on
	myRoute = getRoute(config)

	if len(myRoute.Paths) == 0 {
		panic("Route has no path to go on")
	}

	fmt.Println("Got route, going on it")

	reportChan, lineChan := Connect()

	// create the chan to report back the status
	c := make(chan WalkerStatus)

	duration := time.Duration(config.SimulationStep)
	ticker = time.NewTicker(duration * time.Millisecond)

	go func() {
		for {
			select {
			case <-ticker.C:
				advance(city, reportChan, lineChan)
			case j := <-lineChan:
				fmt.Println("Received answer with line", j)
			}
		}
	}()

	return c
}

func getRoute(config *Config) *Route {

	route, err := generateRandomRoute(config)

	if err != nil {
		panic(err)
	}

	return route
}

func advance(city *common.CityInterface, chanReport chan common.Report, lineChan chan common.LineInfo) {

	if myRoute == nil {
		panic("No route to go on")
	}

	if currentInstructionIndex == -1 {
		currentInstructionIndex = 1

		currentInstruction := myRoute.Paths[0].Instructions[currentInstructionIndex]
		firstPointIndex := currentInstruction.Interval[0]
		secondPointIndex := currentInstruction.Interval[1]
		lineToEnter := myRoute.Paths[0].Points.Coordinates[firstPointIndex:secondPointIndex]

		enterReport := common.Report{}.
			WithCurrentLine(lineToEnter).
			WithReportDetails(common.ReportOnTheLine)

		chanReport <- enterReport
	}

	// Check if all instructions have been passed.
	if currentInstructionIndex >= len(myRoute.Paths[0].Instructions) {
		fmt.Println("Route finished")
		os.Exit(0)
	}

	// Get the current distance to cover
	currentInstruction := myRoute.Paths[0].Instructions[currentInstructionIndex]
	distance := currentInstruction.Distance

	currentPosInInstruction += averageSpeed

	if currentPosInInstruction >= distance {

		offFirstPointIndex := currentInstruction.Interval[0]
		offSecondPointIndex := currentInstruction.Interval[1]
		lineToExit := myRoute.Paths[0].Points.Coordinates[offFirstPointIndex:offSecondPointIndex]

		exitReport := common.Report{}.
			WithCurrentLine(lineToExit).
			WithReportDetails(common.ReportOffFromLine)
		chanReport <- exitReport

		currentInstructionIndex++
		if currentInstructionIndex >= len(myRoute.Paths[0].Instructions) {
			return
		}
		currentInstruction = myRoute.Paths[0].Instructions[currentInstructionIndex]
		onFirstPointIndex := currentInstruction.Interval[0]
		onSecondPointIndex := currentInstruction.Interval[1]
		lineToEnter := myRoute.Paths[0].Points.Coordinates[onFirstPointIndex:onSecondPointIndex]

		enterReport := common.Report{}.
			WithCurrentLine(lineToEnter).
			WithReportDetails(common.ReportOnTheLine)
		chanReport <- enterReport

		currentPosInInstruction = 0
	} else {
		fmt.Println("On", currentInstruction.StreetName, ":",
			currentPosInInstruction, "of", distance)
	}

	firstPointIndex := currentInstruction.Interval[0]
	secondPointIndex := currentInstruction.Interval[1]

	currentLine := myRoute.Paths[0].Points.Coordinates[firstPointIndex:secondPointIndex]

	lineChan <- common.LineInfo{}.
		WithLine(currentLine)
}
