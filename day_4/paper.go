// Read example_input.txt

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println(content())

	total_accessible_rolls := 0
	for line_number, line := range strings.Split(strings.TrimSpace(content()), "\n") {
		total_accessible_rolls += count_accessible_rolls_in_line(line, line_number)
	}

	fmt.Println(total_accessible_rolls)
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

func count_accessible_rolls_in_line(line_content string, line_number int) int {
	accessible_rolls := 0

	for x_coord, char := range line_content {
		if char == '.' {
			continue
		}
		if char == '@' {
			current_coordinate := [2]int{x_coord, line_number}
			accessible_rolls += check_surrounding_content(current_coordinate)
		}
	}

	fmt.Println("Line number:", line_number)
	fmt.Println("Line content:", line_content)
	fmt.Println("Accessible rolls in this line:", accessible_rolls)
	return accessible_rolls
}

func check_surrounding_content(coordinate [2]int) int {
	x := coordinate[0]
	y := coordinate[1]

	sum := 0

	sum += value_at_coordinate(x-1, y-1) // Top Left
	sum += value_at_coordinate(x, y-1) 	// Top
	sum += value_at_coordinate(x+1, y-1) // Top Right
	sum += value_at_coordinate(x-1, y) 	// Left
	sum += value_at_coordinate(x+1, y) 	// Right
	sum += value_at_coordinate(x-1, y+1) // Bottom Left
	sum += value_at_coordinate(x, y+1) 	// Bottom
	sum += value_at_coordinate(x+1, y+1) // Bottom Right

	if sum < 4 {
		return 1
	}

	return 0
}

func value_at_coordinate(x int, y int) int {
	content := content_at_coordinates(x, y)

	if string(content) == "@" {
		return 1
	}

	return 0
}

func content_at_coordinates(x int, y int) string {
	if x < 0 || y < 0 {
		return ""
	}

	if x >= len(strings.Split(content(), "\n")[y]) {
		return ""
	}

	if y >= len(strings.Split(content(), "\n")) {
		return ""
	}

	lines := strings.Split(content(), "\n")
	return string(lines[y][x])
}
