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
    max := -1

    // Create Scanner to read in file line by line
    scanner := bufio.NewScanner(buffer)
    for scanner.Scan() {
        str := strings.Split(scanner.Text(), ",")
        for _, s := range(str) {
            num, err := strconv.Atoi(s)
            if err != nil {
                log.Fatal(err)
            }
            if max < num {
                max = num
            }
            crabs = append(crabs, num)
        }
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    for row:= 0; row < max ; row++ {
        if hash[row] != 0 {
            continue
        }
        for i := 0; i < len(crabs); i++ {
            diff := crabs[i] - row
            if diff < 0 {
                diff = -diff
            }
            cost := 0
            for y := 1; y <= diff; y++ {
                cost = cost + y
            }
            if cost < 0 {
                cost = -cost
            }
            hash[row] += cost
        }
    }

    min := math.MaxInt
    for _, value := range(hash) {
        if min > value {
            min = value
        }
    }
    fmt.Println(min)
}
