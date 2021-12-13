package main

import(
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

// toInt takes in a string that contains the coordinates of a point in the origami graph.
// The function returns x and y value coordinates as ints
func toInt(pair string) (int, int) {
    x := strings.Split(pair, ",")[0]
    y := strings.Split(pair, ",")[1]

    xInt, err := strconv.Atoi(x)
    if err != nil {
        log.Fatal(err)
    }
    yInt, err := strconv.Atoi(y)
    if err != nil {
        log.Fatal(err)
    }

    return xInt, yInt
}

func main() {
    // Open "input.txt"
    buffer, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer buffer.Close()

    hash := make(map[string]int) // hash that maps coordinates to 1/0 determining if they are marked on the map
    instructions := []string{} // slice storing all folding instructions
    line := 1 // int holding current line number

    // Create Scanner to read in file line by line
    scanner := bufio.NewScanner(buffer)
    for scanner.Scan() {
        // not proud of the way I parsed this :)
        if line <= 839 {
            hash[scanner.Text()] = 1
        } else if line >= 841 {
            instructions = append(instructions, strings.Split(scanner.Text(), " ")[2])
        }
        line += 1
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    for _, inst := range(instructions) {

        visible := 0 // int holding number of visible points after each fold
        direction := strings.Split(inst, "=")[0] // string determining horizontal/vertical fold
        point := strings.Split(inst, "=")[1] // string holding the point where we make the fold
        p, err := strconv.Atoi(point) // int value holding point where we make the fold
        if err != nil {
            log.Fatal(err)
        }

        // vertical fold
        if direction == "x" {
            for pair, exist := range(hash) {
                if exist == 1 {
                    x, y := toInt(pair)
                    diff := x - p

                    // find the difference between the fold and the point
                    if diff < 0 {
                        diff = -diff
                    }

                    // if the original point if to the left of the fold...
                    if x < p {
                        // if there is an overlap, remove the right side overlap so we don't double count
                        if hash[strconv.Itoa(p + diff) + "," + strconv.Itoa(y)] == 1 {
                            hash[strconv.Itoa(p + diff) + "," + strconv.Itoa(y)] = 0
                        }
                    } else if x > p {
                        // if there is an overlap, remove the right side overlap so we don't double count
                        if hash[strconv.Itoa(p - diff) + "," + strconv.Itoa(y)] == 1 {
                            hash[strconv.Itoa(p - diff) + "," + strconv.Itoa(y)] = 0
                        }
                    }

                    visible += 1
                }
            }
            fmt.Printf("%s=%s: %d\n", direction, point, visible)
        } else if direction == "y" {
            // horizontal cut
            for pair, exist := range(hash) {
                if exist == 1 {
                    x, y := toInt(pair)
                    diff := y - p

                    if diff < 0 {
                        diff = -diff
                    }

                    // if the original point is above the fold point
                    if y > p {
                        // if there is an overlap, remove the right side overlap so we don't double count
                        if hash[strconv.Itoa(x) + "," + strconv.Itoa(p - diff)] == 1 {
                            hash[strconv.Itoa(x) + "," + strconv.Itoa(p - diff)] = 0
                        }
                    } else if y < p {
                        // if there is an overlap, remove the right side overlap so we don't double count
                        if hash[strconv.Itoa(x) + "," + strconv.Itoa(p + diff)] == 1 {
                            hash[strconv.Itoa(x) + "," + strconv.Itoa(p + diff)] = 0
                        }
                    }
                    visible += 1
                }
            }
            fmt.Printf("%s=%s: %d\n", direction, point, visible)
        }
    }
}
