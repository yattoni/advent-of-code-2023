package main

import (
	"fmt"
	"log"
	"strings"

	util "github.com/yattoni/advent-of-code-2023"
)

const LEFT = 'L'
const RIGHT = 'R'

const START = "AAA"
const END = "ZZZ"

type node struct {
	id      string
	left    string
	right   string
	isStart bool
	isEnd   bool
}

func (node node) String() string {
	return fmt.Sprintf("{id=%s,left=%s,right=%s}", node.id, node.left, node.right)
}

func parseInput(lines []string) ([]rune, map[string]node) {
	instructions := []rune(lines[0])

	fmt.Println("instructions", string(instructions))
	fmt.Println("instructions", instructions)

	nodes := make(map[string]node, len(lines)-2)

	for _, line := range lines[2:] {
		id := string(line[0:3])
		left := string(line[7:10])
		right := string(line[12:15])
		nodes[id] = node{id, left, right, strings.HasSuffix(id, "A"), strings.HasSuffix(id, "Z")}
	}
	fmt.Println("nodes", nodes)
	return instructions, nodes
}

func part1(instructions []rune, nodes map[string]node) {
	current := nodes[START]
	steps := 0
	for current.id != END {
		for _, instruction := range instructions {
			if current.id == END {
				break
			}
			steps += 1
			if instruction == LEFT {
				current = nodes[current.left]
			} else if instruction == RIGHT {
				current = nodes[current.right]
			} else {
				log.Fatalf("Unknown instruction, %c", instruction)
			}
		}
	}

	fmt.Println("part 1 steps", steps) // 20093
}

func part2BruteForce(instructions []rune, nodes map[string]node) {
	currentNodes := make([]node, 0)

	for _, n := range nodes {
		if n.isStart {
			currentNodes = append(currentNodes, n)
		}
	}
	fmt.Println("starting nodes", currentNodes)

	steps := 0
	for !allNodesAtEnd(currentNodes) {
		for _, instruction := range instructions {
			if allNodesAtEnd(currentNodes) {
				break
			}
			steps += 1
			if steps%10_000_000 == 0 {
				fmt.Println("current steps", steps)
			}
			for i, n := range currentNodes {
				if instruction == LEFT {
					currentNodes[i] = nodes[n.left]
				} else if instruction == RIGHT {
					currentNodes[i] = nodes[n.right]
				} else {
					log.Fatalf("Unknown instruction, %c", instruction)
				}
			}
		}
	}
	fmt.Println("part 2 steps", steps)
}

func allNodesAtEnd(nodes []node) bool {
	for _, node := range nodes {
		if !node.isEnd {
			return false
		}
	}
	return true
}

func main() {
	lines := util.ReadFileWithSpacesToLines("input")

	part1(parseInput(lines))

	instructions, nodes := parseInput(util.ReadFileWithSpacesToLines("input"))

	startingNodes := make([]node, 0)

	for _, n := range nodes {
		if n.isStart {
			startingNodes = append(startingNodes, n)
		}
	}
	fmt.Println("starting nodes", startingNodes)
	stepsToEndByNode := make(map[string]int, len(startingNodes))
	steps := make([]int, len(startingNodes))

	for i, current := range startingNodes {
		currentSteps := 0
		for !current.isEnd {
			for _, instruction := range instructions {
				if current.isEnd {
					break
				}
				currentSteps += 1
				if instruction == LEFT {
					current = nodes[current.left]
				} else if instruction == RIGHT {
					current = nodes[current.right]
				} else {
					log.Fatalf("Unknown instruction, %c", instruction)
				}
			}
		}
		stepsToEndByNode[startingNodes[i].id] = currentSteps
		steps[i] = currentSteps
	}

	fmt.Println(stepsToEndByNode)
	fmt.Println(steps)
	fmt.Println("part2 steps", lcmMultiple(steps))
	// least common multiple = 22103062509257
}

// https://stackoverflow.com/questions/147515/least-common-multiple-for-3-or-more-numbers
func lcmMultiple(nums []int) int {
	if len(nums) == 2 {
		return lcm(nums[0], nums[1])
	}
	return lcm(nums[0], lcmMultiple(nums[1:]))
}

// https://stackoverflow.com/questions/3154454/what-is-the-most-efficient-way-to-calculate-the-least-common-multiple-of-two-int
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// https://en.wikipedia.org/wiki/Euclidean_algorithm#Implementations
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
