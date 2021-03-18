package biorw

import (
	"bufio"
	"bytes"
	"io"
)

type Sequence struct {
	Name        string
	Description string
	Residues    string
}

func (s Sequence) String() string {
	return ">" + s.Name + " " + s.Description + "\n" + s.Residues
}

type FastaReader struct {
	reader      *bufio.Reader
	prefix_line []byte
}

func parseHeader(s *Sequence, line []byte) {
	tokens := bytes.Split(line, []byte(" "))
	s.Name = string(tokens[0])
	if len(tokens) > 1 {
		s.Description = string(tokens[1])
	}
}

func (r *FastaReader) Read() (*Sequence, error) {
	var line []byte
	var s = &Sequence{}

	for {
		buff, isPrefix, err := r.reader.ReadLine()
		if err == io.EOF {
			if r.prefix_line != nil {
				parseHeader(s, r.prefix_line)
				r.prefix_line = nil
				return s, nil
			}
			return nil, err
		}

		line = append(line, buff...)
		if isPrefix {
			continue
		}

		line = bytes.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		if bytes.HasPrefix(line, []byte(">")) {
			if r.prefix_line == nil {
				r.prefix_line = line[1:]
				line = nil
			} else {
				parseHeader(s, r.prefix_line)
				r.prefix_line = line[1:]
				line = nil
				return s, nil
			}
		} else {
			s.Residues = s.Residues + string(line)
			line = nil
		}
	}
}

// ReadAll read all
func (r *FastaReader) ReadAll() ([]*Sequence, error) {
	seqs := make([]*Sequence, 0)
	for {
		seq, err := r.Read()
		if err == io.EOF {
			break
		}
		seqs = append(seqs, seq)
	}

	return seqs, nil
}

func NewFastaReader(reader io.Reader) *FastaReader {
	return &FastaReader{
		reader: bufio.NewReader(reader),
	}
}
