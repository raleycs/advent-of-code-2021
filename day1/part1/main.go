package main

import(
    "bufio"
    "fmt"
    "log"
    "os"
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

    previousDepth := -1 // temporary var to hold previous depth measure
    increases := 0 // var to hold number of increases

    // Create Scanner to read in file line by line
    scanner := bufio.NewScanner(buffer)
    for scanner.Scan() {
        currentDepth, err := strconv.Atoi(scanner.Text())
        if err != nil {
            log.Fatal(err)
        }

        // don't compare first value
        if previousDepth != -1 {
            if currentDepth > previousDepth {
                increases += 1 // increment if current depth > past depth
            }
        }
        previousDepth = currentDepth
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Total number of increases: %d\n", increases)
}
