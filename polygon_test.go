package redis_qb

import "testing"

func TestPolygonQuery_Default(t *testing.T) {
	query := PolygonQuery{
		Field:    "geom",
		Operator: Within,
		WKT:      "POLYGON((0 0, 0 100, 100 100, 100 0, 0 0))",
	}

	expects := "@geom:[WITHIN $geometry] PARAMS 2 geometry 'POLYGON((0 0, 0 100, 100 100, 100 0, 0 0))'"
	actual := query.QueryString()

	if expects != actual {
		t.Errorf("polygon WITHIN query does not match.\nexpected \t'%s'\ngot \t\t'%s'", expects, actual)
		t.Fail()
	}
}
