package main

import (
	"fmt"
	"internal/stack"
	"regexp"
	"strconv"
)

func historyCalculation(input string, cb func(*stack.Stack[[]int]) int) int {
	history := []int{}
	re := regexp.MustCompile(`[-]?\d+`)
	numbers := re.FindAllString(input, -1)
	for _, n := range numbers {
		num, err := strconv.Atoi(n)
		check(err)
		history = append(history, num)
	}

	valueStack := stack.NewStack[[]int]()
	currentLine := history
	valueStack.Push(currentLine)
	for {
		// calculate the differences between the different entries
		newLine := []int{}
		allZero := true // are all values zero?
		currentValue := currentLine[0]
		for i := 1; i < len(currentLine); i++ {
			newValue := currentLine[i] - currentValue
			newLine = append(newLine, newValue)
			currentValue = currentLine[i]
			// Check for zero values
			if newValue != 0 {
				allZero = false
			}
		}

		// If all values are zero, we are done with this
		if allZero {
			break
		}

		// Save new line to stack and make it the current line
		currentLine = newLine
		valueStack.Push(currentLine)
	}

	// The final value of the calculations
	return cb(valueStack)
}

func historySuccessor(input string) int {
	return historyCalculation(input,
		func(valueStack *stack.Stack[[]int]) int {
			result := 0
			for valueStack.Length() != 0 {
				lastLine := valueStack.Pop()
				result = lastLine[len(lastLine)-1] + result
			}
			return result
		})
}

func historyPredecessor(input string) int {
	return historyCalculation(input,
		func(valueStack *stack.Stack[[]int]) int {
			result := 0
			for valueStack.Length() != 0 {
				lastLine := valueStack.Pop()
				result = lastLine[0] - result
			}
			return result
		})
}

func day09_part1(input []string) int {
	sum := 0
	for _, line := range input {
		if line == "" {
			continue
		}
		sum += historySuccessor(line)
	}
	return sum
}

func day09_part2(input []string) int {
	sum := 0
	for _, line := range input {
		if line == "" {
			continue
		}
		sum += historyPredecessor(line)
	}
	return sum
}

func day09() {
	input := readFileAsLines("data/day09_input.txt")
	result1 := day09_part1(input)
	result2 := day09_part2(input)
	fmt.Printf("Day09 Result1: %d\n", result1)
	fmt.Printf("Day09 Result2: %d\n", result2)
}
