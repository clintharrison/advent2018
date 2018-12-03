package main

import (
    "io/ioutil"
    "flag"
    "fmt"
    "strconv"
    "strings"
)

var inputFile = flag.String("input", "day1.txt", "")
var part2 = flag.Bool("part2", false, "")

func main() {
    flag.Parse()

    bytes, _ := ioutil.ReadFile(*inputFile)
    contents := string(bytes)

    if !*part2 {
        sum := 0
        for _, line := range strings.Split(contents, "\n") {
            if len(line) == 0 { break }

            change, _ := strconv.Atoi(line)
            sum += change
        }

        fmt.Printf("%v\n", sum)
    } else {
        seen := make(map[int]bool)
        sum := 0
        InfiniteLoop:
        for {
            for _, line := range strings.Split(contents, "\n") {
                if len(line) == 0 { break }

                change, _ := strconv.Atoi(line)
                sum += change
                if _, ok := seen[sum]; ok {
                    fmt.Printf("%v\n", sum)
                    break InfiniteLoop
                } else {
                    seen[sum] = true
                }
            }
        }
    }
}
