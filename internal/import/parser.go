// Package internalimport contains the code for import
package internalimport

// Parser is the interface for a parser.
//
//go:generate mockgen -destination ./mock/parser.go -package mock . Parser
type Parser interface {
	Parse() ([]ParsedRecord, error)
}

// ParsedRecord contains a record of the file.
type ParsedRecord struct {
	Record []string
}

// NewParsedRecord creates a new file record as ParsedRecord.
func NewParsedRecord(record []string) ParsedRecord {
	return ParsedRecord{
		Record: record,
	}
}

// ContainsEmptyFields checks if a parsed record contains empty fields.
func (pr ParsedRecord) ContainsEmptyFields() bool {
	if pr.Record == nil {
		return true
	}

	for _, field := range pr.Record {
		if field == "" {
			return true
		}
	}

	return false
}
