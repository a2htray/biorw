package biorw

import (
	"bufio"
	"bytes"
	"io"
)

const (
	DefaultKey1 = "categories"
	DefaultKey2 = "names"
)

type CategoryRow map[string][]string

type CategoryReader struct {
	reader *bufio.Reader
	header bool
	keys   [][]byte
}

func (r *CategoryReader) Read() (CategoryRow, error) {
	var (
		line []byte
	)
	for {
		buff, isPrefix, err := r.reader.ReadLine()
		if err == io.EOF {
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

		if r.keys == nil {
			if r.header {
				r.keys = bytes.Split(line, []byte("\t"))[0:2]
				line = nil
				continue
			} else {
				r.keys = [][]byte{
					[]byte(DefaultKey1),
					[]byte(DefaultKey2),
				}
			}
		}

		tokens := bytes.Split(line, []byte("\t"))

		m := CategoryRow{}

		cs := bytes.Split(tokens[0], []byte("|"))
		m[string(r.keys[0])] = make([]string, len(cs))

		for i, c := range cs {
			m[string(r.keys[0])][i] = string(c)
		}

		cs = bytes.Split(tokens[1], []byte(","))
		m[string(r.keys[1])] = make([]string, len(cs))
		for i, c := range cs {
			m[string(r.keys[1])][i] = string(c)
		}

		line = nil
		return m, nil
	}
}

func NewCategoryReader(reader io.Reader, header bool) *CategoryReader {
	return &CategoryReader{
		reader: bufio.NewReader(reader),
		header: header,
	}
}
