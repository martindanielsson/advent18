package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

type claim struct {
	id int
	left int
	top int
	width int
	height int
}

func main() {
	input, err := os.Open("../input")

	if err != nil {
		fmt.Println("error reading file", err);
	}

	claims, err := Parse(input)

	if err != nil {
		fmt.Println("error parsing file", err);
	}

	fmt.Println("contested square inches:", FindOverlaps(claims))
}

func Parse(r io.Reader) ([]claim, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []claim
	re := regexp.MustCompile(`\#(\d{1,4})\s@\s(\d{1,3}),(\d{1,3}):\s(\d{1,3})x(\d{1,3})`)

	for scanner.Scan() {
		x := scanner.Text()
		m := re.FindStringSubmatch(x)

		result = append(result, ClaimFromMatch(m))
	}

	return result, scanner.Err()
}

func ClaimFromMatch(m []string) claim {
	var ints []int

	for i, v := range m {
		if i == 0 {
			continue
		}

		n, err := strconv.Atoi(v)

		if err != nil {
			panic(err)
		}

		ints = append(ints, n)
	}

	c := claim{ints[0], ints[1], ints[2], ints[3], ints[4]}

	return c
}

func FindOverlaps(claims []claim) int {
	overlaps := 0
	fabric := [1000][1000]int{}

	for _, c := range claims {
		for y := c.top; y < c.top+c.height; y++ {
			for x := c.left; x < c.left+c.width; x++ {
				if fabric[y][x] == 1 {
					overlaps++
				}

				fabric[y][x]++
			}
		}
	}

	return overlaps
}