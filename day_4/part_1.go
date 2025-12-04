package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println(content())
	fmt.Println(count_accessible_rolls(content()))
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

func count_accessible_rolls(content string) int {
		total_accessible_rolls := 0
	for line_number, line := range strings.Split(strings.TrimSpace(content), "\n") {
		total_accessible_rolls += count_accessible_rolls_in_line(line, line_number)
	}
	return total_accessible_rolls
}

func count_accessible_rolls_in_line(line_content string, line_number int) int {
	accessible_rolls := 0

	for x_coord, char := range line_content {
		if char == '.' {
			continue
		}
		if char == '@' {
			if roll_is_accessible(x_coord, line_number) {
				accessible_rolls += 1
			}
		}
	}

	return accessible_rolls
}

func roll_is_accessible(x int, y int) bool {
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
		return true
	}

	return false
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

	width := len(strings.Split(content(), "\n")[y])
	if x >= width {
		return ""
	}

	height := len(strings.Split(content(), "\n"))
	if y >= height {
		return ""
	}

	lines := strings.Split(content(), "\n")
	return string(lines[y][x])
}
