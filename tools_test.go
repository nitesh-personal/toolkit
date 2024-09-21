package toolkit

import "testing"

func TestTools_RandomString(t *testing.T) {

	var testTools Tools

	str := testTools.RandomString(10)

	if len(str) != 10 {
		t.Error("wrong len string genrated")
	}

}
