package builder

import "testing"

func TestQuery_Equals(t *testing.T) {
	qb := QueryBuilder{}

	expects := "FT.SEARCH test @title:(hello world) @description:-(boringness)"

	actual := qb.
		NewQuery("test").
		Equals("title", "hello world").
		NotEquals("description", "boringness").
		QueryString()

	if actual != expects {
		t.Errorf("query string did not match. expected '%s' but got '%s'", expects, actual)
		t.Fail()
	}
}
