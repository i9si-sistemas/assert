package assert

// Length checks if the provided slice or array has the expected length.
// If the length does not match, it reports a fatal error using the provided fatalMessage.
//
// 	mySlice := []int{1, 2, 3}
// 	assert.Length(t, mySlice, 3, "mySlice should have length 3")
func Length[V any](t T, value []V, expected int, args ...any) {
	tester := initTest(t)
	result := len(value)
	configureTest(tester, result, expected)
	if result != expected {
		tester.Fatal(args...)
	}
}
