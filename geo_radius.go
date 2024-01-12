package redis_qb

import (
	"strconv"
	"strings"
)

type GeoRadiusUnit string

const (
	Metres     = GeoRadiusUnit("m")
	Kilometres = GeoRadiusUnit("km")
	Miles      = GeoRadiusUnit("mi")
	Feet       = GeoRadiusUnit("ft")
)

// Coord is a 2D coordinate.
type Coord struct {
	X, Y float64
}

// GeoQuery represents a GEORADIUS query in Redis
type GeoQuery struct {
	FilterOptions
	Field  string
	Unit   GeoRadiusUnit
	Centre Coord
	Radius float64
}

func (g *GeoQuery) QueryString() string {
	sb := strings.Builder{}

	sb.WriteRune('@')
	sb.WriteString(g.Field)
	sb.WriteRune(':')
	if g.Optional {
		sb.WriteRune('~')
	} else if g.Inverted {
		sb.WriteRune('-')
	}

	sb.WriteRune('[')
	sb.WriteString(strconv.FormatFloat(g.Centre.Y, 'f', -1, 64))
	sb.WriteRune(' ')
	sb.WriteString(strconv.FormatFloat(g.Centre.X, 'f', -1, 64))
	sb.WriteRune(' ')
	sb.WriteString(strconv.FormatFloat(g.Radius, 'f', -1, 64))
	sb.WriteRune(' ')
	sb.WriteString(string(g.Unit))
	sb.WriteRune(']')

	return sb.String()
}
