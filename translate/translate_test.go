package translate

import "testing"

func TestKlingon(t *testing.T) {

	tests := []struct {
		Value      string
		Hex        string
		ShouldPass bool
	}{
		{"Uhura", "0xF8E5 0xF8D6 0xF8E5 0xF8E1 0xF8D0", true},
		{"cg", "", false},
		{"tlh", "0xF8E4", true},
		{"etlh", "0xF8D4 0xF8E4", true},
		{"tlhc", "", false},
		{"tlhgh", "0xF8E4 0xF8D5", true},
		{"ch d", "0xF8D2 0x0020 0xF8D3", true},
		{"ch D", "0xF8D2 0x0020 0xF8D3", true},
		{"q", "0xF8DF", true},
		{"Q", "0xF8E0", true},
		{"QF", "", false},
		{" ", "0x0020", true},
	}

	for i, test := range tests {
		res, err := New(test.Value).Klingon()
		if err != nil && test.ShouldPass {
			t.Error("fail at test case", i+1, test.Value, "should pass the test", "err:", err)
			continue
		}

		if err == nil && !test.ShouldPass {
			t.Error("fail at test case:", i+1, test.Value, "should not pass the test")
			continue
		}

		if test.Hex != res {
			t.Error("fail at test case", i+1, "expected", test.Hex, "got", res)
		}
	}
}
