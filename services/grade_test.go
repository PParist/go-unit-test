package services_test

import (
	"fmt"
	"testing"

	"unittest/services"

	"github.com/stretchr/testify/assert"
)

func TestCheckGrade(t *testing.T) {
	type testcase struct {
		name     string
		score    int
		expected string
	}

	testcases := []testcase{
		{name: "A", score: 90, expected: "A"},
		{name: "B", score: 80, expected: "B"},
		{name: "C", score: 70, expected: "C"},
		{name: "D", score: 60, expected: "D"},
		{name: "F", score: 50, expected: "F"},
	}

	for _, tc := range testcases {

		t.Run(tc.name, func(t *testing.T) {
			grade := services.CheckGrade(tc.score)
			assert.Equal(t, tc.expected, grade)
		})
	}
}

// benchmark is used to measure the performance of the code
func BenchmarkCheckGrade(b *testing.B) {

	for i := 0; i < b.N; i++ {
		services.CheckGrade(90)
	}
}

func ExampleCheckGrade() {
	grade := services.CheckGrade(90)
	fmt.Println(grade)
	// Output: A
}
