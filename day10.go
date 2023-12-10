package main

import (
	"fmt"
)

type mapCoord struct {
	x int
	y int
}

type mapNode struct {
	character rune
	coord     mapCoord
	north     bool
	east      bool
	south     bool
	west      bool
}

func parseMap(input []string) *map[mapCoord]mapNode {
	result := map[mapCoord]mapNode{}

	// insert all nodes into the node map
	for y, line := range input {
		for x, r := range line {
			mapnode := mapNode{
				character: r,
				coord:     mapCoord{x, y},
				north:     r == '|' || r == 'L' || r == 'J',
				east:      r == '-' || r == 'L' || r == 'F',
				south:     r == '|' || r == '7' || r == 'F',
				west:      r == '-' || r == 'J' || r == '7',
			}
			result[mapCoord{x, y}] = mapnode
		}
	}

	for v, k := range result {
		if k.character == 'S' {
			// Find the directions of the starting node
			k.north = result[mapCoord{k.coord.x, k.coord.y - 1}].south
			k.east = result[mapCoord{k.coord.x + 1, k.coord.y}].west
			k.south = result[mapCoord{k.coord.x, k.coord.y + 1}].north
			k.west = result[mapCoord{k.coord.x - 1, k.coord.y}].east
			result[v] = k
			break
		}
	}

	return &result
}

func findStartingNode(nodeMap *map[mapCoord]mapNode) mapCoord {
	for v, k := range *nodeMap {
		if k.character == 'S' {
			return v
		}
	}
	panic("No starting node found")
}

func followPath(nodeMap *map[mapCoord]mapNode) []mapCoord {
	startingCoords := findStartingNode(nodeMap)
	currentCoords := startingCoords
	currentNode := (*nodeMap)[currentCoords]
	lastDirection := ""
	path := []mapCoord{}
	path = append(path, currentCoords)
	for {
		if currentNode.north && lastDirection != "s" {
			currentCoords = mapCoord{currentCoords.x, currentCoords.y - 1}
			lastDirection = "n"
		} else if currentNode.east && lastDirection != "w" {
			currentCoords = mapCoord{currentCoords.x + 1, currentCoords.y}
			lastDirection = "e"
		} else if currentNode.south && lastDirection != "n" {
			currentCoords = mapCoord{currentCoords.x, currentCoords.y + 1}
			lastDirection = "s"
		} else if currentNode.west && lastDirection != "e" {
			currentCoords = mapCoord{currentCoords.x - 1, currentCoords.y}
			lastDirection = "w"
		} else {
			panic("Nowhere to go")
		}

		currentNode = (*nodeMap)[currentCoords]

		if currentNode.character == 'S' {
			break
		}

		path = append(path, currentCoords)
	}
	return path
}

func pathContains(path []mapCoord, pos mapCoord) bool {
	for _, p := range path {
		if p == pos {
			return true
		}
	}
	return false
}

func findEmptyNodes(nodeMap *map[mapCoord]mapNode, path []mapCoord) []mapCoord {
	result := []mapCoord{}
	for k, v := range *nodeMap {
		if !pathContains(path, v.coord) {
			result = append(result, k)
		}
	}
	return result
}

func legalPiece(pos mapCoord, nodeMap *map[mapCoord]mapNode, path []mapCoord) bool {
	north, south, east, west := false, false, false, false

	up, up_ok := (*nodeMap)[mapCoord{pos.x, pos.y - 1}]
	down, down_ok := (*nodeMap)[mapCoord{pos.x, pos.y + 1}]
	left, left_ok := (*nodeMap)[mapCoord{pos.x - 1, pos.y}]
	right, right_ok := (*nodeMap)[mapCoord{pos.x + 1, pos.y}]

	if up_ok && pathContains(path, up.coord) && up.south {
		north = true
	}
	if down_ok && pathContains(path, down.coord) && down.north {
		south = true
	}
	if left_ok && pathContains(path, left.coord) && left.east {
		west = true
	}
	if right_ok && pathContains(path, right.coord) && right.west {
		east = true
	}

	return (west && east) || (north && south) || (north && west) || (south && east)
}

func isNodeInsideOfPath(node mapCoord, path []mapCoord, nodeMap *map[mapCoord]mapNode) bool {
	limit := 10000

	count := 0
	for x, y := node.x, node.y; x < limit && y < limit; x, y = x+1, y+1 {
		currentPos := mapCoord{x, y}
		for _, p := range path {
			if currentPos == p && legalPiece(currentPos, nodeMap, path) {
				count += 1
				break
			}
		}
	}

	return count%2 != 0
}

func day10_part1(input []string) int {
	nodeMap := parseMap(input)
	path := followPath(nodeMap)
	return len(path) / 2
}

func day10_part2(input []string) int {
	nodeMap := parseMap(input)
	path := followPath(nodeMap)
	emptyNodes := findEmptyNodes(nodeMap, path)
	result := 0
	for _, node := range emptyNodes {
		if isNodeInsideOfPath(node, path, nodeMap) {
			result += 1
		}
	}
	return result
}

func day10() {
	input := readFileAsLines("data/day10_input.txt")
	fmt.Printf("day 10 result 1: %d\n", day10_part1(input))
	fmt.Printf("day 10 result 2: %d\n", day10_part2(input))
}
