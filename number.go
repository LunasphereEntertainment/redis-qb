package redis_qb

import (
	"strconv"
	"strings"
)

type RangeFilter struct {
	FilterOptions
	Field    string
	Min, Max float64
}

func (rf *RangeFilter) QueryString() string {
	sb := strings.Builder{}

	sb.WriteRune('@')
	sb.WriteString(rf.Field)
	sb.WriteRune(':')
	if rf.Inverted {
		sb.WriteRune('-')
	} else if rf.Optional {
		sb.WriteRune('~')
	}

	sb.WriteRune('[')
	sb.WriteString(strconv.FormatFloat(rf.Min, 'g', -1, 64))
	sb.WriteRune(' ')
	sb.WriteString(strconv.FormatFloat(rf.Max, 'g', -1, 64))
	sb.WriteRune(']')

	return sb.String()
}
