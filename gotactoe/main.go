package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	gameboard := []string{
        " ", " |", " ", " |", " \n", 
        "--", "+", "--", "+", "--\n", 
        " ", " |", " ", " |", " \n", 
        "--", "+", "--", "+", "--\n", 
        " ", " |", " ", " |", " \n"}
    availableMoves := []int{0, 2, 4, 10, 12, 14, 20, 22, 24}

    fmt.Println(gameboard)

    fmt.Println("Enter an available move: TL, TM, TR, ML, M, MR, BL, BM, BR")

    gameLoop(availableMoves, gameboard)

}

func gameLoop(availableMoves []int, gameboard []string) {
    gameLoop := true
    rand.Seed(time.Now().UnixNano())

    for gameLoop != false {
        var input string
        fmt.Scanf("%s", &input)

        switch input {
        case "TL":
            playerPick := availableMoves[0]
            gameboard = playerMovePlacement(playerPick, gameboard)
        case "TM":
            playerPick := availableMoves[1]
            gameboard = playerMovePlacement(playerPick, gameboard)
        case "TR":
            playerPick := availableMoves[2]
            gameboard = playerMovePlacement(playerPick, gameboard)
        case "ML":
            playerPick := availableMoves[3]
            gameboard = playerMovePlacement(playerPick, gameboard)
        case "M":
            playerPick := availableMoves[4]
            gameboard = playerMovePlacement(playerPick, gameboard)
        case "MR":
            playerPick := availableMoves[5]
            gameboard = playerMovePlacement(playerPick, gameboard)
        case "BL":
            playerPick := availableMoves[6]
            gameboard = playerMovePlacement(playerPick, gameboard)
        case "BM":
            playerPick := availableMoves[7]
            gameboard = playerMovePlacement(playerPick, gameboard)
        case "BR":
            playerPick := availableMoves[8]
            gameboard = playerMovePlacement(playerPick, gameboard)
        }

        computerPick := availableMoves[rand.Intn(8)]
        gameboard = cpuMovePlacement(computerPick, gameboard)

        fmt.Println(gameboard)

        if checkPlayerWinCondition(gameboard) == false {
            gameLoop = false
        } else if checkComputerWinCondition(gameboard) == false {
            gameLoop = false
        } else {
            gameLoop = true
        }
    }
}

func playerMovePlacement(playerPick int, gameboard []string) []string{
    if gameboard[playerPick] == "x" || gameboard[playerPick] == "x\n"{
        fmt.Println("That move is already taken!")
        return gameboard        
    } else {
            if playerPick == 4 || playerPick == 14 || playerPick == 24 {
                gameboard[playerPick] = "x\n"
                return gameboard
        } else {
            gameboard[playerPick] = "x"
            return gameboard
        }
    }
}

func cpuMovePlacement(computerPick int, gameboard []string) []string{
    if gameboard[computerPick] == "c" || gameboard[computerPick] == "c\n"{
        fmt.Println("That move is already taken!")
        return gameboard        
    } else {
            if computerPick == 4 || computerPick == 14 || computerPick == 24 {
                gameboard[computerPick] = "c\n"
                return gameboard
        } else {
            gameboard[computerPick] = "c"
            return gameboard
        }
    }
}

func checkPlayerWinCondition(gameboard []string) bool {
    if gameboard[0] == "x" && gameboard[2] == "x" && gameboard[4] == "x\n"{
        fmt.Println("You win!") 
        return false
    } else if gameboard[10] == "x" && gameboard[12] == "x" && gameboard[14] == "x\n" {
        fmt.Println("You win!") 
        return false 
    } else if gameboard[10] == "x" && gameboard[12] == "x" && gameboard[14] == "x\n" {
        fmt.Println("You win!") 
        return false
    } else if gameboard[20] == "x" && gameboard[22] == "x" && gameboard[24] == "x\n" {
        fmt.Println("You win!") 
        return false
    } else if gameboard[0] == "x" && gameboard[12] == "x" && gameboard[24] == "x\n" {
        fmt.Println("You win!") 
        return false
    } else if gameboard[4] == "x" && gameboard[12] == "x" && gameboard[20] == "x\n" {
        fmt.Println("You win!") 
        return false
    } else if gameboard[0] == "x" && gameboard[10] == "x" && gameboard[20] == "x\n" {
        fmt.Println("You win!") 
        return false
    } else if gameboard[2] == "x" && gameboard[12] == "x" && gameboard[22] == "x\n" {
        fmt.Println("You win!") 
        return false
    } else if gameboard[4] == "x" && gameboard[14] == "x" && gameboard[24] == "x\n" {
        fmt.Println("You win!") 
        return false
    } else {
        return true
    }
}

func checkComputerWinCondition(gameboard []string) bool {
    if gameboard[0] == "c" && gameboard[2] == "c" && gameboard[4] == "c\n"{
        fmt.Println("The computer wins!") 
        return false
    } else if gameboard[0] == "c" && gameboard[2] == "c" && gameboard[4] == "c\n" {
        fmt.Println("The computer wins!") 
        return false 
    } else if gameboard[10] == "c" && gameboard[12] == "c" && gameboard[14] == "c\n" {
        fmt.Println("The computer wins!") 
        return false
    } else if gameboard[20] == "c" && gameboard[22] == "c" && gameboard[24] == "c\n" {
        fmt.Println("The computer wins!") 
        return false
    } else if gameboard[0] == "c" && gameboard[12] == "c" && gameboard[24] == "c\n" {
        fmt.Println("The computer wins!") 
        return false
    } else if gameboard[4] == "c" && gameboard[12] == "c" && gameboard[20] == "c\n" {
        fmt.Println("The computer wins!") 
        return false
    } else if gameboard[0] == "c" && gameboard[10] == "c" && gameboard[20] == "c\n" {
        fmt.Println("The computer wins!") 
        return false
    } else if gameboard[2] == "c" && gameboard[12] == "c" && gameboard[22] == "c\n" {
        fmt.Println("The computer wins!") 
        return false
    } else if gameboard[4] == "c" && gameboard[14] == "c" && gameboard[24] == "c\n" {
        fmt.Println("The computer wins!") 
        return false
    } else {
        return true
    }
}
