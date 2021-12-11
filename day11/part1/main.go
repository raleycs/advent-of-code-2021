package main

import(
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

// printGrid prints out all elements
// in the grid variable
func printGrid(grid [][]int) {
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            fmt.Printf("%d", grid[i][j])
        }
        fmt.Println()
    }
}


// flash is a recursive function that
// increments elements around a given one.
// base case is when the element is less than 10.
// returns the total number of flashes during the call of the function.
func flash(grid [][]int, i int, z int) int {
    flashes := 0
    if grid[i][z] == 0 {
        return 0
    } else if grid[i][z] == -1 {
        grid[i][z] = 1
        return 0
    }
    grid[i][z] += 1

    // Base case
    if grid[i][z] <= 9 {
        return 0
    }

    grid[i][z] = 0

    // top left
    if i == 0 && z == 0 {
        flashes += flash(grid, 0, z + 1)
        flashes += flash(grid, i + 1, z)
        flashes += flash(grid, i + 1, z + 1)
    } else if i == 0 && z == len(grid[i]) - 1 {
        // top right
        flashes += flash(grid, i, z - 1)
        flashes += flash(grid, i + 1, z)
        flashes += flash(grid, i + 1, z - 1)
    } else if i == len(grid) - 1 && z == 0 {
        // bottom left
        flashes += flash(grid, len(grid) - 2, 0)
        flashes += flash(grid, i, 1)
        flashes += flash(grid, len(grid) - 2, 1)
    } else if i == len(grid) - 1 && z == len(grid[i]) - 1 {
        // bottom right
        flashes += flash(grid, len(grid) - 2, len(grid[i]) - 1)
        flashes += flash(grid, len(grid) - 1, len(grid[i]) - 2)
        flashes += flash(grid, len(grid) - 2, len(grid[i]) - 2)
    } else if i == 0 {
        // top edge
        flashes += flash(grid, i, z + 1)
        flashes += flash(grid, i, z - 1)
        flashes += flash(grid, i + 1, z)
        flashes += flash(grid, i + 1, z - 1)
        flashes += flash(grid, i + 1, z + 1)
    } else if z == 0 {
        // left edge
        flashes += flash(grid, i + 1, z)
        flashes += flash(grid, i - 1, z)
        flashes += flash(grid, i, z + 1)
        flashes += flash(grid, i + 1, z + 1)
        flashes += flash(grid, i - 1, z + 1)
    } else if i == len(grid) - 1 {
        // bottom edge
        flashes += flash(grid, i, z + 1)
        flashes += flash(grid, i, z - 1)
        flashes += flash(grid, i - 1, z)
        flashes += flash(grid, i - 1, z - 1)
        flashes += flash(grid, i - 1, z + 1)
    } else if z == len(grid[i]) - 1 {
        // right edge
        flashes += flash(grid, i + 1, z)
        flashes += flash(grid, i - 1, z)
        flashes += flash(grid, i, z - 1)
        flashes += flash(grid, i + 1, z - 1)
        flashes += flash(grid, i - 1, z - 1)
    } else {
        // up + down + right + left
        flashes += flash(grid, i, z + 1)
        flashes += flash(grid, i, z - 1)
        flashes += flash(grid, i + 1, z - 1)
        flashes += flash(grid, i + 1, z + 1)
        flashes += flash(grid, i + 1, z)
        flashes += flash(grid, i - 1, z)
        flashes += flash(grid, i - 1, z - 1)
        flashes += flash(grid, i - 1, z + 1)
    }

    grid[i][z] = 0
    return flashes + 1
}

func main() {
    // Open "input.txt"
    buffer, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer buffer.Close()

    grid := [][]int{} // 2-D slice to act as grid of octopi

    // Create Scanner to read in file line by line
    scanner := bufio.NewScanner(buffer)
    for scanner.Scan() {
        // fill grid by converting string input to int rows
        row := []int{}
        for _, char := range(scanner.Text()) {
            num, err := strconv.Atoi(string(char))
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

    flashes := 0
    for n := 0; n < 100; n++ {
        fmt.Printf("Step %d\n", n)
        printGrid(grid)
        fmt.Println()
        for i := 0; i < len(grid); i++ {
            for j := 0; j < len(grid[i]); j++ {
                if grid[i][j] == 0 {
                    grid[i][j] = -1
                }
            }
        }
        for i := 0; i < len(grid); i++ {
            for j := 0; j < len(grid[i]); j++ {
                flashes += flash(grid, i, j)
            }
        }
        fmt.Printf("Flashes: %d\n", flashes)
        fmt.Println("-------")
    }
}
