package redis_qb

import "strings"

type TagSearch struct {
	FilterOptions
	Field string
	Tags  TextList
}

func (t *TagSearch) QueryString() string {
	sb := strings.Builder{}

	sb.WriteRune('@')
	sb.WriteString(t.Field)
	sb.WriteRune(':')
	if t.Optional {
		sb.WriteRune('~')
	} else if t.Inverted {
		sb.WriteRune('-')
	}
	sb.WriteRune('{')
	sb.WriteString(t.Tags.String(false))
	sb.WriteRune('}')

	return sb.String()
}
