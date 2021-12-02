package main

import(
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

func main() {
    fmt.Printf("Reading in input.txt...\n")

    // Open "input.txt"
    buffer, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer buffer.Close()

    horizontal := 0 // temporary var to hold horizontal position
    depth := 0 // var to hold depth

    // Create Scanner to read in file line by line
    scanner := bufio.NewScanner(buffer)
    for scanner.Scan() {
        command := strings.Fields(scanner.Text()) // split string by whitespace
        t := command[0]  // forward/up/down
        val, err := strconv.Atoi(command[1]) // convert string into int
        if err != nil {
            log.Fatal(err)
        }

        if t == "forward" {
            horizontal += val
        } else if t == "up" {
            depth -= val
        } else if t == "down" {
            depth += val
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Solution: %d\n", horizontal * depth)
}
