package builder

import "testing"

func TestQuery_InRange(t *testing.T) {
	qb := &QueryBuilder{}

	expects := "FT.SEARCH test @age:[18 100]"
	actual := qb.NewQuery("test").InRange("age", 18, 100).QueryString()

	if actual != expects {
		t.Errorf("range query does not match. expected '%s' but got '%s'", expects, actual)
		t.Fail()
	}
}

func TestQuery_NotInRange(t *testing.T) {
	qb := &QueryBuilder{}

	expects := "FT.SEARCH test @age:-[18 100]"
	actual := qb.NewQuery("test").NotInRange("age", 18, 100).QueryString()

	if actual != expects {
		t.Errorf("range query does not match. expected '%s' but got '%s'", expects, actual)
		t.Fail()
	}
}
