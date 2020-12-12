package fibonacci

import "testing"

func TestNew(t *testing.T) {
	c1 := CaseFibonacci{
		n:        0,
		expected: 0,
	}
	c2 := CaseFibonacci{
		n:        1,
		expected: 1,
	}
	c3 := CaseFibonacci{
		n:        2,
		expected: 1,
	}
	cl := []CaseFibonacci{c1, c2, c3}
	fn := New()
	for i, c := range cl {
		if fn.data[i] != c.expected {
			t.Fatalf("Expected %d, but gotted %d", c.expected, fn.data[i])
		}
	}
}

func TestFibNum_Get(t *testing.T) {
	var (
		k1, k2        = 10, 133
		v1, v2        = 100, 122
		notExistedKey = 42
		data          = make(map[int]int)
	)
	keys := []int{k1, k2}
	values := []int{v1, v2}

	for i := range keys {
		data[keys[i]] = values[i]
	}
	fn := FibNum{
		keys: keys,
		data: data,
	}
	for i := range keys {
		val, err := fn.Get(keys[i])
		if err != nil || val != values[i] {
			t.Fatalf("Trying get val: %d by key: %d, with out put err: %v", values[i], keys[i], err,
			)
		}
	}

	_, err := fn.Get(notExistedKey)
	if err == nil {
		t.Fatalf("Must return error, but returned %v", err)
	}
}

func TestFibNum_calcNext(t *testing.T) {
	var (
		k0, k1, k2, k3, k4 = 0, 1, 2, 3, 4
		v0, v1, v2, v3, v4 = 0, 1, 1, 2, 3
		data               = make(map[int]int)
	)
	keys := []int{k0, k1, k2, k3, k4}
	values := []int{v0, v1, v2, v3, v4}
	caseList := []CaseFibonacci{
		{
			n:        5,
			expected: 5,
		},
		{
			n:        6,
			expected: 8,
		},
		{
			n:        7,
			expected: 13,
		},
	}

	for i := range keys {
		data[keys[i]] = values[i]
	}
	fn := FibNum{
		keys: keys,
		data: data,
	}
	for _, c := range caseList {
		keysLen := len(fn.keys)
		_, isExist := fn.data[c.n]
		if isExist {
			t.Fatalf("Wrong data initiated for struct: %+v", fn)
		}
		fn.calcNext()
		if keysLen != len(fn.keys) -1 {
			t.Fatalf("New key: %d wasn't pushed in keys array", c.n)
		}
		val, isExist := fn.data[c.n]
		if !isExist || val != c.expected {
			t.Fatalf("New value: %d must be equal %d", val, c.expected)
		}
	}
}
