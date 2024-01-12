package builder

import (
	redisqb "redis-qb"
	"testing"
)

func TestQuery_WithinRadius(t *testing.T) {
	qb := &QueryBuilder{}

	expects := "FT.SEARCH test @loc:[2.01 50.2 5 km]"
	actual := qb.NewQuery("test").WithinRadius(
		"loc",
		redisqb.Coord{Y: 2.01, X: 50.2},
		5,
		redisqb.Kilometres,
	).QueryString()

	if actual != expects {
		t.Errorf("geo within query does not match. expected '%s' but got '%s'", expects, actual)
		t.Fail()
	}
}

func TestQuery_NotWithinRadius(t *testing.T) {
	qb := &QueryBuilder{}

	expects := "FT.SEARCH test @loc:-[2.01 50.2 5 km]"
	actual := qb.NewQuery("test").NotWithinRadius(
		"loc",
		redisqb.Coord{Y: 2.01, X: 50.2},
		5,
		redisqb.Kilometres,
	).QueryString()

	if actual != expects {
		t.Errorf("geo within query does not match. expected '%s' but got '%s'", expects, actual)
		t.Fail()
	}
}
