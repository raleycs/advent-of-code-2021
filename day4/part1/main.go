package main

import(
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

// readInput reads input.txt and
// returns all binaries in a slice
func readInput()([]string, [][][]string) {
    // Open "input.txt"
    // buffer, err := os.Open("input.txt")
    buffer, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer buffer.Close()


    boards := [][][]string{} // slice of 2D slices to hold bingo boards
    board := [][]string{} // new bingo board to add
    drawn := []string{} // slice of numbers drawn

    // Create Scanner to read file and fill bingo boards
    readBoard := false // determine when we should start reading bingo boards
    scanner := bufio.NewScanner(buffer)
    for scanner.Scan() {
        if readBoard {
            if len(board) != 5 && scanner.Text() != "" {
                board = append(board, strings.Split(strings.Join(strings.Fields(scanner.Text()), " "), " "))
            } else if len(board) == 5 && scanner.Text() == "" {
                boards = append(boards, board)
                board = nil // clear bingo board
            }
        } else {
            // read first line as drawn numbers
            drawn = strings.Split(strings.Join(strings.Fields(scanner.Text()), " "), ",")
            readBoard = true
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return drawn, boards
}

// getScore counts total score
func getScore(board [][]string, winningNumber string) int {
    score := 0
    for x := 0; x < 5; x++ {
        for y := 0; y < 5; y++ {
            if board[x][y] != "x" {
                points, err := strconv.Atoi(board[x][y])
                if err != nil {
                    log.Fatal(err)
                }
                score += points
            }
        }
    }
    winner, err := strconv.Atoi(winningNumber)
    if err != nil {
        log.Fatal(err)
    }
    return winner * score
}

// get board state
func viewBoard(board [][]string) {
    for x := 0; x < 5; x++ {
        for y := 0; y < 5; y++ {
            fmt.Printf("%s \t", board[x][y])
        }
        fmt.Println()
    }
}

// check if we have a horizontal bingo
// return true if we have a bingo
func checkHorizontal(board [][]string) bool {
    for y := 0; y < 5; y++ {
        if board[y][0] == "x" {
            for x := 0; x < 5; x++ {
                if board[y][x] != "x" {
                    return false
                }
            }
            return true
        }
    }
    return false
}

// check if we have a vertical bingo
// return true if we have a bingo
func checkVertical(board [][]string) bool {
    for x := 0; x < 5; x++ {
        if board[0][x] == "x" {
            for y := 0; y < 5; y++ {
                if board[y][x] != "x" {
                    return false
                }
            }
            return true
        }
    }
    return false
}

// check if we have a diagonal bingo
// return true if we have a bingo
func checkDiagonal(board [][]string) bool {
    // top left to bottom right
    if board[0][0] == "x" && board[1][1] == "x" && board[2][2] == "x" && board[3][3] == "x" && board[4][4] == "x" {
        return true
    }
    // bottom left to top right
    if board[4][0] == "x" && board[3][1] == "x" && board[2][2] == "x" && board[1][3] == "x" && board[0][4] == "x" {
        return true
    }
    return false
}

// checkWinner check if a bingo board has won
// and if so calculate the winning score and return it
func checkWinner(board [][]string, lastNumber string) int {
    score := -1

    for x := 0; x < 5; x++ {
        for y := 0; y < 5; y++ {
            // we have the number drawn
            if board[x][y] == lastNumber {
                board[x][y] = "x"
                // check for horizontal bingo
                if checkHorizontal(board) {
                    return getScore(board, lastNumber)
                }
                // check for vertical bingo
                if checkVertical(board) {
                    return getScore(board, lastNumber)
                }
                // check for diagonal bingo
                // if checkDiagonal(board) {
                //     return getScore(board, lastNumber)
                // }
                return score
            }
        }
    }

    return score
}

func main() {
    fmt.Printf("Reading in input.txt...\n")

    drawn, boards := readInput()

    // clean up white spaces in boards + drawn list
    for _, board := range(boards) {
        for x := 0; x < 5; x++ {
            for y := 0; y < 5; y++ {
                board[x][y] = strings.ReplaceAll(board[x][y], " ", "")
            }
        }

    }
    for i, num := range(drawn) {
        drawn[i] = strings.ReplaceAll(num, " ", "")
    }

    // parse through drawn numbers and
    // find the bingo board that wins it all
    for _, num := range(drawn) {
        fmt.Printf("Number drawn: %s\n", num)
        for i, board := range(boards) {
            score := checkWinner(board, num)
            if score != -1 {
                fmt.Printf("Winner: Board %d\n", i + 1)
                fmt.Printf("Score: %d\n", score)
                viewBoard(board)
                os.Exit(0)
            }
        }
    }
}
