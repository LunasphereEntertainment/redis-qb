package redis_qb

import "testing"

func TestRangeFilter_Default(t *testing.T) {
	query := RangeFilter{
		Field: "age",
		Min:   18,
		Max:   100,
	}

	expects := "@age:[18 100]"
	actual := query.QueryString()

	if expects != actual {
		t.Errorf("range query does not match. expected '%s' but got '%s'", expects, actual)
	}
}

func TestRangeFilter_Optional(t *testing.T) {
	query := RangeFilter{
		FilterOptions: FilterOptions{Optional: true},
		Field:         "age",
		Min:           18,
		Max:           100,
	}

	expects := "@age:~[18 100]"
	actual := query.QueryString()

	if expects != actual {
		t.Errorf("range query does not match. expected '%s' but got '%s'", expects, actual)
	}
}

func TestRangeFilter_Inverted(t *testing.T) {
	query := RangeFilter{
		FilterOptions: FilterOptions{Inverted: true},
		Field:         "age",
		Min:           18,
		Max:           100,
	}

	expects := "@age:-[18 100]"
	actual := query.QueryString()

	if expects != actual {
		t.Errorf("range query does not match. expected '%s' but got '%s'", expects, actual)
	}
}
