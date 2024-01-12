package builder

import qb "redis-qb"

func (q *Query) TagsIn(field string, tags ...string) *Query {
	q.filters = append(q.filters, &qb.TagSearch{
		Field: field,
		Tags:  tags,
	})

	return q
}

func (q *Query) TagsNotIn(field string, tags ...string) *Query {
	q.filters = append(q.filters, &qb.TagSearch{
		FilterOptions: qb.FilterOptions{Inverted: true},
		Field:         field,
		Tags:          tags,
	})

	return q
}
