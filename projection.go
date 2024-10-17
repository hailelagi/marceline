package main

// Projection: modification(r/w/order) over columns, changes the shape of output/attributes
// π(a1,a2),. . . , (a)n(R).
// SQL: SELECT b_id-100, a_id FROM R WHERE a_id = 'a2'
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
