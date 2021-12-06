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

    hash := make(map[int]int)
    for _, fish := range(fishes) {
        hash[fish] += 1
    }

    // fmt.Printf("Initial state: ")
    // for k, v := range(hash) {
    //     fmt.Printf("Fish with timer %d: %d,", k, v)
    // }
    // fmt.Println()

    total := len(fishes)

    // decrement fish based on days
    // NOTE: The best best way to improve
    // the speed of the solution is by using
    // hashing tables due to O(1) lookup time
    for day := 0; day < 256; day++ {
        // fmt.Printf("Day %d\n", day)
        // for k, v := range(hash) {
        //     fmt.Printf("Fish with timer %d: %d\n", k, v)
        // }
        // fmt.Println(total)
        // fmt.Println()
        if hash[0] > 0 {
            tmp := hash[8]
            hash[8] = hash[0]
            total += hash[8]
            for i := 1; i < 8; i++ {
                hash[i - 1] = hash[i]
            }
            hash[7] = tmp
            hash[6] += hash[8]
        } else {
            tmp := hash[8]
            for i := 1; i < 8; i++ {
                hash[i - 1] = hash[i]
            }
            hash[7] = tmp
            hash[8] = 0
        }
    }

    fmt.Printf("Total fish: %d\n", total)

}
