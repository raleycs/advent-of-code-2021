package main

import(
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

// windowSum takes a slice of ints and
// sums them together and returns it.
func windowSum(window []int) int {
    sum := 0

    for _, v := range window {
        sum += v
    }
    return sum
}

// getMeasures reads the input
// and returns a slice of integers
func getMeasures() []int {
    // Open "input.txt"
    buffer, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer buffer.Close()

    // Create Scanner to read in file line by line
    scanner := bufio.NewScanner(buffer)
    measures := []int{} // slice holding all measurements from input
    for scanner.Scan() {
        v, err := strconv.Atoi(scanner.Text())
        if err != nil {
            log.Fatal(err)
        }
        measures = append(measures, v)
        if err := scanner.Err(); err != nil {
            log.Fatal(err)
        }
    }
    return measures
}

func main() {
    fmt.Printf("Reading in input.txt...\n")

    measures := getMeasures()

    currentWindow := []int{} // temporary var holding current window
    previousWindow := measures[0:3] // temporary var holding previous window
    increases := 0 // var to hold number of increases
    measures = measures[1:] // ignore first input

    for _, v := range(measures) {
        fmt.Printf("Reading in: %d\n", v)

        currentWindow = append(currentWindow, v) // add current input to sliding window

        fmt.Println(previousWindow)
        fmt.Println(currentWindow)

        // sum when new window is full
        if len(currentWindow) == 3 {
            // determine if there was an increase from the previous window
            if windowSum(currentWindow) > windowSum(previousWindow) {
                increases += 1 // increment
                fmt.Printf("Current sum: %d (increased)\n", windowSum(currentWindow))
            } else {
                fmt.Printf("Current sum: %d (no change/decreased)\n", windowSum(currentWindow))
            }

            // set past window to be current window
            previousWindow = currentWindow

            // set current window to be next window
            currentWindow = currentWindow[1:]
        }
        fmt.Println("-----")
    }
    fmt.Printf("Total number of increases: %d\n", increases)
}
