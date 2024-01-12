package redis_qb

type Query struct {
}

type FilterGroup []Filter

type Filter interface {
	QueryString() string
}

type FilterOptions struct {
	Inverted bool
	Optional bool
}
