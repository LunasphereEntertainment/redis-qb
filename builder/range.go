package builder

import qb "redis-qb"

func (q *Query) InRange(field string, min, max float64) *Query {
	q.filters = append(q.filters, &qb.RangeFilter{
		Field: field,
		Min:   min,
		Max:   max,
	})

	return q
}

func (q *Query) NotInRange(field string, min, max float64) *Query {
	q.filters = append(q.filters, &qb.RangeFilter{
		FilterOptions: qb.FilterOptions{Inverted: true},
		Field:         field,
		Min:           min,
		Max:           max,
	})

	return q
}
