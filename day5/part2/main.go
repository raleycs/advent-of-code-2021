package main

import(
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

func main() {
    fmt.Printf("Reading in input.txt...\n")

    // Open "input.txt"
    buffer, err := os.Open("input.txt")
    // buffer, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer buffer.Close()

    // create map to keep track of danger points
    danger := make(map[string]int)

    // Create Scanner to read in file line by line
    scanner := bufio.NewScanner(buffer)
    for scanner.Scan() {
        endpoints := strings.Split(scanner.Text(), " -> ")
        coordinates := [][]string{}
        for _, pair := range(endpoints) {
            coordinates = append(coordinates, strings.Split(pair, ","))
        }

        // parse coordinates as strings
        x0, err := strconv.Atoi(coordinates[0][0])
        if err != nil {
            log.Fatal(err)
        }
        y0, err := strconv.Atoi(coordinates[0][1])
        if err != nil {
            log.Fatal(err)
        }
        x1, err := strconv.Atoi(coordinates[1][0])
        if err != nil {
            log.Fatal(err)
        }
        y1, err := strconv.Atoi(coordinates[1][1])
        if err != nil {
            log.Fatal(err)
        }

        // verticals
        if coordinates[0][0] == coordinates[1][0] {
            if y1 >= y0 {
                for i := y0; i <= y1; i++ {
                    danger[coordinates[0][0] + "," + strconv.Itoa(i)] += 1
                }
            } else {
                for i := y0; i >= y1; i-- {
                    danger[coordinates[0][0] + "," + strconv.Itoa(i)] += 1
                }
            }
            // horizontals
        } else if coordinates[0][1] == coordinates[1][1] {
            if x1 >= x0 {
                for i := x0; i <= x1; i++ {
                    danger[strconv.Itoa(i) + "," + coordinates[1][1]] += 1
                }
            } else {
                for i := x0; i >= x1; i-- {
                    danger[strconv.Itoa(i) + "," + coordinates[1][1]] += 1
                }
            }
            // diagonals
        } else {
            // fmt.Printf("DEBUG: %d,%d\t", x0, y0)
            // fmt.Printf("DEBUG: %d,%d\n", x1, y1)
            // top right to bottom left
            if x1 >= x0  && y0 >= y1 {
                fmt.Println("here 1")
                for i := x0; i <= x1; i++ {
                    diff := y0 - (i - x0)
                    if diff < 0 {
                        diff = -diff
                    }
                    // fmt.Printf("DEBUG: %d,%d\n", i, diff)
                    danger[strconv.Itoa(i) + "," + strconv.Itoa(diff)] += 1
                }
            } else if x1 >= x0 && y1 >= y0 {
                // bottom left to top right
                fmt.Println("here 2")
                for i := x0; i <= x1; i++ {
                    diff := y0 + (i - x0)
                    if diff < 0 {
                        diff = -diff
                    }
                    // fmt.Printf("DEBUG: %d,%d\n", i, diff)
                    danger[strconv.Itoa(i) + "," + strconv.Itoa(diff)] += 1
                }

            } else if x1 < x0 && y1 < y0 {
                // top left to bottom right
                fmt.Println("here 3")
                for i := x0; i >= x1; i-- {
                    diff := y0 - (x0 - i)
                    if diff < 0 {
                        diff = -diff
                    }
                    // fmt.Printf("DEBUG: %d,%d\n", i, diff)
                    danger[strconv.Itoa(i) + "," + strconv.Itoa(diff)] += 1
                }
            } else if x1 < x0 && y1 >= y0 {
                // bottom right to top left
                fmt.Println("here 4")
                for i := x0; i >= x1; i-- {
                    diff := y0 + (x0 - i)
                    if diff < 0 {
                        diff = -diff
                    }
                    // fmt.Printf("DEBUG: %d,%d\n", i, diff)
                    danger[strconv.Itoa(i) + "," + strconv.Itoa(diff)] += 1
                }
            }
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    overlaps := 0
    for _, v := range(danger) {
        if v >= 2 {
            overlaps += 1
        }
    }

    // for x := 0; x < 1000; x++ {
    //     for y := 0; y < 1000; y++ {
    //         if danger[fmt.Sprintf("%d,%d", y, x)] == 0 {
    //             fmt.Printf(".")
    //         } else {
    //             fmt.Printf("%d", danger[fmt.Sprintf("%d,%d", y, x)])
    //         }
    //     }
    //     fmt.Println()
    // }

    fmt.Printf("Total overlaps: %d\n", overlaps)
}
