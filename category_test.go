package biorw

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func Test_NewCategoryReader(t *testing.T) {
	reader := NewCategoryReader(strings.NewReader(`ortho|OG0000000	AT1G01540,AT1G04520,AT1G05700,AT1G06700,AT1G07550,AT1G07560,AT1G07570,AT1G07870,AT1G09440,AT1G10620,AT1G11050,AT1G14370,AT1G16090,AT1G1611
	ortho|OG0000001	AT1G03510,AT1G03540,AT1G04840,AT1G05750,AT1G06140,AT1G06143,AT1G08070,AT1G09190,AT1G09410,AT1G11290,AT1G13410,AT1G14470,AT1G15510,AT1G1648
	ortho|OG0000002	AT1G07390,AT1G08590,AT1G09970,AT1G12460,AT1G13910,AT1G17230,AT1G17240,AT1G17250,AT1G17750,AT1G25570,AT1G28340,AT1G28440,AT1G33590,AT1G3360
`), false)

	for {
		category, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(category[DefaultKey1], category[DefaultKey2])
	}
}

func Test_NewCategoryReader2(t *testing.T) {
	reader := NewCategoryReader(strings.NewReader(`categories	names
ortho|OG0000000	AT1G01540,AT1G04520,AT1G05700,AT1G06700,AT1G07550,AT1G07560,AT1G07570,AT1G07870,AT1G09440,AT1G10620,AT1G11050,AT1G14370,AT1G16090,AT1G1611
	ortho|OG0000001	AT1G03510,AT1G03540,AT1G04840,AT1G05750,AT1G06140,AT1G06143,AT1G08070,AT1G09190,AT1G09410,AT1G11290,AT1G13410,AT1G14470,AT1G15510,AT1G1648
	ortho|OG0000002	AT1G07390,AT1G08590,AT1G09970,AT1G12460,AT1G13910,AT1G17230,AT1G17240,AT1G17250,AT1G17750,AT1G25570,AT1G28340,AT1G28440,AT1G33590,AT1G3360
`), true)

	for {
		category, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(category[DefaultKey1], category[DefaultKey2])
	}
}