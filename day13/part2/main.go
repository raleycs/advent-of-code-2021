package main
import(
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

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

    hash := make(map[string]int)

    instructions := []string{}

    // Create Scanner to read in file line by line
    scanner := bufio.NewScanner(buffer)
    line := 1
    for scanner.Scan() {
        if line <= 839 {
            hash[scanner.Text()] = 1
        } else if line >= 841 {
            instructions = append(instructions, strings.Split(scanner.Text(), " ")[2])
        }
        // if line <= 18 {
        //     hash[scanner.Text()] = 1
        // } else if line >= 20 {
        //     instructions = append(instructions, strings.Split(scanner.Text(), " ")[2])
        // }
        line += 1
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    for _, inst := range(instructions) {
        direction := strings.Split(inst, "=")[0]
        point := strings.Split(inst, "=")[1]
        p, err := strconv.Atoi(point)
        if err != nil {
            log.Fatal(err)
        }

        if direction == "x" {
            for pair, exist := range(hash) {
                if exist == 1 {
                    x, y := toInt(pair)
                    diff := x - p

                    if diff < 0 {
                        diff = -diff
                    }

                    if x > p {
                        hash[strconv.Itoa(p - diff) + "," + strconv.Itoa(y)] = 1
                        hash[pair] = 0
                    }
                }
            }
        } else if direction == "y" {
            for pair, exist := range(hash) {
                if exist == 1 {
                    x, y := toInt(pair)
                    diff := y - p

                    if diff < 0 {
                        diff = -diff
                    }

                    if y > p {
                        hash[strconv.Itoa(x) + "," + strconv.Itoa(p - diff)] = 1
                        hash[pair] = 0
                    }
                }
            }
        }
    }

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
