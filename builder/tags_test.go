package builder

import "testing"

func TestQuery_TagsIn(t *testing.T) {
	qb := &QueryBuilder{}

	expects := "FT.SEARCH test @tags:{hello | world}"
	actual := qb.NewQuery("test").TagsIn("tags", "hello", "world").QueryString()

	if actual != expects {
		t.Errorf("tag query does not match. expected '%s' but got '%s'", expects, actual)
		t.Fail()
	}
}

func TestQuery_TagsNotIn(t *testing.T) {
	qb := &QueryBuilder{}

	expects := "FT.SEARCH test @tags:-{hello | world}"
	actual := qb.NewQuery("test").TagsNotIn("tags", "hello", "world").QueryString()

	if actual != expects {
		t.Errorf("tag query does not match. expected '%s' but got '%s'", expects, actual)
		t.Fail()
	}
}
