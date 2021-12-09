package main

import(
    "bufio"
    "fmt"
    "log"
    "os"
    "sort"
    "strconv"
    "strings"
)

// findBasin is a recursive function that checks for 9 in the surround values around
// a point in the grid. It returns once it hits 9.
func findBasin(grid [][]int, i int, z int, traversed map[string]bool) int {
    // Base case
    if grid[i][z] == 9 {
        return 0
    } else {
        // ensures we aren't double counting
        if traversed[strconv.Itoa(i) + " " + strconv.Itoa(z)] {
            return 0
        }
    }

    // fmt.Printf("(%d, %d)\n", i, z)
    traversed[strconv.Itoa(i) + " " + strconv.Itoa(z)] = true // add to map to make sure we don't double cross/get in infinite loop

    total := 1

    // top left
    if i == 0 && z == 0 {
        total += findBasin(grid, 0, z + 1, traversed)
        total += findBasin(grid, i + 1, z, traversed)
    } else if i == 0 && z == len(grid[i]) - 1 {
        // top right
        total += findBasin(grid, i, z - 1, traversed)
        total += findBasin(grid, i + 1, z, traversed)
    } else if i == len(grid) - 1 && z == 0 {
        // bottom left
        total += findBasin(grid, len(grid) - 2, 0, traversed)
        total += findBasin(grid, i, 1, traversed)
    } else if i == len(grid) - 1 && z == len(grid[i]) - 1 {
        // bottom right
        total += findBasin(grid, len(grid) - 2, len(grid[i]) - 1, traversed)
        total += findBasin(grid, len(grid) - 1, len(grid[i]) - 2, traversed)
    } else if i == 0 {
        // top edge
        total += findBasin(grid, i, z + 1, traversed)
        total += findBasin(grid, i, z - 1, traversed)
        total += findBasin(grid, i + 1, z, traversed)
    } else if z == 0 {
        // left edge
        total += findBasin(grid, i + 1, z, traversed)
        total += findBasin(grid, i - 1, z, traversed)
        total += findBasin(grid, i, z + 1, traversed)
    } else if i == len(grid) - 1 {
        // bottom edge
        total += findBasin(grid, i, z + 1, traversed)
        total += findBasin(grid, i, z - 1, traversed)
        total += findBasin(grid, i - 1, z, traversed)
    } else if z == len(grid[i]) - 1 {
        // right edge
        total += findBasin(grid, i + 1, z, traversed)
        total += findBasin(grid, i - 1, z, traversed)
        total += findBasin(grid, i, z - 1, traversed)
    } else {
        // up + down + right + left
        total += findBasin(grid, i, z + 1, traversed)
        total += findBasin(grid, i, z - 1, traversed)
        total += findBasin(grid, i + 1, z, traversed)
        total += findBasin(grid, i - 1, z, traversed)
    }

    return total
}

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

    var lowest []string // holds lowest points in grid

    for i, row := range(grid) {
        for z, val := range(row) {
            // top left
            if i == 0 && z == 0 {
                if val < grid[0][1] && val < grid[1][0] {
                    lowest = append(lowest, strconv.Itoa(i) + " " + strconv.Itoa(z))
                }
            } else if i == 0 && z == len(row) - 1 {
                // top right
                if val < grid[1][len(row) - 1] && val < grid[0][len(row) - 2] {
                    lowest = append(lowest, strconv.Itoa(i) + " " + strconv.Itoa(z))
                }
            } else if i == len(grid) - 1 && z == 0 {
                // bottom left
                if val < grid[len(grid) - 2][0] && val < row[z + 1] {
                    lowest = append(lowest, strconv.Itoa(i) + " " + strconv.Itoa(z))
                }
            } else if i == len(grid) - 1 && z == len(row) - 1 {
                // bottom right
                if val < grid[len(grid) - 2][len(row) - 1] && val < grid[len(grid) - 1][len(row) - 2] {
                    lowest = append(lowest, strconv.Itoa(i) + " " + strconv.Itoa(z))
                }
            } else if i == 0 {
                // top edge
                if val < grid[i][z - 1] && val < grid[i][z + 1] && val < grid[1][z] {
                    lowest = append(lowest, strconv.Itoa(i) + " " + strconv.Itoa(z))
                }
            } else if z == 0 {
                // left edge
                if val < grid[i - 1][0] && val < grid[i][1] && val < grid[i + 1][0] {
                    lowest = append(lowest, strconv.Itoa(i) + " " + strconv.Itoa(z))
                }
            } else if i == len(grid) - 1 {
                // bottom edge
                if val < grid[i][z - 1] && val < grid[i][z + 1] && val < grid[i - 1][z] {
                    lowest = append(lowest, strconv.Itoa(i) + " " + strconv.Itoa(z))
                }
            } else if z == len(row) - 1 {
                // right edge
                if val < grid[i + 1][z] && val < grid[i - 1][z] && val < grid[i][z - 1] {
                    lowest = append(lowest, strconv.Itoa(i) + " " + strconv.Itoa(z))
                }
            } else {
                // up + down + right + left
                if val < grid[i][z + 1] && val < grid[i][z - 1] && val < grid[i + 1][z] && val < grid[i - 1][z] {
                    lowest = append(lowest, strconv.Itoa(i) + " " + strconv.Itoa(z))
                }
            }
        }
    }

    var basins []int // holds sizes of all found basins

    for _, pair := range(lowest) {
        i, err := strconv.Atoi(strings.Split(pair, " ")[0])
        if err != nil {
            log.Fatal(err)
        }
        z, err := strconv.Atoi(strings.Split(pair, " ")[1])
        if err != nil {
            log.Fatal(err)
        }
        traversed := make(map[string]bool)
        // fmt.Println("------------------")
        // fmt.Printf("Checking: %s\n", pair)
        basins = append(basins, findBasin(grid, i, z, traversed))
        // fmt.Println("------------------")
    }

    // sort basins from greatest to least
    sort.Sort(sort.Reverse(sort.IntSlice(basins)))

    fmt.Printf("Solution: %d\n", basins[0] * basins[1] * basins[2])
}
