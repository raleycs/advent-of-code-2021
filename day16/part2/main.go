package main

import(
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

func parseBinary(binary string) {
    var sum int64
    sum = 0

    // parse until we reach the end
    // for len(binary) >= 1 {

    version, id := getHeader(binary)
    sum += version

    // literal value
    if id == 4 {
        var literal int64
        var size int64
        literal, size = parseLiteral(binary)
        fmt.Printf("Literal extracted: %d\n", literal)
        for size % 8 != 0 {
            size += 1
        }
    } else {
        // operator
        versionTotal, _ := parseOp(binary)
        sum += versionTotal
    }
    fmt.Printf("Sum of versions: %d\n", sum)
    fmt.Println("--------------------")
}

func getHeader(binary string) (int64, int64) {
    // extract header
    version := binary[0:3]
    versionInt, err := strconv.ParseInt(version, 2, 64)
    if err != nil {
        log.Fatal(err)
    }
    id := binary[3:6]
    idInt, err := strconv.ParseInt(id, 2, 64)
    if err != nil {
        log.Fatal(err)
    }

    return versionInt, idInt
}

func parseOp(binary string) (int64, int64) {
    var sum int64
    var packetSize int64
    lengthType := binary[6]

    // next 15 bits are a number that represents the total length in bits of the sub-packets
    if string(lengthType) == "0" {
        length, err := strconv.ParseInt(binary[7:22], 2, 64)
        if err != nil {
            log.Fatal(err)
        }
        var total int64
        binary = binary[22:] // truncate header + length bits
        packetSize += 22

        for {
            version, id := getHeader(binary)
            sum += version

            if id == 4 {
                literal, size := parseLiteral(binary)
                fmt.Printf("Literal extracted: %d\n", literal)
                binary = binary[size:]
                total += size
                packetSize += size
            } else {
                versionTotal, size := parseOp(binary)
                binary = binary[size:]
                sum += versionTotal
                total += size
                packetSize += size
            }
            if total == length {
                break
            }
        }
    } else if string(lengthType) == "1" {
        // next 11 bits are a number that represents the number of sub-packets immediately contained by this packet
        length, err := strconv.ParseInt(binary[7:18], 2, 64)
        if err != nil {
            log.Fatal(err)
        }

        binary = binary[18:] // truncate header + length bits
        var parsed int64

        packetSize += 18

        for i := int64(0); i < length; i++ {
            version, id := getHeader(binary)
            sum += version

            if id == 4 {
                literal, size := parseLiteral(binary)
                fmt.Printf("Literal extracted: %d\n", literal)
                binary = binary[size:]
                parsed += size
                packetSize += size
            } else {
                versionTotal, size := parseOp(binary)
                sum += versionTotal
                binary = binary[size:]
                parsed += size
                packetSize += size
            }
        }
    }
    return sum, packetSize
}

func parseLiteral(binary string) (int64, int64) {
    groups := []string{} // group of binary to put together literal
    packet := binary[6:] // skip header
    i := 1
    var size int64
    size = 6
    // parse 3 groups of bits
    for i != 0 {
        groups = append(groups, packet[(i - 1) * 5:i * 5][1:])
        size += 5

        // last group
        if string(packet[(i - 1) * 5:i * 5][0]) == "0" {
            break
        }
        i += 1
    }

    literal, err := strconv.ParseInt(strings.Join(groups, ""), 2, 64)
    if err != nil {
        log.Fatal(err)
    }
    return literal, size
}

func main() {
    // Open "input.txt"
    buffer, err := os.Open("input.txt")
    // buffer, err := os.Open("test.txt")
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

        parseBinary(binary)
    }
}
