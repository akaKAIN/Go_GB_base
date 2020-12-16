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
		num: k2,
		data: data,
	}
	for i := range keys {
		val, err := fn.Get(keys[i])
		if err != nil || val != values[i] {
			t.Fatalf("Trying get val: %d by key: %d, with out put err: %v", values[i], keys[i], err)
		}
	}

	_, err := fn.Get(notExistedKey)
	if err == nil {
		t.Fatalf("Must return error, but returned %v", err)
	}
}

func TestFibNum_Calc(t *testing.T) {
	var (
		k0, k1, k2, k3, k4 = 0, 1, 2, 3, 4
		v0, v1, v2, v3, v4 = 0, 1, 1, 2, 3
		data               = make(map[int]int)
	)
	keys := []int{k0, k1, k2, k3, k4}
	values := []int{v0, v1, v2, v3, v4}
	for i := range keys {
		data[i] = values[i]
	}
	fn := FibNum{
		num: keys[len(keys) - 1],
		data: data,
	}

	// Проверяем на получение ошибки для отрицательных чисел.
	_, err := fn.Calc(-1)
	if err == nil {
		t.Fatalf("Negative numbers must call error")
	}

	// Проверяем состояния в случае, когда число уже в "мапе"
	res, err := fn.Calc(k4)
	if err != nil {
		t.Fatalf("Expected error = nil, but got: %v", err)
	}
	if res != v4 {
		t.Fatalf("Expected %d result, but gotted %d", v4, res)
	}

	// Проверяем состояния в случае, когда число в "мапе" отсутствует
	var key, val = 5, 5
	res, err = fn.Calc(key)
	if fn.num != key {
		t.Fatalf("New key was not added")
	}
	res, err = fn.Get(key)
	if err != nil || res != val {
		t.Fatalf("Expected %d, but gotted %d, err is %v", val, res, err)
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
		num: keys[len(keys) - 1],
		data: data,
	}
	for _, c := range caseList {
		_, isExist := fn.data[c.n]
		if isExist {
			t.Fatalf("Wrong data initiated for struct: %+v", fn)
		}
		fn.calcNext()
		if fn.num != c.n {
			t.Fatalf("Wrong currend num: expected %d, gotted %d", c.n, fn.num)
		}
		val, isExist := fn.data[c.n]
		if !isExist || val != c.expected {
			t.Fatalf("New value: %d must be equal %d", val, c.expected)
		}
	}
}

func BenchmarkCalcFibonacciBine(b *testing.B) {
	caseList := []CaseFibonacci{
		{
			n:        10,
			expected: 0,
		},
		{
			n:        15,
			expected: 0,
		},
		{
			n:        20,
			expected: 0,
		},
	}
	for i := 0; i < b.N; i++ {
		for _, c := range caseList {
			_ = CalcFibonacciBine(c.n)
		}
	}
}

func BenchmarkCalcFibonacciRecursive(b *testing.B) {
	caseList := []CaseFibonacci{
		{
			n:        10,
			expected: 0,
		},
		{
			n:        15,
			expected: 0,
		},
		{
			n:        20,
			expected: 0,
		},
	}
	for i := 0; i < b.N; i++ {
		for _, c := range caseList {
			_ = CalcFibonacciRecursive(c.n)
		}
	}
}

func BenchmarkStructFibonacci(b *testing.B) {
	caseList := []CaseFibonacci{
		{
			n:        10,
			expected: 0,
		},
		{
			n:        15,
			expected: 0,
		},
		{
			n:        20,
			expected: 0,
		},
	}
	fn := New()
	for i := 0; i < b.N; i++ {
		for _, c := range caseList {
			_, err := fn.Calc(c.n)
			if err != nil {
				b.Fatalf("Gotted error: %v", err)
			}
		}
	}
}
