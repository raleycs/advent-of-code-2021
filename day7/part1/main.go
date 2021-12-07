package main

import(
    "bufio"
    "fmt"
    "log"
    "math"
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

    var crabs []int // crab submarine
    hash := make(map[int]int) // map to determine lowest total costs for each row

    // Create Scanner to read in file line by line
    scanner := bufio.NewScanner(buffer)
    for scanner.Scan() {
        str := strings.Split(scanner.Text(), ",")
        for _, s := range(str) {
            num, err := strconv.Atoi(s)
            if err != nil {
                log.Fatal(err)
            }
            crabs = append(crabs, num)
        }
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    // iterate through each row that has a crab
    // and find the cheapest row to align to
    for _, current := range(crabs) {
        if hash[current] != 0 {
            continue
        }
        for i := 0; i < len(crabs); i++ {
            cost := crabs[i] - current
            if cost < 0 {
                cost = -cost
            }
            hash[current] += cost
        }
    }

    // search through map to find cheapest row
    min := math.MaxInt
    for _, value := range(hash) {
        if min > value {
            min = value
        }
    }
    fmt.Println(min)
}
