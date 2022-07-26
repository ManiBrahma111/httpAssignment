package unitTesting

import "testing"

type s struct {
	a   int
	b   int
	res int
	exp int
}

func TestUnit(t *testing.T) {
	var arr = []s{
		{a: 10, b: 5, res: 2, exp: div(10, 5)},
		{a: 0, b: 5, res: 0, exp: div(0, 5)},
		{a: 10, b: 0, res: 0, exp: div(10, 0)},
	}
	for _, val := range arr {
		if val.res != val.exp {
			t.Errorf("your output is wrong")
		}
	}
}
