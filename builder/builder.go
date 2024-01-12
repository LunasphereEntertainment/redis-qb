package builder

import (
	"github.com/go-redis/redis"
	redisqb "redis-qb"
	"strings"
)

type QueryBuilder struct {
	client *redis.Client
}

func (qb *QueryBuilder) NewQuery(index string) *Query {
	return &Query{
		qb:      qb,
		index:   index,
		filters: make([]redisqb.Filter, 0),
	}
}

type Query struct {
	qb      *QueryBuilder
	index   string
	filters []redisqb.Filter
}

func (q *Query) QueryString() string {
	sb := strings.Builder{}

	sb.WriteString("FT.SEARCH ")
	sb.WriteString(q.index)
	sb.WriteRune(' ')

	for i, q := range q.filters {
		if i > 0 {
			sb.WriteRune(' ')
		}
		sb.WriteString(q.QueryString())
	}

	return sb.String()

}
