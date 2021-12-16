package main

import( "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

func toString(x int, y int) string {
    xStr := strconv.Itoa(x)
    yStr := strconv.Itoa(y)

    return xStr + "," + yStr
}

func toInt(s string) (int, int) {
    num1 := strings.Split(s, ",")[0]
    num2 := strings.Split(s, ",")[1]
    i1, err := strconv.Atoi(num1)
    if err != nil {
        log.Fatal(err)
    }
    i2, err := strconv.Atoi(num2)

    return i1, i2
}

func main() {
    // Open "input.txt"
    buffer, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer buffer.Close()

    graph := [][]int{}

    // Create Scanner to read in file line by line
    scanner := bufio.NewScanner(buffer)
    for scanner.Scan() {
        iArr := []int{}
        for _, c := range(scanner.Text()) {
            i, err := strconv.Atoi(string(c))
            if err != nil {
                log.Fatal(err)
            }
            iArr = append(iArr, i)
        }
        graph = append(graph, iArr)
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }


    goal := toString(len(graph[len(graph) - 1]) - 1,len(graph) - 1)

    fmt.Printf("Total risk: %d\n", findPath(graph, goal))
}

func findPath(graph [][]int, goal string) int {
    for i := 2; i < len(graph); i++ {
        graph[i][0] += graph[i - 1][0]
    }
    for i := 2; i < len(graph); i++ {
        graph[0][i] += graph[0][i - 1]
    }

    for i := 1; i < len(graph); i++ {
        for j := 1; j < len(graph); j++ {
            min := 9999
            if min > graph[i - 1][j] {
                min = graph[i - 1][j]
            }
            if min > graph[i][j - 1] {
                min = graph[i][j - 1]
            }
            graph[i][j] += min
        }
    }

    j, i := toInt(goal)

    return graph[i][j]
}
