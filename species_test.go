package biorw

import (
	"fmt"
	"strings"
	"testing"
)

func Test_NewSpeciesReaderTSV(t *testing.T) {
	reader := NewSpeciesReader(strings.NewReader(`name	identity	image
Zea mays PH207 v1.1	Zea mays [identity]	D:\\workspace\\zju_orthology\\examples\\species\\multiple\\ZmaPH207.png
Volvox carteri v2.1	Volvox carteri [identity]	D:\\workspace\\zju_orthology\\examples\\species\\multiple\\aaa.jpg
`), '\t')

	for {
		species, err := reader.Read()

		if err != nil {
			break
		}

		fmt.Println()
		for key, value := range species {
			fmt.Println(key, ":", value)
		}
	}
}

func Test_NewSpeciesReaderCSV(t *testing.T) {
	reader := NewSpeciesReader(strings.NewReader(`name,identity,image
Zea mays PH207 v1.1,Zea mays [identity],D:\\workspace\\zju_orthology\\examples\\species\\multiple\\ZmaPH207.png
Volvox carteri v2.1,Volvox carteri [identity],D:\\workspace\\zju_orthology\\examples\\species\\multiple\\aaa.jpg
`), ',')

	for {
		species, err := reader.Read()

		if err != nil {
			break
		}

		fmt.Println()
		for key, value := range species {
			fmt.Println(key, ":", value)
		}
	}
}
