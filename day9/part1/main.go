package main

import(
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    // "strings"
)

func main() {
    // Open "input.txt"
    buffer, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer buffer.Close()

    var grid [][]int

    // Create Scanner to read in file line by line
    scanner := bufio.NewScanner(buffer)
    for scanner.Scan() {
        var row []int
        for _, c := range(scanner.Text()) {
            num, err := strconv.Atoi(string(c))
            if err != nil {
                log.Fatal(err)
            }
            row = append(row, num)
        }
        grid = append(grid, row)
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    var lowest []int

    for i, row := range(grid) {
        for z, val := range(row) {
            if i == 0 && z == 0 {
                if val < grid[0][1] && val < grid[1][0] {
                    lowest = append(lowest, val)
                }
            } else if i == 0 && z == len(row) - 1 {
                if val < grid[1][len(row) - 1] && val < grid[0][len(row) - 2] {
                    lowest = append(lowest, val)
                }
            } else if i == len(grid) - 1 && z == 0 {
                if val < grid[len(grid) - 2][0] && val < row[z + 1] {
                    lowest = append(lowest, val)
                }
            } else if i == len(grid) - 1 && z == len(row) - 1 {
                if val < grid[len(grid) - 2][len(row) - 1] && val < grid[len(grid) - 1][len(row) - 2] {
                    lowest = append(lowest, val)
                }
            } else if i == 0 {
                if val < grid[i][z - 1] && val < grid[i][z + 1] && val < grid[1][z] {
                    lowest = append(lowest, val)
                }
            } else if z == 0 {
                if val < grid[i - 1][0] && val < grid[i][1] && val < grid[i + 1][0] {
                    lowest = append(lowest, val)
                }
            } else if i == len(grid) - 1 {
                if val < grid[i][z - 1] && val < grid[i][z + 1] && val < grid[i - 1][z] {
                    lowest = append(lowest, val)
                }
            } else if z == len(row) - 1 {
                if val < grid[i + 1][z] && val < grid[i - 1][z] && val < grid[i][z - 1] {
                    lowest = append(lowest, val)
                }
            } else {
                if val < grid[i][z + 1] && val < grid[i][z - 1] && val < grid[i + 1][z] && val < grid[i - 1][z] {
                    lowest = append(lowest, val)
                }
            }
        }
    }
    fmt.Println(lowest)

    risk := 0
    for _, n := range(lowest) {
        risk += 1 + n
    }
    fmt.Println(risk)
}
