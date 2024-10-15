package main

import (
	"bytes"
	"fmt"
)

type Row []string

type Relation struct {
	colNames []string
	rows     []Row
}

func (r Relation) String() string {
	var buf bytes.Buffer

	for i, col := range r.colNames {
		if i > 0 {
			buf.WriteString("| ")
		}

		fmt.Fprintf(&buf, "%-13s", col)
	}

	buf.WriteByte('\n')

	for range r.colNames {
		buf.WriteString("------------")
	}

	buf.WriteByte('\n')
	for _, row := range r.rows {
		for _, col := range row {
			fmt.Fprintf(&buf, "%-15s", col)
		}
		buf.WriteByte('\n')
	}

	return buf.String()
}

func main() {
	r := Relation{
		colNames: []string{"hi", "hello", "world"},
		rows: []Row{
			{"go", "is", "fast"},
			{"rust", "is", "slow"},
			{"c", "has", "runtime"},
		},
	}

	fmt.Println(r)
}
