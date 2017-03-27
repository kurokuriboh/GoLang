/*
Model to handle data of Candy Crush game
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// BoardInfo struct for json object
type BoardInfo struct {
	Rows    int
	Columns int
	Data    []int
}

// Definition struct for json object
type Definition struct {
	GameID         int
	ExtensionColor BoardInfo
	BoardState     BoardInfo
	Colors         int
}

// Message struct for json object
type Message struct {
	GameDef Definition
}

// CandyData struct for storing information of candy
type CandyData struct {
	Row      int
	Column   int
	Selected bool
}

var extArray [][]CandyData
var boardArray [][]CandyData
var offsetArray []int

var movesSoFar int
var currentScore int

/*
SetupModel sets up the game with the given file
*/
func SetupModel(filename string) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	var m Message
	err = json.Unmarshal(file, &m)
	if err != nil {
		fmt.Printf("Invalid Json: %v\n", err)
		os.Exit(1)
	}

	// TODO: Transfer values to boardArray and extArray

	fmt.Printf("%v", boardArray)
}
