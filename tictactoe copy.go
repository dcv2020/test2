//testing git branch1
package main

import (
	"fmt"
	"time"
)

//test
//declaring global variables here
var board [21]string
var userMove int
var computerWin = false

//these are making sure the computer doesn't go twice in the same move
var secondMove = false
var thirdMove = false
var fourthMove = false

func main() {

	//beginning of game
	firstDrawBoard() //method puts board in terminal for the first time
	fmt.Println("Type in your move on the board (1-9) and press enter to start playing")

	fmt.Scanln(&userMove) //users first move entered here

	updateBoard(userMove) //changing user input to proper values

	drawBoard() //redrawing board with updated vales on it

	////////////////////////////////////////////////////////////////////////////////////////////////

	computerFirstMove() //logic for first computer move

	time.Sleep(1 * time.Second) //this is to have the computer wait a second before moving

	drawBoard() //drawing board with computers move on it

	////////////////////////////////////////////////////////////////////////////////////////////////

	fmt.Scanln(&userMove) //users second move entered here

	updateBoard(userMove)

	drawBoard()

	////////////////////////////////////////////////////////////////////////////////////////////////

	//computer second move here
	checkImmediateDanger() //this checks if the human if about to win

	if secondMove == false { //if human isn't about to win, then computer moves
		computerSecondMove()
	}

	time.Sleep(1 * time.Second)

	drawBoard()

	////////////////////////////////////////////////////////////////////////////////////////////////

	fmt.Scanln(&userMove) //users third move entered here

	updateBoard(userMove) //changing user input to proper values

	drawBoard() //redrawing board with updated vales on it

	////////////////////////////////////////////////////////////////////////////////////////////////

	//computer third move here

	checkImmediateWin()

	if thirdMove == false {
		checkImmediateDanger()
	}

	if thirdMove == false {
		computerNextMove()
	}

	time.Sleep(1 * time.Second)

	drawBoard()

	////////////////////////////////////////////////////////////////////////////////////////////////

	if computerWin == false { //don't play move if computer already won

		fmt.Scanln(&userMove) //users fourth move entered here

		updateBoard(userMove)

		drawBoard()
	}

	////////////////////////////////////////////////////////////////////////////////////////////////

	//computer fourth move here

	if computerWin == false {

		checkImmediateWin()

		if fourthMove == false {
			checkImmediateDanger()
		}

		if fourthMove == false {
			computerNextMove()
		}

		time.Sleep(1 * time.Second)

		drawBoard()
	}

	////////////////////////////////////////////////////////////////////////////////////////////////

	if computerWin == false {

		fmt.Scanln(&userMove) //users last move entered here

		updateBoard(userMove)

		drawBoard()
	}

	////////////////////////////////////////////////////////////////////////////////////////////////

	if computerWin == false {
		fmt.Print("Tie game. Want to play again?")
	} else {
		fmt.Print("I win! Play again?")
	}
	time.Sleep(2 * time.Second)

}

func firstDrawBoard() {
	fmt.Print("\n\n\n\n\n\n\n\n\n\n\n") //clearing terminal
	board[0] = "  1   2   3"
	board[1] = " "
	board[2] = "_" //square1
	board[3] = "|"
	board[4] = "_" //square2
	board[5] = "|"
	board[6] = "_" //square3
	board[7] = "4"
	board[8] = "_" //square4
	board[9] = "|"
	board[10] = "_" //square5
	board[11] = "|"
	board[12] = "_" //square6
	board[13] = "6"
	board[14] = " "
	board[15] = " " //square7
	board[16] = "|"
	board[17] = " " //square8
	board[18] = "|"
	board[19] = " " //square9
	board[20] = "  7   8   9"
	fmt.Println("\n" + board[0])
	fmt.Println(board[1], board[2], board[3], board[4], board[5], board[6])
	fmt.Println(board[7], board[8], board[9], board[10], board[11], board[12], board[13])
	fmt.Println(board[14], board[15], board[16], board[17], board[18], board[19])
	fmt.Println(board[20] + "\n")
}

