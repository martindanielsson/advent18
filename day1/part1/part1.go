package main

import "fmt"
import "bufio"
import "io"
import "os"
import "strconv"

func main() {
	input, err := os.Open("../input")
	if err != nil {
	    fmt.Println("error reading file", err);
	}
	frequencies, err := Parse(input)
	if err != nil {
	    fmt.Println("error parsing file", err);
	}
	fmt.Println("frequency is", Sum(frequencies))
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

func Sum(input []int) int {
    sum := 0

    for i := range input {
        sum += input[i]
    }

    return sum
}
