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
    hash := make(map[int]int)

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

    for _, current := range(crabs) {
        // fmt.Printf("Changing to %d\n", current)
        if hash[current] != 0 {
            continue
        }
        for i := 0; i < len(crabs); i++ {
            cost := crabs[i] - current
            // if current == 2 {
            //     fmt.Printf("From %d to %d: ", current, crabs[i])
            //     fmt.Println(cost)
            // }
            if cost < 0 {
                cost = -cost
            }
            hash[current] += cost
        }
    }

    min := math.MaxInt
    row := -1
    for key, value := range(hash) {
        if min > value {
            min = value
            row = key
        }
    }
    // fmt.Println(hash)
    fmt.Println(row)
    fmt.Println(min)
}
