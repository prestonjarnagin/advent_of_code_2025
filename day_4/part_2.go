package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Grid struct {
	lines  []string
	width  int
	height int
}

func NewGrid(content string) *Grid {
	lines := strings.Split(strings.TrimSpace(content), "\n")
	width := len(lines[0])
	height := len(lines)
	return &Grid{lines: lines, width: width, height: height}
}

func (g *Grid) PrintGrid() {
	for _, line := range g.lines {
		fmt.Println(line)
	}
}

func (g *Grid) ContentAt(x int, y int) string {
	if x < 0 || y < 0 {
		return ""
	}

	if x >= g.width || y >= g.height {
		return ""
	}

	return string(g.lines[y][x])
}

func (g *Grid) GetAccessibleRollCount() int {
	accessible_rolls := 0
	for y := 0; y < g.height; y++ {
		accessible_rolls += g.GetAccessibleRollCountForLine(y)
	}
	return accessible_rolls
}

func (g *Grid) GetAccessibleRollCountForLine(line int) int {
	accessible_rolls := 0
	for x := 0; x < g.width; x++ {
		content := g.ContentAt(x, line)
		if content == "." {
			continue
		}
		if content == "@" || content == "x" {
			if g.CoordinatesAreAccessible(x, line) {
				accessible_rolls += 1
			}
		}
	}
	return accessible_rolls
}

func (g *Grid) ContentValueAt(x int, y int) int {
	content := g.ContentAt(x, y)
	if content == "." {
		return 0
	}
	if content == "@" || content == "x" {
		return 1
	}
	return 0
}

func (g *Grid) CoordinatesAreAccessible(x int, y int) bool {
	sum := 0

	sum += g.ContentValueAt(x-1, y-1) // Top Left
	sum += g.ContentValueAt(x, y-1)   // Top
	sum += g.ContentValueAt(x+1, y-1) // Top Right
	sum += g.ContentValueAt(x-1, y)   // Left
	sum += g.ContentValueAt(x+1, y)   // Right
	sum += g.ContentValueAt(x-1, y+1) // Bottom Left
	sum += g.ContentValueAt(x, y+1)   // Bottom
	sum += g.ContentValueAt(x+1, y+1) // Bottom Right

	if sum < 4 {
		g.MarkAccessibleRollForDeletion(x, y)
		return true
	}

	return false
}

func (g *Grid) MarkAccessibleRollForDeletion(x int, y int) {
	line := g.lines[y]
	new_line := line[:x] + "x" + line[x+1:]

	g.lines[y] = new_line
}

func content() string {
	// filename := "example_input.txt"
	filename := "input.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return ""
	}
	return string(content)
}

func (g *Grid) RemoveRollsMarkedForDeletion() {
	for y := 0; y < g.height; y++ {
		g.lines[y] = strings.Replace(g.lines[y], "x", ".", -1)
	}
}

func (g *Grid) ProcessGrid() int {
	running_total := 0

	fmt.Println(">>>")
	fmt.Println(g.GetAccessibleRollCount())
	for {
		accessible_rolls := g.GetAccessibleRollCount()

		if accessible_rolls == 0 {
			break
		}

		fmt.Println("Accessible rolls:", accessible_rolls)
		running_total += accessible_rolls
		g.RemoveRollsMarkedForDeletion()

		fmt.Println(">>>")
		fmt.Println("Running total:", running_total)
	}

	return running_total
}

func main() {
	grid := NewGrid(content())
	grid.PrintGrid()

	fmt.Println(grid.ProcessGrid())
}
