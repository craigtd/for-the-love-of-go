package calculator_test

import (
	"calculator"
	"testing"
)

type testCase struct {
    name string
    a, b float64
    want float64
}

func TestAdd(t *testing.T) {
    t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
		{name: "Two negative numbers that sum to a negative", a: -1, b: -2, want: -3},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s: Add(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel() // run this test concurrently with other tests
    testCases := []testCase{
        {a: 4, b: 2, want: 2},
        {a: 6, b: 1, want: 5},
        {a: -1, b: -1, want: 0},
    }
    for _, tc := range testCases {
        got := calculator.Subtract(tc.a, tc.b)
        if tc.want != got {
            t.Errorf("Subtract(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
        }
    }
}

func TestMultiply(t *testing.T) {
	t.Parallel()
    testCases := []testCase{
        {a: 4, b: 9, want: 36},
        {a: 4, b: 0, want: 0},
        {name: "Two negative numbers with a positive product", a: -2, b: -3, want: 6},
    }
    for _, tc := range testCases {
        got := calculator.Multiply(tc.a, tc.b)
        if tc.want != got {
            t.Errorf("%s: Multiply(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
        }
    }
}

func TestDivide(t *testing.T) {
	t.Parallel()
	var want float64 = 3
	got := calculator.Divide(12, 4)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
