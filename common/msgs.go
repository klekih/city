package common

// CityInterface is a representation to a city entity
type CityInterface interface {
	SendVector()
	RetrieveLineData()
}

// Report is the base type for reporting status and vectors
// to a city entity
type Report struct {
	CurrentLine [][]float64
}

// WithCurrentLine adds a line to the location
func (r Report) WithCurrentLine(line [][]float64) Report {
	r.CurrentLine = line
	return r
}

// Line is the message send back and forth: from actor to city
// and from the city with information about a line
type Line struct {
	Coordinates [][]float64
}

const (
	// SendReport is a message passed from an actor to the city
	// indicating its status (e.g. location).
	SendReport = iota

	// AskForLine is a message passed from an actor to the city.
	// A response is awaited.
	AskForLine = iota

	// RespondWithLine is a message passed from the city to
	// an actor and it contains line data.
	RespondWithLine = iota
)

// Envelope is the container for different messages sent back
// and forth between an actor and a city
type Envelope struct {
	MessageType int
	Payload     interface{}
}
