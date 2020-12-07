package main

func add(numA, numB int) int {
	return numA + numB
}
func sub(numA, numB int) int {
	return numA - numB
}
func mul(numA, numB int) int {
	return numA * numB
}
func div(numA, numB int) int {
	// подстава с int по итогу ... TODO: float, someday
	if numB == 0 {
		return 0
	}
	return numA / numB
}

func getOperationsMap() map[string]func(int, int) int {
	m := make(map[string]func(int, int) int)
	m["+"] = add
	m["-"] = sub
	m["*"] = mul
	m["/"] = div
	return m
}
