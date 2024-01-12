package builder

import "testing"

func TestGroupQuery(t *testing.T) {
	qb := &QueryBuilder{}

	q := qb.NewQuery("test")

	q.CaptureGroup().Equals("test", "hello world")

	expects := "FT.SEARCH test (@test:(hello world))"
	actual := q.QueryString()

	if expects != actual {
		t.Errorf("search query is not correct.\nexpected: \t%s\nactual: \t\t%s", expects, actual)
	}
}

func TestGroupQuery_Or(t *testing.T) {
	qb := &QueryBuilder{}

	q := qb.NewQuery("test")

	q.CaptureGroup().Equals("test", "hello world")
	q.CaptureGroup(Or).Equals("description", "important document")

	expects := "FT.SEARCH test (@test:(hello world)) | (@description:(important document))"
	actual := q.QueryString()

	if expects != actual {
		t.Errorf("search query is not correct.\nexpected: \t%s\nactual: \t\t%s", expects, actual)
	}
}
