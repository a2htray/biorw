package biorw

import (
	"encoding/csv"
	"io"
)

type Species map[string]string

type SpeciesReader struct {
	reader     *csv.Reader
	headerRead bool
	headers    []string
}

func (r *SpeciesReader) Read() (Species, error) {
	for {
		tokens, err := r.reader.Read()
		if err == io.EOF {
			return Species{}, err
		}

		if !r.headerRead {
			r.headers = tokens
			r.headerRead = true
			continue
		}

		species := Species{}
		for i, header := range r.headers {
			species[header] = tokens[i]
		}

		return species, nil
	}
}

func NewSpeciesReader(reader io.Reader, delimeter rune) *SpeciesReader {
	r := &SpeciesReader{
		reader:     csv.NewReader(reader),
		headerRead: false,
	}

	r.reader.Comma = delimeter

	return r
}
