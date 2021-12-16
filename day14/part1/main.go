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

    polymer := []rune(input)

    steps := 10
    for i := 0; i < steps; i++ {
        pairs := make(map[string][]int)
        for x := 0; x < len(polymer) - 1; x++ {
            pair := string(polymer[x]) + string(polymer[x + 1])
            pairs[pair] = append(pairs[pair], x + 1)
        }
        for key, indexes := range(pairs) {
            if hash[key] != "" {
                for _, index := range(indexes) {
                    val := []rune(hash[key])
                    polymer = append(polymer[:index + 1], polymer[index:]...)
                    polymer[index] = val[0]

                    for k, in := range(pairs) {
                        for slot, v := range(in){
                            if v > index {
                                pairs[k][slot] += 1
                            }
                        }
                    }
                }

            }
        }
    }

    count := make(map[string]int)

    for _, s := range(polymer) {
        count[string(s)] += 1
    }

    ordered := []int{}
    for _, val := range(count) {
        ordered = append(ordered, val)
    }

    sort.Ints(ordered)

    fmt.Printf("%d\n", ordered[len(ordered) - 1] - ordered[0])
}
