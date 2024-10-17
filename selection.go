package main

// selection table/relation + predicate (expr = true | false | unknown)
// σ predicate(R).
// 1. constant
// 2. equality selection
// SQL: SELECT * FROM R WHERE a_id = 'a2'

func ConstantSelect(relation Relation, idx int, expr string) Relation {
	result := make([]Row, 0)

	for _, r := range relation.rows {
		if r[idx] == expr {
			result = append(result, r)
		}
	}

	return Relation{
		colNames: relation.colNames,
		rows:     result,
	}
}

func EqualSelect(relation Relation, x, y int) Relation {
	result := make([]Row, 0)

	for _, row := range relation.rows {
		if row[x] == row[y] {
			result = append(result, row)
		}
	}

	return Relation{
		colNames: relation.colNames,
		rows:     result,
	}
}
