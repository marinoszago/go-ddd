// Package csv contains the parser struct for csv parsing
package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	internalimport "go-ddd/internal/import"
)

// Parser is a csv file parser.
type Parser struct {
	filePath string
}

// NewParser creates a new file Parser.
func NewParser(filePath string) Parser {
	return Parser{filePath: filePath}
}

// Parse parses a csv file and returns a slice of parsed records.
func (p Parser) Parse() ([]internalimport.ParsedRecord, error) {
	f, err := os.Open(p.filePath)
	if err != nil {
		log.Printf("Unable to read input file "+p.filePath, err)
		return nil, fmt.Errorf("[err] open file: %w", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)

	records, err := csvReader.ReadAll()
	if err != nil {
		log.Printf("Unable to read input file "+p.filePath, err)
		return nil, fmt.Errorf("[err] read file: %w", err)
	}

	if records == nil {
		log.Printf("file is empty: " + p.filePath)
		return nil, fmt.Errorf("[err] file is empty")
	}

	var parsed []internalimport.ParsedRecord

	for _, r := range records[1:] {
		pRecord := internalimport.NewParsedRecord(r)
		parsed = append(parsed, pRecord)
	}

	return parsed, nil
}
