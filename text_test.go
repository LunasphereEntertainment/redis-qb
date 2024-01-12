package redis_qb

import "testing"

func TestTextMatch_Default(t *testing.T) {
	query := TextMatch{
		Field:  "title",
		Values: []string{"hello", "world"},
	}

	expects := "@title:(hello | world)"
	actual := query.QueryString()

	if actual != expects {
		t.Errorf("text query did not match. expected '%s' but got '%s'", expects, actual)
	}
}

func TestTextMatch_Exact(t *testing.T) {
	query := TextMatch{
		Field:   "title",
		Options: TextMatchOptions{Exact: true},
		Values:  []string{"hello world"},
	}

	expects := "@title:(\"hello world\")"
	actual := query.QueryString()

	if actual != expects {
		t.Errorf("text query did not match. expected '%s' but got '%s'", expects, actual)
	}
}

func TestTextMatch_Optional(t *testing.T) {
	query := TextMatch{
		Field:   "title",
		Options: TextMatchOptions{FilterOptions: FilterOptions{Optional: true}},
		Values:  []string{"hello world"},
	}

	expects := "@title:~(hello world)"
	actual := query.QueryString()

	if actual != expects {
		t.Errorf("text query did not match. expected '%s' but got '%s'", expects, actual)
	}
}
