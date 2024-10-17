package main

// Product takes in two relations and outputs a relation that contains all possible combinations for tuples
// from the input relations. : (R × S)
func Cross(relOne, relTwo Relation) Relation {
	result := make([]Row, 0)

	// cross product of rows
	for _, row1 := range relOne.rows {
		for _, row2 := range relTwo.rows {
			pair := append(make(Row, 0), row1...)
			product := append(pair, row2...)
			result = append(result, product)
		}
	}

	// cross product of columns
	crossColNames := append(make([]string, 0), relOne.colNames...)

	return Relation{
		colNames: append(crossColNames, relOne.colNames...),
		rows:     result,
	}
}
