package main

import(
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

func main() {
    // Open "input.txt"
    buffer, err := os.Open("input.txt")
    // buffer, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer buffer.Close()

    // Create Scanner to read in file line by line
    scanner := bufio.NewScanner(buffer)
    for scanner.Scan() {
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
