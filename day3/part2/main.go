package main

import(
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

// readInput reads input.txt and
// returns all binaries in a slice
func readInput()[]string {
    // Open "input.txt"
    buffer, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer buffer.Close()


    inputs := []string{} // slice of binaries to return

    // Create Scanner to read in file line by line
    scanner := bufio.NewScanner(buffer)
    for scanner.Scan() {
        inputs = append(inputs, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return inputs
}

func main() {
    fmt.Printf("Reading in input.txt...\n")

    inputs := readInput()

    // create list of maps to keep track of most common bits in each place
    list := make([]map[rune]int, 12)
    for i := 0; i < 12; i++ {
        m := make(map[rune]int)
        list[i] = m
    }

    // parse through binaries from input
    for _, input := range(inputs) {
        for index, bit := range(input) {
            list[index][bit] += 1
        }
    }

    // create tmp list to prevent changes to original slice
    tmp := []string{}
    for _, input := range(inputs) {
        tmp = append(tmp, input)
    }

    // filter out results to find oxygen generator rating
    for i := 0; i < 12; i++ {
        // 1 is the common bit
        if list[i]['1'] >= list[i]['0'] {
            for index, binary := range(tmp) {
                if tmp[index] != "!" && binary[i] == '0' {
                    tmp[index] = "!"
                }
            }
        } else {
            // 0 is the common bit
            for index, binary := range(tmp) {
                if tmp[index] != "!" && binary[i] == '1' {
                    tmp[index] = "!"
                }
            }
        }
        // check if there is only 1 number left
        left := 0
        for _, binary := range(tmp) {
            if binary != "!" {
                left += 1
            }
        }
        if left == 1 {
            break
        }

        // clear initial mappings
        for index, _ := range(list) {
            list[index]['0'] = 0
            list[index]['1'] = 0
        }
        // re-parse and update common bits
        for _, binary := range(tmp) {
            if binary != "!" {
                for index, bit := range(binary) {
                    list[index][bit] += 1
                }
            }
        }
    }

    // find oxygen rating from remaining values
    oxygenRating := ""
    for _, rating := range(tmp) {
        if rating != "!" {
            oxygenRating = rating
            break
        }
    }
    fmt.Printf("Oxygen Generator Rating: %s\n", oxygenRating)

    // reset tmp list to prevent changes to original slice
    for index, input := range(inputs) {
        tmp[index] = input
    }

    // filter out results to find co2 scrubber rating
    for i := 0; i < 12; i++ {
        // 1 is the common bit
        if list[i]['1'] >= list[i]['0'] {
            for index, binary := range(tmp) {
                if tmp[index] != "!" && binary[i] == '1' {
                    tmp[index] = "!"
                }
            }
        } else {
            // 0 is the common bit
            for index, binary := range(tmp) {
                if tmp[index] != "!" && binary[i] == '0' {
                    tmp[index] = "!"
                }
            }
        }

        // check if there is only 1 number left
        left := 0
        for _, binary := range(tmp) {
            if binary != "!" {
                left += 1
            }
        }
        // if left == 2 {
        //     fmt.Println(i)
        //     for _, binary := range(tmp) {
        //         if binary != "!" {
        //             fmt.Println(binary)
        //         }
        //     }
        // }
        if left == 1 {
            // fmt.Println(i)
            break
        }

        // clear initial mappings
        for index, _ := range(list) {
            list[index]['0'] = 0
            list[index]['1'] = 0
        }
        // re-parse and update common bits
        for _, binary := range(tmp) {
            if binary != "!" {
                for index, bit := range(binary) {
                    list[index][bit] += 1
                }
            }
        }
    }

    // find co2 scrubber rating from remaining values
    scrubberRating := ""
    for _, rating := range(tmp) {
        if rating != "!" {
            scrubberRating = rating
            break
        }
    }
    fmt.Printf("CO2 Scrubber Rating: %s\n", scrubberRating)

    // Convert strings to ints
    o2Rating, err := strconv.ParseInt(oxygenRating, 2, 64)
    if err != nil {
        log.Fatal(err)
    }
    co2Rating, err := strconv.ParseInt(scrubberRating, 2, 64)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Life Support Rating: %d\n", o2Rating * co2Rating)
}
