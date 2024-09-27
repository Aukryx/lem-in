package src

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// nameAnts generates names for all ants (L1, L2, ..., Ln).
func (l *LemInData) NameAnts() {
	antsNumber := l.NumAnts
	for i := 1; i <= antsNumber; i++ {
		l.TabAntNames = append(l.TabAntNames, "L"+strconv.Itoa(i))
	}
}

// DistributeAnts assigns ants to paths to minimize the number of turns.
func DistributeAnts(paths [][]string, numAnts int) [][]int {
	distribution := make([][]int, len(paths))
	pathLengths := make([]int, len(paths))
	for i, path := range paths {
		pathLengths[i] = len(path) - 1
	}

	// Distribute ants in a specific order
	for i := 1; i <= numAnts; i++ {
		bestPathIndex := 0
		bestArrivalTime := math.MaxInt32
		for j := range paths {
			arrivalTime := len(distribution[j]) + pathLengths[j]
			if arrivalTime < bestArrivalTime {
				bestPathIndex = j
				bestArrivalTime = arrivalTime
			}
		}
		distribution[bestPathIndex] = append(distribution[bestPathIndex], i)
	}
	fmt.Println(distribution)
	return distribution
}

// SimulateAntMovement simulates and prints the movement of ants through the colony.
func SimulateAntMovement(paths [][]string, antDistribution [][]int) {
	type AntPosition struct {
		ant  int
		path int
		step int
	}
	var antPositions []AntPosition
	for pathIndex, ants := range antDistribution {
		for _, ant := range ants {
			antPositions = append(antPositions, AntPosition{ant, pathIndex, 0})
		}
	}
	for len(antPositions) > 0 {
		var moves []string
		var newPositions []AntPosition
		usedLinks := make(map[string]bool)
		for _, pos := range antPositions {
			if pos.step < len(paths[pos.path])-1 {
				currentRoom := paths[pos.path][pos.step]
				nextRoom := paths[pos.path][pos.step+1]
				link := currentRoom + "-" + nextRoom
				if !usedLinks[link] {
					moves = append(moves, fmt.Sprintf("L%d-%s", pos.ant, nextRoom))
					newPositions = append(newPositions, AntPosition{pos.ant, pos.path, pos.step + 1})
					usedLinks[link] = true
				} else {
					newPositions = append(newPositions, pos)
				}
			}
		}
		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}
		antPositions = newPositions
	}
}
