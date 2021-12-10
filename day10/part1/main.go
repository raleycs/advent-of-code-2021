package main

import(
    "bufio"
    "fmt"
    "log"
    "os"
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

    hash := make(map[string]int)

    // Create Scanner to read in file line by line
    scanner := bufio.NewScanner(buffer)
    for scanner.Scan() {
        chunk := []string{} // slice holding opening characters -- operates like a stack
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
                // if so add closing character total to hash
                // else pop the stack
                if char == "}" {
                    if chunk[len(chunk) - 1] != "{" {
                        hash[char] += 1
                        break
                    } else {
                        chunk = pop(chunk)
                    }
                } else if char == "]" {
                    if chunk[len(chunk) - 1] != "[" {
                        hash[char] += 1
                        break
                    } else {
                        chunk = pop(chunk)
                    }
                } else if char == ")" {
                    if chunk[len(chunk) - 1] != "(" {
                        hash[char] += 1
                        break
                    } else {
                        chunk = pop(chunk)
                    }
                } else if char == ">" {
                    if chunk[len(chunk) - 1] != "<" {
                        hash[char] += 1
                        break
                    } else {
                        chunk = pop(chunk)
                    }
                }
            }
        }
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    score := 0 // int to hold score

    // calculate score
    for k, v := range(hash) {
        if k == ")" {
            score += 3 * v
        }
        if k == "]" {
            score += 57 * v
        }
        if k == "}" {
            score += 1197 * v
        }
        if k == ">" {
            score += 25137 * v
        }
    }

    fmt.Printf("Score: %d\n", score)
}
