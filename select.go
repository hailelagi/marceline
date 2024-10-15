package main

// selection table/relation + predicate (expr = true | false | unknown)
// 1. constant
// 2. equality selection

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
