package builder

import (
	redisqb "redis-qb"
)

func (q *Query) Equals(field string, values ...string) *Query {
	q.filters = append(q.filters, &redisqb.TextMatch{
		Field:  field,
		Values: values,
	})

	return q
}

func (q *Query) NotEquals(field string, values ...string) *Query {
	q.filters = append(q.filters, &redisqb.TextMatch{
		Options: redisqb.TextMatchOptions{FilterOptions: redisqb.FilterOptions{Inverted: true}},
		Field:   field,
		Values:  values,
	})

	return q
}

func (q *Query) EqualsExact(field string, values ...string) *Query {
	q.filters = append(q.filters, &redisqb.TextMatch{
		Options: redisqb.TextMatchOptions{Exact: true},
		Field:   field,
		Values:  values,
	})

	return q
}

func (q *Query) NotEqualsExact(field string, values ...string) *Query {
	q.filters = append(q.filters, &redisqb.TextMatch{
		Options: redisqb.TextMatchOptions{Exact: true, FilterOptions: redisqb.FilterOptions{Inverted: true}},
		Field:   field,
		Values:  values,
	})

	return q
}
