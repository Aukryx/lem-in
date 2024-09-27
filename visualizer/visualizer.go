package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Room représente une salle dans la colonie de fourmis.
type Room struct {
	Name    string   // Nom de la salle
	X, Y    int      // Coordonnées de la salle
	IsStart bool     // Indique si c'est la salle de départ
	IsEnd   bool     // Indique si c'est la salle d'arrivée
	Links   []string // Noms des salles connectées à cette salle
}

// LemInData contient toutes les informations sur la colonie de fourmis.
type LemInData struct {
	NumAnts     int              // Nombre total de fourmis
	TabAntNames []string         // Noms de toutes les fourmis
	Rooms       map[string]*Room // Map de toutes les salles, indexées par nom
	StartRoom   string           // Nom de la salle de départ
	EndRoom     string           // Nom de la salle d'arrivée
}

// AntPosition représente la position d'une fourmi dans la colonie.
type AntPosition struct {
	ant  int
	path int
	step int
}

// NewLemInData crée et initialise une nouvelle structure LemInData.
func NewLemInData() *LemInData {
	return &LemInData{
		Rooms: make(map[string]*Room),
	}
}

// AddRoom ajoute une nouvelle salle à la structure LemInData.
func (l *LemInData) AddRoom(name string, x, y int) {
	l.Rooms[name] = &Room{
		Name:  name,
		X:     x,
		Y:     y,
		Links: []string{},
	}
}

// SetStartRoom marque une salle comme salle de départ.
func (l *LemInData) SetStartRoom(name string) {
	if room, exists := l.Rooms[name]; exists {
		room.IsStart = true
		l.StartRoom = name
	}
}

// SetEndRoom marque une salle comme salle d'arrivée.
func (l *LemInData) SetEndRoom(name string) {
	if room, exists := l.Rooms[name]; exists {
		room.IsEnd = true
		l.EndRoom = name
	}
}

// AddLink crée un lien bidirectionnel entre deux salles.
func (l *LemInData) AddLink(room1, room2 string) {
	if r1, exists := l.Rooms[room1]; exists {
		r1.Links = append(r1.Links, room2)
	}
	if r2, exists := l.Rooms[room2]; exists {
		r2.Links = append(r2.Links, room1)
	}
}

// main est le point d'entrée du programme.
func main() {
	// Vérifie si un chemin de fichier est fourni en argument
	if len(os.Args) < 2 {
		fmt.Println("Veuillez fournir un chemin de fichier")
		return
	}

	filePath := os.Args[1]

	// Analyse le fichier d'entrée et crée une structure LemInData
	lemInData, err := parseInputFile(filePath)
	if err != nil {
		fmt.Println("Erreur lors de l'analyse du fichier :", err)
		return
	}

	// Génère des noms pour toutes les fourmis
	lemInData.nameAnts()

	// Affiche les données analysées pour vérification
	fmt.Printf("Nombre de fourmis : %d\n", lemInData.NumAnts)
	fmt.Printf("Salle de départ : %s\n", lemInData.StartRoom)
	fmt.Printf("Salle d'arrivée : %s\n", lemInData.EndRoom)
	fmt.Printf("Noms des fourmis : %s\n", lemInData.TabAntNames)
	fmt.Println("Salles :")

	// Trouve tous les chemins possibles du départ à l'arrivée en utilisant BFS
	allPaths := FindAllPathsBFS(lemInData.Rooms, lemInData.StartRoom, lemInData.EndRoom)

	// Sélectionne les meilleurs chemins pour le mouvement des fourmis
	BestPath := FilterPath(allPaths, lemInData.StartRoom, lemInData.EndRoom)
	fmt.Println("Meilleurs chemins : ", BestPath)

	// Distribue les fourmis parmi les chemins sélectionnés
	antDistribution := DistributeAnts(BestPath, lemInData.NumAnts)

	// Affiche les données d'entrée (informations sur les salles et les liens)
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
			// Pour éviter les doublons de liens, n'affichez que si le nom de la salle est inférieur au nom du lien
			if name < link {
				fmt.Printf("%s-%s\n", name, link)
			}
		}
	}
	fmt.Println() // Ligne vide avant les mouvements des fourmis

	// Simule et visualise les mouvements des fourmis
	VisualizeAntMovements(BestPath, antDistribution, lemInData.Rooms)
}

