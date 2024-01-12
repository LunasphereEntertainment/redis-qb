package redis_qb

import "testing"

func TestGeoQuery_Default(t *testing.T) {
	query := GeoQuery{
		Field:  "geography",
		Unit:   Miles,
		Centre: Coord{X: 52.2341, Y: -1.052},
		Radius: 10,
	}

	expects := "@geography:[-1.052 52.2341 10 mi]"
	actual := query.QueryString()

	if actual != expects {
		t.Errorf("geo radius query did not match. expected '%s' but got '%s'", expects, actual)
		t.Fail()
	}
}

func TestGeoQuery_Optional(t *testing.T) {
	query := GeoQuery{
		FilterOptions: FilterOptions{Optional: true},
		Field:         "geography",
		Unit:          Miles,
		Centre:        Coord{X: 52.2341, Y: -1.052},
		Radius:        10,
	}

	expects := "@geography:~[-1.052 52.2341 10 mi]"
	actual := query.QueryString()

	if actual != expects {
		t.Errorf("geo radius query did not match. expected '%s' but got '%s'", expects, actual)
		t.Fail()
	}
}

func TestGeoQuery_Inverted(t *testing.T) {
	query := GeoQuery{
		FilterOptions: FilterOptions{Inverted: true},
		Field:         "geography",
		Unit:          Miles,
		Centre:        Coord{X: 52.2341, Y: -1.052},
		Radius:        10,
	}

	expects := "@geography:-[-1.052 52.2341 10 mi]"
	actual := query.QueryString()

	if actual != expects {
		t.Errorf("geo radius query did not match. expected '%s' but got '%s'", expects, actual)
		t.Fail()
	}
}
