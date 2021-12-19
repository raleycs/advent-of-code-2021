package main

import(
    "bufio"
    "fmt"
    "log"
    "os"
    "sort"
    "strings"
)

func main() {
    // Open "input.txt"
    buffer, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer buffer.Close()

    line := 1 // int holding current line number
    input := ""
    hash := make(map[string]string)

    // Create Scanner to read in file line by line
    scanner := bufio.NewScanner(buffer)
    for scanner.Scan() {
        if line == 1 {
            input = scanner.Text()
        } else if line > 2 {
            key := strings.Split(scanner.Text(), " -> ")[0]
            value := strings.Split(scanner.Text(), " -> ")[1]

            hash[key] = value
        }
        line += 1
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    count := make(map[string]int)

    count = polymerize(40, input, hash)

    ordered := []int{}
    for _, val := range(count) {
        ordered = append(ordered, val)
    }

    sort.Ints(ordered)

    fmt.Printf("Solution: %d\n", ordered[len(ordered) - 1] - ordered[0])
}

func polymerize(steps int, polymer string, hash map[string]string) map[string]int {
    totalPairs := make(map[string]int)

    // get all pairs
    for x := 0; x < len(polymer) - 1; x++ {
        pair := string(polymer[x]) + string(polymer[x + 1])
        totalPairs[pair] += 1
    }

    for i := 0; i < steps; i++ {
        perIteration := make(map[string]int)
        for pair, pairTotal := range(totalPairs) {
            if pairTotal != 0 {
                totalPairs[pair] -= pairTotal
                insertionElement := hash[pair]

                leftPair := string(pair[0]) + insertionElement
                rightPair := insertionElement + string(pair[1])

                perIteration[leftPair] += pairTotal
                perIteration[rightPair] += pairTotal
            }
        }

        // update
        for pair, pairTotal := range(perIteration) {
            totalPairs[pair] = pairTotal
        }
    }

    count := make(map[string]int) // holds mapping for letter to total occurrences

    for pair, pairTotal := range(totalPairs) {
        if pairTotal != 0 {
            count[string(pair[0])] += pairTotal
        }
    }
    
    count[string(polymer[len(polymer) - 1])] += 1

    return count
}
