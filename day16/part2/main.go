package main

import(
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func doOp(id int64, values []int64) int64 {
    // sum
    if id == 0 {
        var sum int64
        for _, v := range(values) {
            sum += v
        }
        return sum
    }
    // product
    if id == 1 {
        var total = values[0]
        for i := 1; i < len(values); i++ {
            total *= values[i]
        }
        return total
    }
    // min
    if id == 2 {
        var min = values[0]
        for i := 1; i < len(values); i++ {
            if min > values[i] {
                min = values[i]
            }
        }
        return min
    }
    // max
    if id == 3 {
        var max = values[0]
        for i := 1; i < len(values); i++ {
            if max < values[i] {
                max = values[i]
            }
        }
        return max
    }
    // literal
    if id == 4 {
        return values[0]
    }
    // greater than
    if id == 5 {
        if values[0] > values[1] {
            return 1
        } else {
            return 0
        }
    }
    // less than
    if id == 6 {
        if values[0] < values[1] {
            return 1
        } else {
            return 0
        }
    }
    // equal to
    if id == 7 {
        if values[0] == values[1] {
            return 1
        } else {
            return 0
        }
    }
    return 99999
}

func parseBinary(binary string, index int64) (int64, int64) {
    var results []int64
    id, err := strconv.ParseInt(binary[index + 3:index + 6], 2, 64)
    if err != nil {
        log.Fatal(err)
    }
    index += 6

    // literal value
    if id == 4 {
        results = []int64{0}

        for {
            results[0] *= 16
            val, err := strconv.ParseInt(binary[index + 1:index + 5], 2, 64)
            if err != nil {
                log.Fatal(err)
            }
            results[0] += val
            index += 5

            if binary[index - 5] == '0' {
                return index, doOp(id, results)
            }
        }
    } else {
        results = []int64{}
        if binary[index] == '0' {
            endIndex, err := strconv.ParseInt(binary[index + 1:index + 16], 2, 64)
            if err != nil {
                log.Fatal(err)
            }
            endIndex += index + 16
            index += 16

            var res int64
            for index < endIndex {
                index, res = parseBinary(binary, index)
                results = append(results, res)
            }
            return index, doOp(id, results)
        } else {
            numPackets, err := strconv.ParseInt(binary[index + 1:index + 12], 2, 64)
            if err != nil {
                log.Fatal(err)
            }
            index += 12

            var res int64
            for i := int64(0); i < numPackets; i++ {
                index, res = parseBinary(binary, index)
                results = append(results, res)
            }
        }
        return index, doOp(id, results)
    }
}

func main() {
    // Open "input.txt"
    buffer, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer buffer.Close()

    encoded := []string{}

    // Create Scanner to read in file line by line
    scanner := bufio.NewScanner(buffer)
    for scanner.Scan() {
        encoded = append(encoded, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    for i, e := range(encoded) {
        fmt.Printf("Parsing packet %d\n", i + 1)
        // convert hex to binary
        binary := ""
        for _, digit := range(e) {
            n, err := strconv.ParseInt(string(digit), 16, 64)
            if err != nil {
                log.Fatal(err)
            }
            bin := strconv.FormatInt(n, 2)
            for len(bin) != 4 {
                bin = "0" + bin
            }
            binary += bin
        }
        _, result := parseBinary(binary, int64(0))
        fmt.Printf("Solution: %d\n", result)
    }
}
