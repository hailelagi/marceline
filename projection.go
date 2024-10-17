package main

// Projection
//
//	 modification(r/w/order) over columns, changes the shape of output/attributes
//		πA1,A2,. . . ,An(R).
func Projection(relation Relation, columns []int) Relation {
	result := make([]Row, 0)

	for _, row := range relation.rows {
		newRow := make(Row, len(columns))

		for i, colIdx := range columns {
			newRow[i] = row[colIdx]
		}

		result = append(result, newRow)
	}

	colNames := make([]string, len(columns))

	for i, colIdx := range columns {
		colNames[i] = relation.colNames[colIdx]
	}

	return Relation{
		colNames: colNames,
		rows:     result,
	}
}
