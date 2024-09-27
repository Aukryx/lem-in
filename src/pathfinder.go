package src

import (
	"container/list"
)

func FindAllPathsBFS(rooms map[string]*Room, start, end string) [][]string {
	var paths [][]string
	queue := list.New()
	queue.PushBack([]string{start})

	for queue.Len() > 0 {
		// Print the current state of the queue
		// fmt.Println("Current queue:")
		// for e := queue.Front(); e != nil; e = e.Next() {
		// 	fmt.Printf("%v ", e.Value)
		// }
		// fmt.Println()

		path := queue.Remove(queue.Front()).([]string)
		lastRoom := path[len(path)-1]
		// fmt.Println("Processing room:", lastRoom)

		if lastRoom == end {
			paths = append(paths, path)
			continue
		}
		// Explore other rooms as before
		for _, nextRoom := range rooms[lastRoom].Links {
			if !Contains(path, nextRoom) {
				newPath := make([]string, len(path))
				copy(newPath, path)
				newPath = append(newPath, nextRoom)
				queue.PushBack(newPath)
			}
		}
	}
	return paths
}
func Contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// SelectOptimalPaths chooses the best paths for ant movement.
func FilterPath(AllPaths [][]string, start string, end string) [][]string {
	BestSolution := [][]string{}

	// Parcourir tous les chemins comme point de départ potentiel
	for i := 0; i < len(AllPaths); i++ {
		CurrentSolution := [][]string{AllPaths[i]} // Commence avec le premier chemin

		// Essayer de combiner ce chemin avec d'autres
		for j := 0; j < len(AllPaths); j++ {
			if i != j && CheckPath(CurrentSolution, AllPaths[j], start, end) {
				CurrentSolution = append(CurrentSolution, AllPaths[j])
			}
		}

		// Mettre à jour la meilleure solution si la solution courante est meilleure
		if len(CurrentSolution) > len(BestSolution) {
			BestSolution = CurrentSolution
		}
	}

	return BestSolution
}

// CheckPath vérifie si le chemin "current" peut être ajouté à la solution courante "path"
// sans partager de pièces autres que start et end
func CheckPath(path [][]string, current []string, start string, end string) bool {
	// Vérifier chaque chemin déjà dans la solution
	for i := 0; i < len(path); i++ {
		for _, room := range path[i] {
			// Ignorer les pièces de départ et d'arrivée
			if room == start || room == end {
				continue
			}

			// Si une pièce du chemin actuel existe déjà dans le chemin en cours, on retourne false
			for _, curRoom := range current {
				if curRoom == room && curRoom != start && curRoom != end {
					return false
				}
			}
		}
	}
	return true
}
