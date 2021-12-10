package main

import(
    "bufio"
    "fmt"
    "log"
    "os"
    "sort"
)

// pop emulates function of popping from
// a stack data structure for slices
func pop(s []string) []string {
    i := len(s) - 1
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func main() {
    // Open "input.txt"
    buffer, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer buffer.Close()

    scores := []int{}

    // Create Scanner to read in file line by line
    scanner := bufio.NewScanner(buffer)
    for scanner.Scan() {
        chunk := []string{} // slice holding opening characters -- operates like a stack
        skip := false // boolean used to skip if we have a corrupted chunk
        for _, c := range(scanner.Text()) {
            char := string(c) // convert rune to string for convenience

            // add opening characters to chunk slice
            if char == "(" {
                chunk = append(chunk, char)
            } else if char == "[" {
                chunk = append(chunk, char)
            } else if char == "{" {
                chunk = append(chunk, char)
            } else if char == "<" {
                chunk = append(chunk, char)
            } else if len(chunk) > 1 {
                // determine if the chunk is corrupted
                // else pop the stack
                if char == "}" {
                    if chunk[len(chunk) - 1] != "{" {
                        skip = true
                        break
                    } else {
                        chunk = pop(chunk)
                    }
                } else if char == "]" {
                    if chunk[len(chunk) - 1] != "[" {
                        skip = true
                        break
                    } else {
                        chunk = pop(chunk)
                    }
                } else if char == ")" {
                    if chunk[len(chunk) - 1] != "(" {
                        skip = true
                        break
                    } else {
                        chunk = pop(chunk)
                    }
                } else if char == ">" {
                    if chunk[len(chunk) - 1] != "<" {
                        skip = true
                        break
                    } else {
                        chunk = pop(chunk)
                    }
                }
            }
        }

        // corrupted stack, no need to count its score
        if skip {
            continue
        }

        score := 0 // int used to hold score for given chunk

        // calculate score in reverse order (order matters here)
        // since we counted characters in the opposite direction
        for i := len(chunk) - 1; i >= 0; i-- {
            if chunk[i] == "(" {
                score += 1
            }
            if chunk[i] == "[" {
                score += 2
            }
            if chunk[i] == "{" {
                score += 3
            }
            if chunk[i] == "<" {
                score += 4
            }

            // prevents double counting
            if i != 0 {
                score *= 5
            }
        }
        scores = append(scores, score)
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    sort.Sort(sort.IntSlice(scores)) // sorts the scores

    fmt.Println(scores[(len(scores)-1)/2]) // print middle score
}
