package main

import(
    "bufio"
    "container/heap"
    "fmt"
    "log"
    "math"
    "os"
    "strconv"
    "strings"
)

// An Item is something we manage in a priority queue.
type Item struct {
    coordinates string
    priority int    // The priority of the item in the queue.
    // The index is needed by update and is maintained by the heap.Interface methods.
    index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
    // We want Pop to give us lowest, not highest, priority so we use greater than here.
    return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].index = i
    pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
    n := len(*pq)
    item := x.(*Item)
    item.index = n
    *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    old[n-1] = nil  // avoid memory leak
    item.index = -1 // for safety
    *pq = old[0 : n-1]
    return item
}

// update modifies the priority of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
    item.coordinates = value
    item.priority = priority
    heap.Fix(pq, item.index)
}

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

func expandGraph(grid [][]int) [][]int {
    tiles := [][]int{}
    length := len(grid)
    for n := 0; n < 5; n++ {
        if n == 0 {
            tiles = append(tiles, grid...)
        }
        for i := 0; i < len(grid); i++ {
            tmp := []int{}
            for j := 0; j < len(grid); j++ {
                risk := -1
                if n == 0 {
                    risk = grid[i][j] + n + 1
                } else {
                    risk = grid[i][j] + n
                }
                if risk >= 10 {
                    risk = (risk % 9)
                }
                tmp = append(tmp, risk)
            }
            if n != 0 {
                tiles = append(tiles, tmp)
            }
        }
    }

    for x := 0; x < 5; x++ {
        for n := 1; n < 5; n++ {
            for i := 0; i < len(grid); i++ {
                row := []int{}
                for j := 0; j < len(grid); j++ {
                    risk := grid[i][j] + x + n
                    if risk >= 10 {
                        risk = (risk % 9)
                    }
                    row = append(row, risk)
                }
                tiles[i + length * x] = append(tiles[i + length * x], row...)
            }
        }
    }

    return tiles
}

func main() {
    // Open "input.txt"
    buffer, err := os.Open("input.txt")
    // buffer, err := os.Open("test.txt")
    // buffer, err := os.Open("simple.txt")
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

    graph = expandGraph(graph)

    goal := toString(len(graph[len(graph) - 1]) - 1,len(graph) - 1)

    fmt.Printf("Total risk: %d\n", findPath(graph, goal))
}

func findPath(graph [][]int, goal string) int {

    dp := [][]int{}
    visited := [][]bool{}

    for i := 0; i < len(graph); i++ {
        row := []int{}
        visitedRow := []bool{}
        for j := 0; j < len(graph); j++ {
            visitedRow = append(visitedRow, false)
            row = append(row, math.MaxInt)
        }
        dp = append(dp, row)
        visited = append(visited, visitedRow)
    }

    dp[0][0] = 0 // don't count starting point risk score

    var pq PriorityQueue
    heap.Init(&pq)

    item := &Item {
        coordinates: "0,0",
        priority: 0,
    }

    heap.Push(&pq, item)

    for pq.Len() > 0 {

        item := heap.Pop(&pq).(*Item)
        x, y := toInt(item.coordinates)

        if visited[x][y] {
            continue
        }

        visited[x][y] = true

        // up
        upX := x + 1
        upY := y

        if upX < len(graph) && visited[upX][upY] == false {
            min := dp[upX][upY]

            if min > dp[x][y] + graph[upX][upY] {
                min = dp[x][y] + graph[upX][upY]
            }

            dp[upX][upY] = min

            item = &Item {
                coordinates: toString(upX, upY),
                priority: dp[upX][upY],
            }

            heap.Push(&pq, item)
        }

        // down
        downX := x - 1
        downY := y

        if downX >= 0 && visited[downX][downY] == false {
            min := dp[downX][downY]

            if min > dp[x][y] + graph[downX][downY] {
                min = dp[x][y] + graph[downX][downY]
            }

            dp[downX][downY] = min

            item = &Item {
                coordinates: toString(downX, downY),
                priority: dp[downX][downY],
            }

            heap.Push(&pq, item)
        }

        // right
        rightX := x
        rightY := y + 1

        if rightY < len(graph) && visited[rightX][rightY] == false {
            min := dp[rightX][rightY]

            if min > dp[x][y] + graph[rightX][rightY] {
                min = dp[x][y] + graph[rightX][rightY]
            }

            dp[rightX][rightY] = min

            item = &Item {
                coordinates: toString(rightX, rightY),
                priority: dp[rightX][rightY],
            }

            heap.Push(&pq, item)
        }

        // left
        leftX := x
        leftY := y - 1
        if leftY >= 0 && visited[leftX][leftY] == false {
            min := dp[leftX][leftY]

            if min > dp[x][y] + graph[leftX][leftY] {
                min = dp[x][y] + graph[leftX][leftY]
            }

            dp[leftX][leftY] = min

            item = &Item {
                coordinates: toString(leftX, leftY),
                priority: dp[leftX][leftY],
            }

            heap.Push(&pq, item)
        }
    }

    goalX, goalY := toInt(goal)

    return dp[goalX][goalY]
}