// parseInputFile lit et analyse le fichier d'entrée, créant une structure LemInData.
func parseInputFile(filePath string) (*LemInData, error) {
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
			// Analyse le nombre de fourmis (première ligne du fichier)
			lemInData.NumAnts, err = strconv.Atoi(line)
			if err != nil || lemInData.NumAnts < 1 {
				return nil, fmt.Errorf("nombre de fourmis invalide : %s", line)
			}
			hasAntsNumber = true
			continue
		}

		if line == "##start" {
			nextIsStart = true
		} else if line == "##end" {
			nextIsEnd = true
		} else if strings.Contains(line, " ") && !strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "L") {
			// Définition de salle
			parts := strings.Fields(line)
			if len(parts) != 3 {
				return nil, fmt.Errorf("définition de salle invalide : %s", line)
			}
			name := parts[0]
			x, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, fmt.Errorf("coordonnée x invalide : %s", parts[1])
			}
			y, err := strconv.Atoi(parts[2])
			if err != nil {
				return nil, fmt.Errorf("coordonnée y invalide : %s", parts[2])
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
			// Définition de lien
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				return nil, fmt.Errorf("définition de lien invalide : %s", line)
			}
			if parts[0] == parts[1] {
				return nil, fmt.Errorf("une salle ne peut pas être liée à elle-même : %s", line)
			}
			lemInData.AddLink(parts[0], parts[1])
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if lemInData.StartRoom == "" || lemInData.EndRoom == "" {
		return nil, fmt.Errorf("salle de départ ou d'arrivée non définie")
	}

	return lemInData, nil
}

// nameAnts génère des noms pour toutes les fourmis (L1, L2, ..., Ln).
func (l *LemInData) nameAnts() {
	antsNumber := l.NumAnts
	for i := 1; i <= antsNumber; i++ {
		l.TabAntNames = append(l.TabAntNames, "L"+strconv.Itoa(i))
	}
}

