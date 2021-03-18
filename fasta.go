package biorw

import (
	"bufio"
	"bytes"
	"io"
)

type Sequencer interface {
	GetName() string
	GetDescription() string
	GetResidues() string
	GetLen() int
}

type Sequence struct {
	Name        string
	Description string
	Residues    string
}

func (s Sequence) GetName() string {
	return s.Name
}

func (s Sequence) GetDescription() string {
	return s.Description
}

func (s Sequence) GetResidues() string {
	return s.Residues
}

func (s Sequence) GetLen() int {
	return len(s.Residues)
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

type FastaWriter struct {
	writer io.Writer
	width  int
}

func (w *FastaWriter) Write(s Sequencer) (n int, err error) {
	var _n int
	name, description := s.GetName(), s.GetDescription()

	n, err = w.writer.Write([]byte(">" + name + " " + description + "\n"))
	if err != nil {
		return
	}

	segNum := s.GetLen() / w.width
	residuesBytes := []byte(s.GetResidues())

	for i := 0; i < segNum+1; i++ {

		src := residuesBytes[i*w.width : (i+1)*w.width]
		out := make([]byte, len(src))
		copy(out, src)
		out = append(out, '\n')
		_n, err = w.writer.Write(out)
		n = n + _n
		if err != nil {
			return
		}
	}
	return
}

func NewFastaWriter(writer io.Writer, width int) *FastaWriter {
	return &FastaWriter{
		writer: writer,
		width:  width,
	}
}
