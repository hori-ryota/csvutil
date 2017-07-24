package recenc_test

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/jszwec/recenc"
)

type Bar int

func (b *Bar) UnmarshalField(s string) error {
	n, err := strconv.Atoi(s)
	*b = Bar(n)
	return err
}

type Foo struct {
	Bar Bar `recenc:"bar"`
}

func ExampleDecoder_unmarshaler() {
	csvReader := csv.NewReader(strings.NewReader("10\n5"))

	dec, err := recenc.NewDecoder(csvReader, "bar")
	if err != nil {
		log.Fatal(err)
	}

	var foos []Foo
	for {
		var f Foo
		if err := dec.Decode(&f); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		foos = append(foos, f)
	}

	fmt.Println(foos)

	// Output:
	// [{10} {5}]
}