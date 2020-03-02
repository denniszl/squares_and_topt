package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sums = []int{}

// getSquare requirement:
/*
For each test case, calculate the sum of squares of the integers, excluding any negatives, and print the calculated sum in the output.
*/
func getSquare(in string) (int, error) {
	num, err := strconv.Atoi(strings.TrimSuffix(in, "\n"))
	if err != nil {
		return 0, err
	}

	// exclude negatives
	if num <= 0 {
		return 0, nil
	}

	// return square
	return num * num, nil
}

// sumNums sums a line of numbers based on requirements given
func sumNums(numsRemaining int, numbers []string) (int, error) {
	if numsRemaining <= 0 {
		return 0, nil
	}

	sq, err := getSquare(numbers[0])
	if err != nil {
		return 0, err
	}
	numsRemaining--
	// recursion is a bit ugly here but this guarantees code safety due to handling error
	next, err := sumNums(numsRemaining, numbers[1:])
	return sq + next, err
}

// readTestInput recursively reads input as long as testsRemaining > 0
func readTestInput(testsRemaining int) error {
	if testsRemaining <= 0 {
		return nil
	}

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("got error %s", err)
	}

	numbers, err := strconv.Atoi(strings.TrimSuffix(text, "\n"))
	if err != nil {
		return fmt.Errorf("failed to parse input %s, must be integer", strings.TrimSuffix(text, "\n"))
	}
	reader = bufio.NewReader(os.Stdin)
	nums, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("got error %s", err)
	}
	splitNums := strings.Split(strings.TrimSuffix(nums, "\n"), " ")

	s, err := sumNums(numbers, splitNums)
	if err != nil {
		return fmt.Errorf("got error while attempting to sum input %s", err)
	}
	sums = append(sums, s)

	testsRemaining--
	return readTestInput(testsRemaining)
}

// prints sums
func printSums(sums []int) {
	if len(sums) <= 0 {
		return
	}

	fmt.Println(sums[0])
	printSums(sums[1:])
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(fmt.Sprintf("got error %s", err))
	}

	nTests, err := strconv.Atoi(strings.TrimSuffix(text, "\n"))
	if err != nil {
		panic(fmt.Sprintf("failed to parse input %s, must be integer", strings.TrimSuffix(text, "\n")))
	}

	err = readTestInput(nTests)
	if err != nil {
		panic(fmt.Sprintf("error when reading input %s", err))
	}

	printSums(sums)
	return
}
