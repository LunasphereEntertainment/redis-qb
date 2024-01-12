package redis_qb

import "strings"

// TextMatchOptions extends the FilterOptions with an "exact" option.
type TextMatchOptions struct {
	FilterOptions
	Exact bool
}

// TextMatch represents a text query in Redis. Can be configured for exact matching
type TextMatch struct {
	Field   string
	Options TextMatchOptions
	Values  TextList
}

func (t *TextMatch) QueryString() string {
	sb := strings.Builder{}

	sb.WriteRune('@')
	sb.WriteString(t.Field)
	sb.WriteRune(':')
	if t.Options.Inverted {
		sb.WriteRune('-')
	} else if t.Options.Optional {
		sb.WriteRune('~')
	}
	sb.WriteRune('(')
	sb.WriteString(t.Values.String(t.Options.Exact))
	sb.WriteRune(')')

	return sb.String()
}
