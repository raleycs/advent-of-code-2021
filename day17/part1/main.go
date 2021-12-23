package main

import(
    "fmt"
    "log"
    "strconv"
    "strings"
)

func toString(x int, y int) string {
    xStr := strconv.Itoa(x)
    yStr := strconv.Itoa(y)

    return xStr + "," + yStr
}

func toInt(coordinates string) (int, int) {
    split := strings.Split(coordinates, ",")
    x, err := strconv.Atoi(split[0])
    if err != nil {
        log.Fatal(err)
    }
    var y int
    if split[1][0] == '-' {
        split[1] = string(split[1][1:])
        y, err = strconv.Atoi(split[1])
        if err != nil {
            log.Fatal(err)
        }
        y = -y
    } else {
        y, err = strconv.Atoi(split[1])
        if err != nil {
            log.Fatal(err)
        }
    }

    return x, y
}

func withinBound(coordinates string, targetArea map[string]int) bool {
    x, y := toInt(coordinates)
    if x > targetArea["xMax"] {
        return false
    } else if y < targetArea["yMax"] {
        return false
    } else {
        return true
    }
}

func main() {
    targetArea := make(map[string]int)

    xFloor := 253
    xCeil := 280
    yFloor := -73
    yCeil := -64

    // xFloor := 20
    // xCeil := 30
    // yFloor := -10
    // yCeil := -5

    targetArea["xMax"] = xCeil
    targetArea["yMax"] = yFloor

    for y := yFloor; y <= yCeil; y++ {
        for x := xFloor; x <= xCeil; x++ {
            targetArea[toString(x, y)] = 1
        }
    }

    solution := 0
    for i := -1000; i < 1000; i++ {
        for j := 0; j < xCeil; j++ {
            velocity := toString(j, i)
            position := "0,0"
            highest := 0
            skip := true
            fmt.Printf("Trying: %s\n", velocity)
            for withinBound(position, targetArea) {
                _, newY := toInt(position)
                if newY > highest {
                    highest = newY
                }

                if targetArea[position] == 1 {
                    fmt.Printf("Highest y position reached: %d\n", highest)
                    if solution < highest {
                        solution = highest
                    }
                    break
                } else {
                    xV, yV := toInt(velocity)
                    x, y := toInt(position)

                    if skip {
                        skip = false
                        position = toString(x + xV, y + yV)
                        continue
                    }

                    if xV > 0 {
                        xV -= 1
                    }
                    yV -= 1

                    x += xV
                    y += yV

                    velocity = toString(xV, yV)
                    position = toString(x, y)
                }
            }
        }
    }

    fmt.Println(solution)
}
