package main

import (
	"fmt"
	"strings"
)

type day12_node struct {
	result   int
	valid    bool
	children [16]*day12_node
}

var day12_cache map[string]*day12_node

func fetchCacheResult(node *day12_node, nums []int) (int, bool) {
	if len(nums) == 0 {
		return node.result, node.valid
	} else {
		if node.children[nums[0]] == nil {
			return 0, false
		}

		return fetchCacheResult(node.children[nums[0]], nums[1:])
	}
}

func fetchCacheValue(input string, nums []int) (int, bool) {
	node, ok := day12_cache[input]

	if ok {
		value, ok := fetchCacheResult(node, nums)

		if ok {
			return value, true
		}
	}

	return 0, false
}

func storeCacheResult(node *day12_node, nums []int, result int) {
	if len(nums) == 0 {
		node.result = result
		node.valid = true
	} else {
		if node.children[nums[0]] == nil {
			node.children[nums[0]] = &day12_node{}
		}
		storeCacheResult(node.children[nums[0]], nums[1:], result)
	}
}

func storeCacheValue(input string, nums []int, result int) {
	if day12_cache == nil {
		day12_cache = map[string]*day12_node{}
	}

	val, ok := day12_cache[input]
	if !ok {
		val = &day12_node{}
		day12_cache[input] = val
	}
	storeCacheResult(val, nums, result)
}

func day12_count(input string, nums []int) int {
	numlen := len(nums)
	inputlen := len(input)

	if input == "" {
		if numlen == 0 {
			return 1
		} else {
			return 0
		}
	}

	if numlen == 0 {
		if strings.ContainsAny(input, "#") {
			return 0
		} else {
			return 1
		}
	}

	result, isCached := fetchCacheValue(input, nums)

	if isCached {
		return result // shortcut :D
	}

	if strings.ContainsAny(".?", string(input[0])) {
		result += day12_count(input[1:], nums)
	}

	if strings.ContainsAny("#?", string(input[0])) {
		if nums[0] <= inputlen &&
			!strings.ContainsAny(input[:nums[0]], ".") &&
			(nums[0] == inputlen || input[nums[0]] != '#') {
			result += day12_count(input[min(nums[0]+1, inputlen):], nums[1:])
		}
	}

	storeCacheValue(input, nums, result)

	return result
}

func day12_part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		splits := strings.Split(line, " ")
		input := splits[0]
		nums := []int{}
		for _, n := range strings.Split(splits[1], ",") {
			nums = append(nums, parseInt(n))
		}

		sum += day12_count(input, nums)
	}
	return sum
}

func day12_part2(lines []string) int {
	sum := 0
	for _, line := range lines {
		splits := strings.Split(line, " ")
		input := splits[0]
		input = strings.Join([]string{input, input, input, input, input}, "?")

		nums := []int{}
		for _, n := range strings.Split(splits[1], ",") {
			nums = append(nums, parseInt(n))
		}

		buf := []int{}
		for i := 0; i < 5; i++ {
			buf = append(buf, nums...)
		}

		nums = buf

		sum += day12_count(input, nums)
	}
	return sum
}

func day12() {
	input := readFileAsLines("data/day12_input.txt")

	result1 := day12_part1(input)
	fmt.Printf("Day12 part1 result: %d\n", result1)

	result2 := day12_part2(input)
	fmt.Printf("Day12 part2 result: %d\n", result2)
}
