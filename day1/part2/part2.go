package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("../input")

	if err != nil {
		fmt.Println("error reading file", err);
	}

	frequencies, err := Parse(input)

	if err != nil {
		fmt.Println("error parsing file", err);
	}

	first := FindFirst(frequencies)

	fmt.Println("first frequency that occurs twice:", first)
}

func Parse(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int

	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())

		if err != nil {
			return result, err
		}

		result = append(result, x)
	}

	return result, scanner.Err()
}

func FindFirst(frequencies []int) int {
	var f = make(map[int]bool)
	f[0] = true

	sum := 0

	found := false

	for !found {
		for i := range frequencies {
			sum += frequencies[i]

			_, ok := f[sum]

			if ok {
				found = true
				break
			}

			f[sum] = true
		}
	}

	return sum
}
