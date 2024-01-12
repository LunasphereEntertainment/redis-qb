package redis_qb

import "testing"

func TestTagSearch_Default(t *testing.T) {
	query := TagSearch{
		Field: "topics",
		Tags:  []string{"apache kafka", "github", "microsoft"},
	}

	expects := "@topics:{apache kafka | github | microsoft}"
	actual := query.QueryString()

	if actual != expects {
		t.Errorf("tag query does not match. expected '%s' but got '%s'", expects, actual)
		t.Fail()
	}
}

func TestTagSearch_Optional(t *testing.T) {
	query := TagSearch{
		FilterOptions: FilterOptions{Optional: true},
		Field:         "topics",
		Tags:          []string{"apache kafka", "github", "microsoft"},
	}

	expects := "@topics:~{apache kafka | github | microsoft}"
	actual := query.QueryString()

	if actual != expects {
		t.Errorf("tag query does not match. expected '%s' but got '%s'", expects, actual)
		t.Fail()
	}
}

func TestTagSearch_Inverted(t *testing.T) {
	query := TagSearch{
		FilterOptions: FilterOptions{Inverted: true},
		Field:         "topics",
		Tags:          []string{"apache kafka", "github", "microsoft"},
	}

	expects := "@topics:-{apache kafka | github | microsoft}"
	actual := query.QueryString()

	if actual != expects {
		t.Errorf("tag query does not match. expected '%s' but got '%s'", expects, actual)
		t.Fail()
	}
}
