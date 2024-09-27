package src

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// parseInputFile reads and parses the input file, creating a LemInData struct.
func ParseInputFile(filePath string) (*LemInData, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lemInData := NewLemInData()
	nextIsStart := false
	nextIsEnd := false
	hasAntsNumber := false

	for scanner.Scan() {
		line := scanner.Text()

		if !hasAntsNumber {
			// Parse the number of ants (first line of the file)
			lemInData.NumAnts, err = strconv.Atoi(line)
			if err != nil || lemInData.NumAnts < 1 {
				return nil, fmt.Errorf("invalid number of ants: %s", line)
			}
			hasAntsNumber = true
			continue
		}

		if line == "##start" {
			nextIsStart = true
		} else if line == "##end" {
			nextIsEnd = true
		} else if strings.Contains(line, " ") && !strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "L") {
			// Room definition
			parts := strings.Fields(line)
			if len(parts) != 3 {
				return nil, fmt.Errorf("invalid room definition: %s", line)
			}
			name := parts[0]
			x, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, fmt.Errorf("invalid x coordinate: %s", parts[1])
			}
			y, err := strconv.Atoi(parts[2])
			if err != nil {
				return nil, fmt.Errorf("invalid y coordinate: %s", parts[2])
			}
			lemInData.AddRoom(name, x, y)

			if nextIsStart {
				lemInData.SetStartRoom(name)
				nextIsStart = false
			} else if nextIsEnd {
				lemInData.SetEndRoom(name)
				nextIsEnd = false
			}
		} else if strings.Contains(line, "-") && !strings.HasPrefix(line, "#") {
			// Link definition
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid link definition: %s", line)
			}
			if parts[0] == parts[1] {
				return nil, fmt.Errorf("room cannot link to itself: %s", line)
			}
			lemInData.AddLink(parts[0], parts[1])
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if lemInData.StartRoom == "" || lemInData.EndRoom == "" {
		return nil, fmt.Errorf("start or end room not defined")
	}

	return lemInData, nil
}
