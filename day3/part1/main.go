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

    // create list of maps to keep track of most common bits in each place
    list := make([]map[rune]int, 12)
    for i := 0; i < 12; i++ {
        m := make(map[rune]int)
        list[i] = m
    }

    // Create Scanner to read in file line by line
    scanner := bufio.NewScanner(buffer)
    for scanner.Scan() {
        for index, bit := range scanner.Text() {
            list[index][bit] += 1
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    // determine gamma rate
    gamma := ""
    for i := 0; i < 12; i++ {
        if list[i]['1'] >= list[i]['0'] {
            gamma += string('1')
        } else {
            gamma += string('0')
        }
    }
    convertedGamma, err := strconv.ParseInt(gamma, 2, 64)
    if err != nil {
        log.Fatal(err)
    }

    // determine epsilon rate
    epsilon := ""
    for i := 0; i < 12; i++ {
        if gamma[i] == '1' {
            epsilon += "0"
        } else {
            epsilon += "1"
        }
    }
    convertedEpsilon, err := strconv.ParseInt(epsilon, 2, 64)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Power consumption: %d\n", convertedGamma * convertedEpsilon)
}
