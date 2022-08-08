package output

type OutputEvent struct {
	Timestamp         int64 `json:"timestamp"`
	DurationInSeconds int   `json:"duration_in_seconds"`
	UniqueUsers       int   `json:"unique_users"`
}
