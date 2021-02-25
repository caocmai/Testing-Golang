package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error(("Expected 2 + 2 to equal 4"))
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{10, 12},
		{20, 22},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}

func TestDivision(t *testing.T) {
	tests := []struct {
		x      float64
		y      float64
		result float64
		err    error
	}{
		{x: 1.0, y: 2.0, result: 0.5, err: nil},
		{x: -1.0, y: 2.0, result: -0.5, err: nil},
		// {x: 1.0, y: 0.0, result: +Inf, err: nil},
	}
	for _, test := range tests {
		result := divide(test.x, test.y)
		// assert.IsType(t, test.err, err)
		assert.Equal(t, test.result, result)
	}
}

func BenchmarkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calculate(2)
	}
}

func benchmarkCalculate(input int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Calculate(input)
	}
}

func BenchmarkCalculate100(b *testing.B)         { benchmarkCalculate(100, b) }
func BenchmarkCalculateNegative100(b *testing.B) { benchmarkCalculate(-100, b) }
func BenchmarkCalculateNegative1(b *testing.B)   { benchmarkCalculate(-1, b) }
