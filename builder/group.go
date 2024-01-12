package builder

import (
	"strings"
)

type GroupOperator string

const (
	And = GroupOperator("")
	Or  = GroupOperator("| ")
)

type GroupQuery struct {
	*Query
	operator GroupOperator
	//parent *Query
	//parent  *Query
	//filters []redisqb.Filter
}

func (group *GroupQuery) QueryString() string {
	sb := strings.Builder{}

	sb.WriteString(string(group.operator))

	sb.WriteRune('(')
	for i, filter := range group.filters {
		if i > 0 {
			sb.WriteRune(' ')
		}
		sb.WriteString(filter.QueryString())
	}
	sb.WriteRune(')')

	return sb.String()
}

func (q *Query) CaptureGroup(operator ...GroupOperator) *Query {
	group := &GroupQuery{
		Query: q.qb.NewQuery(""),
	}

	if len(operator) > 0 {
		group.operator = operator[0]
	} else {
		group.operator = And
	}

	q.filters = append(q.filters, group)

	return group.Query
}