func drawBoard() {
	fmt.Print("\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Println(board[0])
	fmt.Println(board[1], board[2], board[3], board[4], board[5], board[6])
	fmt.Println(board[7], board[8], board[9], board[10], board[11], board[12], board[13])
	fmt.Println(board[14], board[15], board[16], board[17], board[18], board[19])
	fmt.Println(board[20] + "\n\n")

}

func updateBoard(int) {
	switch userMove {
	case 1:
		board[2] = "X"
	case 2:
		board[4] = "X"
	case 3:
		board[6] = "X"
	case 4:
		board[8] = "X"
	case 5:
		board[10] = "X"
	case 6:
		board[12] = "X"
	case 7:
		board[15] = "X"
	case 8:
		board[17] = "X"
	case 9:
		board[19] = "X"
	}
}

//computer logic here
func computerFirstMove() {
	if board[10] == "X" {
		board[2] = "O"
	} else {
		board[10] = "O"
	}
}

func checkImmediateDanger() {

	//this is to check if the human is about to win and blocking it (8 3-in-a-row possibilities x 3 = 24 possibilities)
	secondMove = true
	thirdMove = true
	fourthMove = true
	if board[2] == "X" && board[4] == "X" && board[6] == "_" { //horizontal 2-in-a-row checks
		board[6] = "O"
	} else if board[2] == "X" && board[6] == "X" && board[4] == "_" {
		board[4] = "O"
	} else if board[4] == "X" && board[6] == "X" && board[2] == "_" {
		board[2] = "O"
	} else if board[8] == "X" && board[10] == "X" && board[12] == "_" {
		board[12] = "O"
	} else if board[8] == "X" && board[12] == "X" && board[10] == "_" {
		board[10] = "O"
	} else if board[10] == "X" && board[12] == "X" && board[8] == "_" {
		board[8] = "O"
	} else if board[15] == "X" && board[17] == "X" && board[19] == " " {
		board[19] = "O"
	} else if board[15] == "X" && board[19] == "X" && board[17] == " " {
		board[17] = "O"
	} else if board[17] == "X" && board[19] == "X" && board[15] == " " {
		board[15] = "O"
	} else if board[2] == "X" && board[8] == "X" && board[15] == " " { //vertical checks
		board[15] = "O"
	} else if board[2] == "X" && board[15] == "X" && board[8] == "_" {
		board[8] = "O"
	} else if board[8] == "X" && board[15] == "X" && board[2] == "_" {
		board[2] = "O"
	} else if board[4] == "X" && board[10] == "X" && board[17] == " " {
		board[17] = "O"
	} else if board[4] == "X" && board[17] == "X" && board[10] == "_" {
		board[10] = "O"
	} else if board[10] == "X" && board[17] == "X" && board[4] == "_" {
		board[4] = "O"
	} else if board[6] == "X" && board[12] == "X" && board[19] == " " {
		board[19] = "O"
	} else if board[6] == "X" && board[19] == "X" && board[12] == "_" {
		board[12] = "O"
	} else if board[12] == "X" && board[19] == "X" && board[6] == "_" {
		board[6] = "O"
	} else if board[2] == "X" && board[10] == "X" && board[19] == " " { //diagonal checks
		board[19] = "O"
	} else if board[2] == "X" && board[19] == "X" && board[10] == "_" {
		board[10] = "O"
	} else if board[10] == "X" && board[19] == "X" && board[2] == "_" {
		board[2] = "O"
	} else if board[6] == "X" && board[10] == "X" && board[15] == " " {
		board[15] = "O"
	} else if board[6] == "X" && board[15] == "X" && board[10] == "_" {
		board[10] = "O"
	} else if board[10] == "X" && board[15] == "X" && board[6] == "_" {
		board[6] = "O"
	} else {
		secondMove = false
		thirdMove = false
		fourthMove = false
	}
}

func checkImmediateWin() {

	//this is to check if the computeris about to win
	computerWin = true
	thirdMove = true
	fourthMove = true

	if board[2] == "O" && board[4] == "O" && board[6] == "_" { //horizontal 2-in-a-row checks
		board[6] = "O"
	} else if board[2] == "O" && board[6] == "O" && board[4] == "_" {
		board[4] = "O"
	} else if board[4] == "O" && board[6] == "O" && board[2] == "_" {
		board[2] = "O"
	} else if board[8] == "O" && board[10] == "O" && board[12] == "_" {
		board[12] = "O"
	} else if board[8] == "O" && board[12] == "O" && board[10] == "_" {
		board[10] = "O"
	} else if board[10] == "O" && board[12] == "O" && board[8] == "_" {
		board[8] = "O"
	} else if board[15] == "O" && board[17] == "O" && board[19] == " " {
		board[19] = "O"
	} else if board[15] == "O" && board[19] == "O" && board[17] == " " {
		board[17] = "O"
	} else if board[17] == "O" && board[19] == "O" && board[15] == " " {
		board[15] = "O"
	} else if board[2] == "O" && board[8] == "O" && board[15] == " " { //vertical checks
		board[15] = "O"
	} else if board[2] == "O" && board[15] == "O" && board[8] == "_" {
		board[8] = "O"
	} else if board[8] == "O" && board[15] == "O" && board[2] == "_" {
		board[2] = "O"
	} else if board[4] == "O" && board[10] == "O" && board[17] == " " {
		board[17] = "O"
	} else if board[4] == "O" && board[17] == "O" && board[10] == "_" {
		board[10] = "O"
	} else if board[10] == "O" && board[17] == "O" && board[4] == "_" {
		board[4] = "O"
	} else if board[6] == "O" && board[12] == "O" && board[19] == " " {
		board[19] = "O"
	} else if board[6] == "O" && board[19] == "O" && board[12] == "_" {
		board[12] = "O"
	} else if board[12] == "O" && board[19] == "O" && board[6] == "_" {
		board[6] = "O"
	} else if board[2] == "O" && board[10] == "O" && board[19] == " " { //diagonal checks
		board[19] = "O"
	} else if board[2] == "O" && board[19] == "O" && board[10] == "_" {
		board[10] = "O"
	} else if board[10] == "O" && board[19] == "O" && board[2] == "_" {
		board[2] = "O"
	} else if board[6] == "O" && board[10] == "O" && board[15] == " " {
		board[15] = "O"
	} else if board[6] == "O" && board[15] == "O" && board[10] == "_" {
		board[10] = "O"
	} else if board[10] == "O" && board[15] == "O" && board[6] == "_" {
		board[6] = "O"
	} else {
		computerWin = false
		thirdMove = false
		fourthMove = false
	}

}

func computerSecondMove() {
	if board[17] == "X" && board[12] == "X" {
		board[19] = "O"
	} else if board[15] == "X" && board[6] == "X" {
		board[4] = "O"
	} else if board[12] == "X" && board[15] == "X" {
		board[19] = "O"
	} else if board[2] == "X" && board[19] == "X" {
		board[4] = "O"
	} else if board[2] == "X" && board[17] == "X" {
		board[15] = "O"
	} else if board[2] == "_" {
		board[2] = "O"
	} else if board[6] == "_" {
		board[6] = "O"
	} else if board[15] == " " {
		board[15] = "O"
	}
}

//this is going to look for open spots on the corner and move there then the rest of the open spots
func computerNextMove() {
	if board[6] == "_" {
		board[6] = "O"
	} else if board[15] == " " {
		board[15] = "O"
	} else if board[19] == " " {
		board[19] = "O"
	} else if board[4] == "_" {
		board[4] = "O"
	} else if board[8] == "_" {
		board[8] = "O"
	} else if board[12] == "_" {
		board[12] = "O"
	} else if board[17] == " " {
		board[17] = "O"
	}
}
