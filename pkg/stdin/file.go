package stdin

import (
	"data-engineer-challenge/pkg/counting"
	"data-engineer-challenge/pkg/input"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// FileReader struct holds the file content and the current line number.
type FileReader struct {
	filecontent []counting.Event
	lineNumber  int
}

// NewFileReader reads content of a file and parses the content into filecontent and returns a new FileReader
func NewFileReader(filename string) *FileReader {
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var filecontent []counting.Event
	err = json.Unmarshal(byteValue, &filecontent)
	if err != nil {
		log.Fatal(err)
	}
	return &FileReader{
		filecontent: filecontent,
		lineNumber:  0,
	}
}

// Read returns the next line from the file and counts the line number.
func (fr *FileReader) Read() (input.Event, error) {
	if fr.lineNumber >= len(fr.filecontent) {
		return input.Event{}, fmt.Errorf("no more lines in file")
	}
	event := fr.filecontent[fr.lineNumber]
	fr.lineNumber++
	return input.Event{
		Ts:  event.Ts,
		Uid: event.Uid,
	}, nil
}
