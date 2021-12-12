package main
import(
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "unicode"
)

var total = 0 // global int to tally total number of paths founds

// isUpper determines if all characters in a string are upper case
func isUpper(s string) bool {
    for _, r := range s {
        if !unicode.IsUpper(r) && unicode.IsLetter(r) {
            return false
        }
    }
    return true
}

// traverse is a recursive function that uses depth-first search for finding all paths
func traverse(node string, visited map[string]int, graph map[string][]string, path []string) {

    visited[node] += 1 // tell visited map that we have visited the current node

    path = append(path, node) // add node to current path

    // if we reached the end, print path and add one to total number of paths found
    if node == "end" {
        fmt.Println("[*] " + strings.Join(path, ","))
        total += 1
        return
    }

    // for each neighboring node, call traverse on that node
    for _, edge := range(graph[node]) {
        dst := strings.Split(edge, "-")[1] // get destination node from edge

        // fmt.Printf("Currently @ node %s\n", node)
        // fmt.Printf("Trying to visit node %s\n", dst)

        // check for small caves
        if dst != "end" && dst != "start" && !isUpper(dst) {
            // ensure that for each small cave, we are visiting once
            if visited[dst] < 1 {
                traverse(dst, visited, graph, path)
            }
        } else if dst != "start" {
            traverse(dst, visited, graph, path)
        }
    }

    // remove last element from path
    path = path[:len(path) - 1]
    visited[node] -= 1
}

func main() {
    // Open "input.txt"
    buffer, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer buffer.Close()

    graph := make(map[string][]string) // graph mapping nodes to all available edges

    // Create Scanner to read in file line by line
    scanner := bufio.NewScanner(buffer)
    for scanner.Scan() {
        edges := strings.Split(scanner.Text(), "-")

        // obtain nodes from edges
        sourceNode := edges[0]
        destinationNode := edges[1]

        // add edges to graph
        graph[sourceNode] = append(graph[sourceNode], scanner.Text())
        graph[destinationNode] = append(graph[destinationNode], destinationNode + "-" + sourceNode)

        if err := scanner.Err(); err != nil {
            log.Fatal(err)
        }

    }

    visited := make(map[string]int) // map to determine number of times a node was visited
    path := []string{} // slice to hold nodes that make up a path to the end

    traverse("start", visited, graph, path) // traverse from the start to find all paths that lead to the end

    fmt.Println(total)
}