func FindAllPathsBFS(rooms map[string]*Room, start, end string) [][]string {
	var paths [][]string
	queue := list.New()
	queue.PushBack([]string{start})

	for queue.Len() > 0 {
		path := queue.Remove(queue.Front()).([]string)
		lastRoom := path[len(path)-1]

		if lastRoom == end {
			paths = append(paths, path)
			continue
		}

		for _, nextRoom := range rooms[lastRoom].Links {
			if !contains(path, nextRoom) {
				newPath := make([]string, len(path))
				copy(newPath, path)
				newPath = append(newPath, nextRoom)
				queue.PushBack(newPath)
			}
		}
	}
	fmt.Println(paths)
	return paths
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// FilterPath sélectionne les meilleurs chemins pour le mouvement des fourmis.
func FilterPath(AllPaths [][]string, start string, end string) [][]string {
	BestSolution := [][]string{}

	// Parcourt tous les chemins comme points de départ potentiels
	for i := 0; i < len(AllPaths); i++ {
		CurrentSolution := [][]string{AllPaths[i]} // Commence avec le premier chemin

		// Tente de combiner ce chemin avec d'autres
		for j := 0; j < len(AllPaths); j++ {
			if i != j && CheckPath(CurrentSolution, AllPaths[j], start, end) {
				CurrentSolution = append(CurrentSolution, AllPaths[j])
			}
		}

		// Met à jour la meilleure solution si la solution actuelle est meilleure
		if len(CurrentSolution) > len(BestSolution) {
			BestSolution = CurrentSolution
		}
	}

	return BestSolution
}

// CheckPath vérifie si le chemin "current" peut être ajouté à la solution courante "path"
// sans partager de salles autres que départ et arrivée
func CheckPath(path [][]string, current []string, start string, end string) bool {
	// Vérifie chaque chemin déjà dans la solution
	for i := 0; i < len(path); i++ {
		for _, room := range path[i] {
			// Ignore les salles de départ et d'arrivée
			if room == start || room == end {
				continue
			}

			// Si une salle du chemin actuel existe déjà dans le chemin en cours, retourne false
			for _, curRoom := range current {
				if curRoom == room && curRoom != start && curRoom != end {
					return false
				}
			}
		}
	}
	return true
}

// DistributeAnts assigne les fourmis aux chemins pour minimiser le nombre de tours.
func DistributeAnts(paths [][]string, numAnts int) [][]int {
	distribution := make([][]int, len(paths))
	pathLengths := make([]int, len(paths))
	for i, path := range paths {
		pathLengths[i] = len(path) - 1
	}

	// Distribue les fourmis dans un ordre spécifique
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

	return distribution
}

// VisualizeAntMovements simule et génère des fichiers DOT pour visualiser les mouvements des fourmis.
func VisualizeAntMovements(paths [][]string, antDistribution [][]int, rooms map[string]*Room) {
	var antPositions []AntPosition
	for pathIndex, ants := range antDistribution {
		for _, ant := range ants {
			antPositions = append(antPositions, AntPosition{ant, pathIndex, 0})
		}
	}

	occupiedRooms := make(map[string]bool)
	endRoom := paths[0][len(paths[0])-1]

	turn := 0
	for len(antPositions) > 0 {
		var moves []string
		var newPositions []AntPosition

		for _, pos := range antPositions {
			if pos.step < len(paths[pos.path])-1 {
				nextRoom := paths[pos.path][pos.step+1]
				if !occupiedRooms[nextRoom] || nextRoom == endRoom {
					moves = append(moves, fmt.Sprintf("L%d-%s", pos.ant, nextRoom))
					newPositions = append(newPositions, AntPosition{pos.ant, pos.path, pos.step + 1})
					if nextRoom != endRoom {
						occupiedRooms[nextRoom] = true
					}
					if pos.step > 0 {
						delete(occupiedRooms, paths[pos.path][pos.step])
					}
				} else {
					newPositions = append(newPositions, pos)
				}
			}
		}

		antPositions = newPositions

		// Génère le fichier DOT pour le tour actuel
		generateDOTFile(rooms, antPositions, paths, turn)

		fmt.Printf("Tour %d: %s\n", turn+1, strings.Join(moves, " "))
		turn++
	}
}

// generateDOTFile génère un fichier DOT représentant l'état actuel du graphe et des fourmis.
func generateDOTFile(rooms map[string]*Room, antPositions []AntPosition, paths [][]string, turn int) {
	fileName := fmt.Sprintf("step_%d.dot", turn)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier DOT:", err)
		return
	}
	defer file.Close()

	fmt.Fprintln(file, "graph G {")
	fmt.Fprintln(file, "    layout=neato;")
	fmt.Fprintln(file, "    size=\"10,7.5!\";") // Définit la taille en pouces, '!' force la taille exacte
	fmt.Fprintln(file, "    ratio=fill;")       // Remplit l'espace défini
	fmt.Fprintln(file, "    dpi=96;")           // Résolution de l'image
	fmt.Fprintln(file, "    node [shape=circle, style=filled];")
	fmt.Fprintln(file, "    overlap=false;")
	fmt.Fprintln(file, "    splines=true;")
	fmt.Fprintln(file, "    sep=0.1;")
	fmt.Fprintln(file, "    margin=0;")
	fmt.Fprintln(file, "    edge [color=gray];")

	// Map pour les positions des fourmis
	antPositionsMap := make(map[string]string)
	for _, pos := range antPositions {
		currentRoom := paths[pos.path][pos.step]
		antPositionsMap[currentRoom] = fmt.Sprintf("L%d", pos.ant)
	}

	// Définir les nœuds avec les coordonnées
	for _, room := range rooms {
		label := room.Name
		color := "white"
		if ant, ok := antPositionsMap[room.Name]; ok {
			label = fmt.Sprintf("%s (%s)", room.Name, ant)
			color = "lightblue"
		}
		if room.IsStart {
			color = "green"
		} else if room.IsEnd {
			color = "red"
		}
		fmt.Fprintf(file, "    \"%s\" [pos=\"%d,%d!\", label=\"%s\", fillcolor=\"%s\"];\n", room.Name, room.X*100, room.Y*100, label, color)
	}

	// Définir les arêtes (tunnels)
	edgesAdded := make(map[string]bool)
	for _, room := range rooms {
		for _, linkName := range room.Links {
			edgeKey := fmt.Sprintf("%s-%s", room.Name, linkName)
			reverseEdgeKey := fmt.Sprintf("%s-%s", linkName, room.Name)
			if !edgesAdded[edgeKey] && !edgesAdded[reverseEdgeKey] {
				fmt.Fprintf(file, "    \"%s\" -- \"%s\";\n", room.Name, linkName)
				edgesAdded[edgeKey] = true
			}
		}
	}

	fmt.Fprintln(file, "}")
}
