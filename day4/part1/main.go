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
func readInput()([]int, [][][]int) {
    // Open "input.txt"
    buffer, err := os.Open("input.txt")
    // buffer, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer buffer.Close()


    boards := [][][]int{} // slice of 2D slices to hold bingo boards
    board := [][]int{} // new bingo board to add
    drawn := []int{} // slice of numbers drawn

    // Create Scanner to read file and fill bingo boards
    readBoard := false // determine when we should start reading bingo boards
    scanner := bufio.NewScanner(buffer)
    for scanner.Scan() {
        if readBoard {
            if len(board) != 5 && scanner.Text() != "" {
                i := []int{}
                for _, s := range(strings.Split(strings.Join(strings.Fields(scanner.Text()), " "), " ")) {
                    num, err := strconv.Atoi(s)
                    if err != nil {
                        log.Fatal(err)
                    }
                    i = append(i, num)
                }
                board = append(board, i)
            } else if len(board) == 5 && scanner.Text() == "" {
                boards = append(boards, board)
                board = nil // clear bingo board
            }
        } else {
            // read first line as drawn numbers
            for _, s := range(strings.Split(strings.Join(strings.Fields(scanner.Text()), " "), ",")) {
                num, err := strconv.Atoi(s)
                if err != nil {
                    log.Fatal(err)
                }
                drawn = append(drawn, num)
            }
            readBoard = true
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return drawn, boards
}

// getScore counts total score
func getScore(board [][]int, winningNumber int) int {
    score := 0
    for x := 0; x < 5; x++ {
        for y := 0; y < 5; y++ {
            if board[x][y] != -1 {
                score += board[x][y]
            }
        }
    }
    fmt.Printf("Calculating final score: %d * %d\n", score, winningNumber)
    return score * winningNumber
}

// get board state
func viewBoard(board [][]int) {
    for x := 0; x < 5; x++ {
        for y := 0; y < 5; y++ {
            fmt.Printf("%d \t", board[x][y])
        }
        fmt.Println()
    }
}

// check if we have a horizontal bingo
// return true if we have a bingo
func checkHorizontal(board [][]int) bool {
    for y := 0; y < 5; y++ {
        if board[y][0] == -1 {
            allMarked := true
            for x := 0; x < 5; x++ {
                if board[y][x] != -1 {
                    allMarked = false
                    break
                }
            }
            if allMarked {
                return true
            }
        }
    }
    return false
}

// check if we have a vertical bingo
// return true if we have a bingo
func checkVertical(board [][]int) bool {
    for x := 0; x < 5; x++ {
        if board[0][x] == -1 {
            allMarked := true
            for y := 0; y < 5; y++ {
                fmt.Println(board[y][x])
                if board[y][x] != -1 {
                    allMarked = false
                    break
                }
            }
            if allMarked {
                return true
            }
        }
    }
    return false
}

// check if we have a diagonal bingo
// return true if we have a bingo
func checkDiagonal(board [][]int) bool {
    // top left to bottom right
    if board[0][0] == -1 && board[1][1] == -1 && board[2][2] == -1 && board[3][3] == -1 && board[4][4] == -1 {
        return true
    }
    // bottom left to top right
    if board[4][0] == -1 && board[3][1] == -1 && board[2][2] == -1 && board[1][3] == -1 && board[0][4] == -1 {
        return true
    }
    return false
}

// checkWinner check if a bingo board has won
// and if so calculate the winning score and return it
func checkWinner(board [][]int, lastNumber int) int {
    for x := 0; x < 5; x++ {
        for y := 0; y < 5; y++ {
            // we have the number drawn
            if board[x][y] == lastNumber {
                board[x][y] = -1
                // check for vertical bingo
                if checkVertical(board) {
                    return getScore(board, lastNumber)
                }
                // check for horizontal bingo
                if checkHorizontal(board) {
                    return getScore(board, lastNumber)
                }
                // check for diagonal bingo
                // if checkDiagonal(board) {
                //     return getScore(board, lastNumber)
                // }
            }
        }
    }

    return -1
}

func main() {
    fmt.Printf("Reading in input.txt...\n")

    drawn, boards := readInput()

    // parse through drawn numbers and
    // find the bingo board that wins it all
    for _, num := range(drawn) {
        fmt.Printf("Number drawn: %d\n", num)
        score := -1
        for i, board := range(boards) {
            score = checkWinner(board, num)
            if score != -1 {
                for x, b := range(boards) {
                    fmt.Printf("Board %d\n", x + 1)
                    viewBoard(b)
                    fmt.Println("----------")
                }
                fmt.Printf("Winner: Board %d\n", i + 1)
                fmt.Printf("Score: %d\n", score)
                os.Exit(0)
            }
        }
    }
}
