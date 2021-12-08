package main

import(
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

// isUnique determines if the given output value is a
// unique string
func isUnique(s string) bool {
    if len(s) == 2 || len(s) == 4 || len(s) == 3 || len(s) == 7 {
        return true
    }
    return false
}

func main() {
    fmt.Printf("Reading in input.txt...\n")

    // Open "input.txt"
    buffer, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer buffer.Close()

    var input []string // slice holding unique signal patterns and output values delimited by |
    // var patterns []string // slice holding unique signal paterns
    var outputs []string // slice holding output values
    var unique int // int holding number of unique output values

    // Create Scanner to read in file line by line
    scanner := bufio.NewScanner(buffer)
    for scanner.Scan() {
        input = strings.Split(scanner.Text(), "|")
        // patterns = strings.Fields(input[0], " ")
        outputs = strings.Fields(input[1])

        for _, out := range(outputs) {
            if isUnique(out) {
                unique += 1
            }
        }
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Unique outputs: %d\n", unique)
}
