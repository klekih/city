package common

// CityInterface is a representation to a city entity
type CityInterface interface {
	SendVector()
	RetrieveJunction()
}

// Location aggregates the information about the place
// where an actor is at a certain moment in time
type Location struct {
	CurrentLine [][]float64
}

// WithCurrentLine adds a line to the location
func (l Location) WithCurrentLine(line [][]float64) Location {
	l.CurrentLine = line
	return l
}

// Report is the base type for reporting status and vectors
// to a city entity
type Report struct {
	Loc Location
}

// Junction is the message send back from the city with information
// about a junction
type Junction struct {
	Loc Location
}

const (
	// SendReport is a message passed from an actor to the city
	// indicating its status (e.g. location).
	SendReport = iota

	// AskForJunction is a message passed from an actor to the city.
	// A response is awaited.
	AskForJunction = iota

	// RespondWithJunction is a message passed from the city to
	// an actor and it contains junction data.
	RespondWithJunction = iota
)

// Envelope is the container for different messages sent back
// and forth between an actor and a city
type Envelope struct {
	MessageType int
	Payload     interface{}
}
