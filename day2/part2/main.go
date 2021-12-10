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
    aim := 0 // var to hold aim

    // Create Scanner to read in file line by line
    scanner := bufio.NewScanner(buffer)
    for scanner.Scan() {
        command := strings.Fields(scanner.Text()) // split string by whitespace
        t := command[0]  // forward/up/down
        val, err := strconv.Atoi(command[1]) // convert string into int
        if err != nil {
            log.Fatal(err)
        }

        // change values based on input commands
        if t == "forward" {
            horizontal += val
            depth += aim * val
        } else if t == "up" {
            aim -= val
        } else if t == "down" {
            aim += val
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Solution: %d\n", horizontal * depth)
}
