package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	input, err := os.Open("../input")

	barcodes, err := Parse(input)

	if err != nil {
		fmt.Println("error reading file", err);
	}

	fmt.Println("barcode checksum", CountBarcodes(barcodes))
}

func Parse(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []string

	for scanner.Scan() {
		x := scanner.Text()
		result = append(result, x)
	}

	return result, scanner.Err()
}

func CountBarcodes(barcodes []string) int {
	two := 0
	three := 0

	for i := range barcodes {
		hasTwo := false
		hasThree := false
		split := strings.Split(barcodes[i], "")

		chars := make(map[string]int)

		for s := range split {
			chars[split[s]]++
		}

		for c := range chars {
			if chars[c] == 2 && !hasTwo {
				hasTwo = true
				two++;
			} else if chars[c] == 3 && !hasThree {
				hasThree = true
				three++
			}
		}
	}

	return two * three
}