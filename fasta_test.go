package biorw

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_ReadOne(t *testing.T) {
	reader := strings.NewReader(`>sequence1 description...
ABABABABABABABABABABABABABABAABABAB
CDCDCDCDCDCDCD
`)

	fastaReader := NewFastaReader(reader)

	for {
		seq, err := fastaReader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(seq)
	}
}

func Test_ReadTwo(t *testing.T) {
	reader := strings.NewReader(`>sequence1 description...
ABABABABABABABABABABABABABABAABABAB
CDCDCDCDCDCDCD1
>sequence2 description...
ABABABABABABABABABABABABABABAABABAB
CDCDCDCDCDCDCD2
`)

	fastaReader := NewFastaReader(reader)

	for {
		seq, err := fastaReader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(seq)
	}
}

func Test_ReadThree(t *testing.T) {
	reader := strings.NewReader(`>sequence1 description1...
ABABABABABABABABABABABABABABAABABAB
CDCDCDCDCDCDCD1
>sequence2 description2...
ABABABABABABABABABABABABABABAABABAB
CDCDCDCDCDCDCD2
>sequence3 description3...
ABABABABABABABABABABABABABABAABABAB
CDCDCDCDCDCDCD3
`)

	fastaReader := NewFastaReader(reader)

	for {
		seq, err := fastaReader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(seq)
	}
}

func Test_ReadAllThree(t *testing.T) {
	reader := strings.NewReader(`>sequence1 description1...
ABABABABABABABABABABABABABABAABABAB
CDCDCDCDCDCDCD1
>sequence2 description2...
ABABABABABABABABABABABABABABAABABAB
CDCDCDCDCDCDCD2
>sequence3 description3...
ABABABABABABABABABABABABABABAABABABCDCDCDCDCDCDCD3
`)

	fastaReader := NewFastaReader(reader)

	seqs, _ := fastaReader.ReadAll()

	for _, seq := range seqs {
		fmt.Println(seq)
	}
}

func Test_NewFastaWriter(t *testing.T) {
	writer := NewFastaWriter(os.Stdout, 5)

	writer.Write(Sequence{
		Name:        "sequence1",
		Description: "description",
		Residues:    "ABABABABABABABABABABABABABABABABABAAB",
	})

	writer = NewFastaWriter(os.Stdout, 10)

	writer.Write(Sequence{
		Name:        "sequence1",
		Description: "description",
		Residues:    "ABABABABABABABABABABABABABABABABABAAB",
	})
}

