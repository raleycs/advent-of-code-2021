package main

import(
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

// identity fills out the hash map based on the known unique inputs
func identify(hash map[string]string, s string) {
    if len(s) == 2 {
        hash["1"] = s
    } else if len(s) == 4 {
        hash["4"] = s
    } else if len(s) == 3 {
        hash["7"] = s
    } else if len(s) == 7 {
        hash["8"] = s
    }
}

func main() {
    // Open "input.txt"
    buffer, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer buffer.Close()

    var input []string // slice holding unique signal patterns and output values delimited by |
    var patterns []string // slice holding unique signal paterns
    var sum int // int holding sum of all decoded values
    var outputs []string // slice holding output values
    hash := make(map[string]string) // mapping to hold patterns to corresponding values

    // Create Scanner to read in file line by line
    scanner := bufio.NewScanner(buffer)
    for scanner.Scan() {
        input = strings.Split(scanner.Text(), "|")
        patterns = strings.Fields(input[0])
        outputs = strings.Fields(input[1])

        // find unique values
        for _, p := range(patterns) {
            identify(hash, p)
        }

        // deduce rest of values
        for _, p := range(patterns) {
            diff := []string{} // slice to hold differences between patterns
            // can be 0, 6, 9
            if len(p) == 6 {
                for _, c := range(hash["8"]) {
                    if strings.Contains(p, string(c)) == false {
                        diff = append(diff, string(c))
                    }
                }
                if len(diff) == 1 {
                    isNine := true
                    for _, c := range(hash["4"]) {
                        if strings.Contains(p, string(c)) == false {
                            isNine = false
                            // can be 0 or 6
                            if strings.Contains(p, string(hash["1"][0])) == false || strings.Contains(p, string(hash["1"][1])) == false {
                                hash["6"] = p // 6 has one missing from 1
                            } else {
                                hash["0"] = p // 0 over laps with 1
                            }
                            break
                        }
                    }
                    // can only be 9
                    if isNine {
                        hash["9"] = p
                    }
                } else {
                    fmt.Println("Something went wrong...")
                }
            } else if len(p) == 5 {
                // can be 2, 3, 5
                for _, c := range(hash["4"]) {
                    if strings.Contains(p, string(c)) == false {
                        diff = append(diff, string(c))
                    }
                }
                if len(diff) == 2 {
                    hash["2"] = p // 2 has exactly 2 missing from 4
                } else if len(diff) == 1 {
                    if strings.Contains(p, string(hash["1"][0])) == false || strings.Contains(p, string(hash["1"][1])) == false {
                        hash["5"] = p // 5 has 1 difference
                    } else {
                        hash["3"] = p // 3 has no differences
                    }
                } else {
                    fmt.Println("Something went wrong!")
                }
            }
        }

        // start decoding values
        var num string // string to hold complete number after decoding all digits
        for _, out := range(outputs) {
            // parse through mappings of digits
            for key, val := range(hash) {
                if len(val) == len(out) {
                    decoded := true // boolean to determine if current output is decoded
                    for _, c := range(out) {
                        if strings.Contains(val, string(c)) == false {
                            decoded = false
                            break
                        }
                    }
                    if decoded {
                        num += key
                    }
                }
            }
        }

        // convert number to int and add to sum
        n, err := strconv.Atoi(num)
        if err != nil {
            log.Fatal(err)
        }
        sum += n
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println(hash)
    fmt.Println(sum)
}
