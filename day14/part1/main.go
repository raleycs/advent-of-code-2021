package main

import(
    "bufio"
    "fmt"
    "log"
    "os"
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


    // for i := 0; i < len(polymer) - 1; i++ {
    //     pair := string(polymer[i]) + string(polymer[i + 1])
    //     pairs[pair] = i + 1
    // }
    for _, s := range(polymer) {
        fmt.Printf("%s", string(s))
    }
    fmt.Println()

    steps := 10
    for i := 0; i < steps; i++ {
        pairs := make(map[string][]int)
        for x := 0; x < len(polymer) - 1; x++ {
            pair := string(polymer[x]) + string(polymer[x + 1])
            pairs[pair] = append(pairs[pair], x + 1)
        }
        // fmt.Println(pairs)
        for key, indexes := range(pairs) {
            if hash[key] != "" {
                for _, index := range(indexes) {
                    val := []rune(hash[key])
                    polymer = append(polymer[:index + 1], polymer[index:]...)
                    polymer[index] = val[0]

                    fmt.Printf("Pattern %s: Added %s @ %d\n", key, string(val[0]), index)

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
        for _, s := range(polymer) {
            fmt.Printf("%s", string(s))
        }
        fmt.Println()
    }

    count := make(map[string]int)
    for _, s := range(polymer) {
        str := string(s)

        count[str] += 1
    }

    for k, v := range(count) {
        fmt.Printf("%s: %d\n", k, v)
    }

}
