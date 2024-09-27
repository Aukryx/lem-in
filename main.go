package main

import (
	"fmt"
	"lem-in/src"
	"os"
)

// main is the entry point of the program.
func main() {
	// Check if a file path is provided as a command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file path")
		return
	}

	filePath := os.Args[1]

	// Parse the input file and create a LemInData struct
	lemInData, err := src.ParseInputFile(filePath)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	// Generate names for all ants
	lemInData.NameAnts()

	// Print parsed data for verification
	fmt.Printf("Number of ants: %d\n", lemInData.NumAnts)
	fmt.Printf("Start room: %s\n", lemInData.StartRoom)
	fmt.Printf("End room: %s\n", lemInData.EndRoom)
	fmt.Printf("Name of ants: %s\n", lemInData.TabAntNames)
	fmt.Println("Rooms:")

	// Find all possible paths from start to end using DFS
	allPaths := src.FindAllPathsBFS(lemInData.Rooms, lemInData.StartRoom, lemInData.EndRoom)

	// Select the optimal paths for ant movement
	BestPath := src.FilterPath(allPaths, lemInData.StartRoom, lemInData.EndRoom)
	fmt.Println("Best paths: ", BestPath)

	// Distribute ants among the selected paths
	antDistribution := src.DistributeAnts(BestPath, lemInData.NumAnts)

	// Print the input data (room information and links)
	for _, room := range lemInData.Rooms {
		if room.IsStart {
			fmt.Println("##start")
		} else if room.IsEnd {
			fmt.Println("##end")
		}
		fmt.Printf("%s %d %d\n", room.Name, room.X, room.Y)
	}
	for name, room := range lemInData.Rooms {
		for _, link := range room.Links {
			fmt.Printf("%s-%s\n", name, link)
		}
	}
	fmt.Println() // Empty line before ant movements

	// Simulate and print ant movements
	src.SimulateAntMovement(BestPath, antDistribution)
}
