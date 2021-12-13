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
    // buffer, err := os.Open("test.txt")
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

                    if diff < 0 {
                        diff = -diff
                    }

                    // if the original point is on the right, remove it and mark it on the left
                    if x > p {
                        hash[strconv.Itoa(p - diff) + "," + strconv.Itoa(y)] = 1
                        hash[pair] = 0
                    }
                }
            }
        } else if direction == "y" {
            // horizontal fold
            for pair, exist := range(hash) {
                if exist == 1 {
                    x, y := toInt(pair)
                    diff := y - p

                    if diff < 0 {
                        diff = -diff
                    }

                    // if the original point is above the fold point, remove it and add it to the top
                    if y > p {
                        hash[strconv.Itoa(x) + "," + strconv.Itoa(p - diff)] = 1
                        hash[pair] = 0
                    }
                }
            }
        }
    }

    // find the max lengths of what we need to print out
    xMax := -1
    yMax := -1
    for pair, marked := range(hash) {
        if marked == 1 {
            x, y := toInt(pair)
            if x > xMax {
                xMax = x
            }
            if y > yMax {
                yMax = y
            }
        }
    }

    // print out the secret message
    for i := 0; i < yMax + 2; i++ {
        for j := 0; j < xMax + 2; j++ {
            if hash[strconv.Itoa(j) + "," + strconv.Itoa(i)] == 1 {
                fmt.Printf("#")
            } else {
                fmt.Printf(".")
            }
        }
        fmt.Println()
    }
}
