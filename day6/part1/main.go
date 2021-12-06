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

    // list of fishes
    fishes := []int{}

    // Create Scanner to read in file line by line
    scanner := bufio.NewScanner(buffer)
    for scanner.Scan() {
        s := strings.Split(scanner.Text(), ",")
        for _, fish := range(s) {
            fishInt, err := strconv.Atoi(fish)
            if err != nil {
                log.Fatal(err)
            }
            fishes = append(fishes, fishInt)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Initial state: ")
    for _, fish := range(fishes) {
        fmt.Printf("%d,", fish)
    }
    fmt.Println()

    // decrement fish based on days
    for day := 0; day < 80; day++ {
        limit := len(fishes)
        for i := 0; i < limit; i++ {
            if fishes[i] == 0 {
                fishes[i] = 6
                fishes = append(fishes, 8) // new fish with time of 8
            } else {
                fishes[i] -= 1
            }
        }
        fmt.Printf("After\t %d day(s): ", day)
        for i := 0; i < len(fishes); i++ {
            fmt.Printf("%d,", fishes[i])
        }
        fmt.Println()
    }

    fmt.Printf("Total fish: %d\n", len(fishes))

}
