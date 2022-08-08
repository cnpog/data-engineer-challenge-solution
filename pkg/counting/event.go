package counting

// Event is the input event
type Event struct {
	Ts  int64
	Uid string
}

// OutgoingEvent is the output of the counting service.
type OutgoingEvent struct {
	Timestamp         int64
	DurationInSeconds int
	UniqueUsers       int
}
