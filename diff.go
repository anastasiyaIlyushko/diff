package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func difference(slice1 []int, slice2 []int) []int {
	var diff []int

	for _, a := range slice1{
		fl := false

		for _, b := range slice2{
			if a == b{
				fl = true
				break
			}
		}
		
		if !fl {
			diff = append(diff, a)
			fmt.Printf("%+v\n", a)
		}		
	}

	return diff
}

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadInts(r io.Reader) ([]int, error) {
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

func main() {
	//slice1 := []string{"foo", "bar", "hello"}
	//slice2 := []string{"foo", "world", "bar", "foo"}

	f1, err := os.Open(os.Args[1])
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", os.Args[1], err))
	}
	defer f1.Close()

	f2, err := os.Open(os.Args[2])
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", os.Args[2], err))
	}
	defer f2.Close()

	ints1, err := ReadInts(f1)

	ints2, err := ReadInts(f2)

	difference(ints1, ints2)
}
