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

	if err != nil {
		fmt.Println("error reading file", err);
	}

	barcodes, err := Parse(input)

	if err != nil {
		fmt.Println("error parsing file", err);
	}

	fmt.Println("the common characters are", FindCommon(barcodes))
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

func FindCommon(barcodes []string) string {
	common := ""
	length := len(barcodes)

	for i := range barcodes {
		a := strings.Split(barcodes[i], "")

		for ii := range barcodes {
			if barcodes[i] == barcodes[length-1-ii] {
				break
			}

			if len(barcodes[i]) != len(barcodes[length-1-ii]) {
				continue
			}

			b := strings.Split(barcodes[length-1-ii], "")

			counter := 0

			for iii, c := range a {
				if c != b[iii] {
					counter++
				} else {
					common += c
				}

				if counter > 1 {
					break
				}
			}

			if counter == 1 {
				return common
			} else {
				common = ""
			}
		}

	}

	return common
}