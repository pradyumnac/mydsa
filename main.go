package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Result(data string) {
	fmt.Println(data)
}

func main() {
	input := []string{
		"6",
		"This line has junk text.",
		"121.18.19.20",
		"2001:0db8:0000:0000:0000:ff00:0042:8329",
		"999.999.999.999",
		"-1.127.127.127",
		"2X01:0db8:0000:0000:0000:ff00:0042:8329",
	}

	numLines, _ := strconv.ParseInt(input[0], 10, 64)
	ipv6_pattern, _ := regexp.Compile("[a-f1-90]{4}")

DataItemLabel:
	for i := int64(1); i <= numLines; i = i + 1 {
		line := input[i]

		// if ip4
		parts_ipv4 := strings.Split(line, ".")
		parts_ipv6 := strings.Split(line, ":")

		if len(parts_ipv4) == 4 {
			// ipv4 further check
			for _, part := range parts_ipv4 {
				partNum, err := strconv.ParseInt(part, 10, 32)
				if err != nil {
					Result("Neither")
					continue DataItemLabel
				}

				if partNum <= 0 || partNum > 255 {
					Result("Neither")
					continue DataItemLabel
				}
			}
			Result("IPV4")
		} else if len(parts_ipv6) == 8 {
			for _, part := range parts_ipv6 {
				if len(part) != 4 {
					Result("Neither")
					continue DataItemLabel
				}

				if !ipv6_pattern.Match([]byte(part)) {
					Result("Neither")
					continue DataItemLabel
				}
			}
			Result("IPV6")
		} else {
			fmt.Println("Neither")
		}
	}
}
