package builder

import qb "redis-qb"

func (q *Query) WithinRadius(field string, centre qb.Coord, radius float64, unit qb.GeoRadiusUnit) *Query {
	q.filters = append(q.filters, &qb.GeoQuery{
		Field:  field,
		Unit:   unit,
		Centre: centre,
		Radius: radius,
	})

	return q
}

func (q *Query) NotWithinRadius(field string, centre qb.Coord, radius float64, unit qb.GeoRadiusUnit) *Query {
	q.filters = append(q.filters, &qb.GeoQuery{
		FilterOptions: qb.FilterOptions{Inverted: true},
		Field:         field,
		Unit:          unit,
		Centre:        centre,
		Radius:        radius,
	})

	return q
}
