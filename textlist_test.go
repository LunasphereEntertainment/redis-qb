package redis_qb

import "testing"

func TestTextList_String(t *testing.T) {
	list := make(TextList, 3)
	list[0] = "hello"
	list[1] = "world"
	list[2] = "test"

	expects := "hello | world | test"
	actual := list.String(false)

	if expects != actual {
		t.Errorf("text list output did not match. expected '%s' but got '%s'", expects, actual)
		t.Fail()
	}

	expects = "\"hello\" | \"world\" | \"test\""
	actual = list.String(true)

	if expects != actual {
		t.Errorf("text list output did not match. expected '%s' but got '%s'", expects, actual)
		t.Fail()
	}
}
