package main

import (
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/wellsjo/advent-of-code-2021/common"
)

func main() {
	p := parseInput("input.txt")

	start := time.Now()
	p.traverse(p.start, map[string]struct{}{}, map[string]int{}, []*node{}, true)
	end := time.Now()

	fmt.Println("Found", len(p.foundPaths), "Paths", end.Sub(start))
}

type puzzle struct {
	start      *node
	foundPaths map[string]struct{}
	nodes      map[string]*node
}

func (p puzzle) traverse(
	currentNode *node,
	noReturn map[string]struct{},
	nodeCounts map[string]int,
	path []*node,
	canRevisitSmallNode bool,
) {
	if currentNode.name == "end" {
		pathStr := pathToString(path)
		p.foundPaths[pathStr] = struct{}{}
		return
	}
	if _, ok := noReturn[currentNode.name]; ok {
		return
	}

	noReturnCopy := copyNoReturnMap(noReturn)
	nodeCountsCopy := copyCountMap(nodeCounts)

	if currentNode.role == start {
		noReturnCopy[currentNode.name] = struct{}{}
	}

	if currentNode.role == small {
		switch nodeCounts[currentNode.name] {
		case 0:
			nodeCountsCopy[currentNode.name]++
			if !canRevisitSmallNode {
				noReturnCopy[currentNode.name] = struct{}{}
			}

		case 1:
			if canRevisitSmallNode {
				noReturnCopy[currentNode.name] = struct{}{}
				canRevisitSmallNode = false
			} else {
				return
			}

		default:
			panic("invalid state")
		}
	}

	for i := range currentNode.connections {
		p.traverse(
			currentNode.connections[i],
			noReturnCopy,
			nodeCountsCopy,
			append(path, currentNode),
			canRevisitSmallNode,
		)
	}
}

func pathToString(path []*node) string {
	str := ""
	for i := range path {
		str += path[i].name
	}
	return str
}

func copyCountMap(original map[string]int) map[string]int {
	newMap := map[string]int{}
	for k, v := range original {
		newMap[k] = v
	}
	return newMap
}

func copyNoReturnMap(original map[string]struct{}) map[string]struct{} {
	newMap := map[string]struct{}{}
	for k, v := range original {
		newMap[k] = v
	}
	return newMap
}

func (p puzzle) print() {
	for _, node := range p.nodes {
		fmt.Println(node)
	}
}

type node struct {
	name        string
	connections []*node
	role        role
}

func (n node) String() string {
	s := n.name + "->"
	for i := range n.connections {
		s += fmt.Sprintf("%v(%v) ",
			n.connections[i].name, n.connections[i].role)
	}
	return s
}

type role string

const (
	start role = "start"
	end   role = "end"
	small role = "small"
	big   role = "big"
)

func parseInput(path string) puzzle {
	lines, err := common.ReadLines(path)
	if err != nil {
		panic(err)
	}

	nodes := map[string]*node{}

	for i := range lines {
		lineParts := strings.Split(lines[i], "-")
		nodeName := lineParts[0]
		connectionName := lineParts[1]

		n, exists := nodes[nodeName]
		if !exists {
			n = makeNode(nodeName)
		}

		n2, exists := nodes[connectionName]
		if !exists {
			n2 = makeNode(connectionName)
		}

		n.connections = append(n.connections, n2)
		n2.connections = append(n2.connections, n)
		nodes[nodeName] = n
		nodes[connectionName] = n2
	}

	return puzzle{
		start:      nodes["start"],
		nodes:      nodes,
		foundPaths: map[string]struct{}{},
	}
}

func makeNode(name string) *node {
	return &node{
		name:        name,
		connections: []*node{},
		role:        getRole(name),
	}
}

func getRole(s string) role {
	if s == "start" {
		return start
	}
	if s == "end" {
		return end
	}
	if unicode.IsLower(rune(s[0])) {
		return small
	}
	if unicode.IsUpper(rune(s[0])) {
		return big
	}
	panic("invalid input")
}
