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

        // fill out maps with any danger zones
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

    fmt.Printf("Total overlaps: %d\n", overlaps)
}
