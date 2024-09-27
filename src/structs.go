package src

// Room represents a single room in the ant colony.
type Room struct {
	Name    string   // Name of the room
	X, Y    int      // Coordinates of the room
	IsStart bool     // Indicates if this is the start room
	IsEnd   bool     // Indicates if this is the end room
	Links   []string // Names of rooms connected to this rooms
}

// LemInData holds all the information about the ant colony and its configuration.
type LemInData struct {
	NumAnts     int              // Total number of ants
	TabAntNames []string         // Names of all ants
	Rooms       map[string]*Room // Map of all rooms, keyed by room name
	StartRoom   string           // Name of the start room
	EndRoom     string           // Name of the end room
}

// NewLemInData creates and initializes a new LemInData struct.
func NewLemInData() *LemInData {
	return &LemInData{
		Rooms: make(map[string]*Room),
	}
}

// AddRoom adds a new room to the LemInData struct.
func (l *LemInData) AddRoom(name string, x, y int) {
	l.Rooms[name] = &Room{
		Name:  name,
		X:     x,
		Y:     y,
		Links: []string{},
	}
}

// SetStartRoom marks a room as the start room.
func (l *LemInData) SetStartRoom(name string) {
	if room, exists := l.Rooms[name]; exists {
		room.IsStart = true
		l.StartRoom = name
	}
}

// SetEndRoom marks a room as the end room.
func (l *LemInData) SetEndRoom(name string) {
	if room, exists := l.Rooms[name]; exists {
		room.IsEnd = true
		l.EndRoom = name
	}
}

// AddLink creates a bidirectional link between two rooms.
func (l *LemInData) AddLink(room1, room2 string) {
	if r1, exists := l.Rooms[room1]; exists {
		r1.Links = append(r1.Links, room2)
	}
	if r2, exists := l.Rooms[room2]; exists {
		r2.Links = append(r2.Links, room1)
	}
}
