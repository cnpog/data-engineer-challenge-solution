package stdout

import (
	"data-engineer-challenge/pkg/output"
	"encoding/json"
	"log"
)

type Console struct{}

func NewConsoleLogger() *Console {
	return &Console{}
}

// Write writes the given outputEvent to the console.
func (c *Console) Write(outputEvent output.OutputEvent) error {
	evt, err := json.Marshal(outputEvent)
	if err != nil {
		return err
	}
	log.Println(string(evt))
	return nil
}
