package redis_qb

import "strings"

type PolygonOperator string

const (
	Within   = PolygonOperator("WITHIN")
	Contains = PolygonOperator("CONTAINS")
)

type PolygonQuery struct {
	FilterOptions
	Field    string
	Operator PolygonOperator
	WKT      string
}

func (p *PolygonQuery) QueryString() string {
	sb := strings.Builder{}

	sb.WriteRune('@')
	sb.WriteString(p.Field)
	sb.WriteRune(':')
	if p.Optional {
		sb.WriteRune('~')
	} else if p.Inverted {
		sb.WriteRune('-')
	}

	sb.WriteRune('[')
	sb.WriteString(string(p.Operator))
	sb.WriteString(" $geometry")
	sb.WriteRune(']')
	sb.WriteString(" PARAMS 2 geometry ")
	sb.WriteRune('\'')
	sb.WriteString(p.WKT)
	sb.WriteRune('\'')

	return sb.String()
}
