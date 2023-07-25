package main

import "testing"

func TestID[T ID](t *testing.T, id T) {
	testCases := []struct {
		name string
		id   T
	}{
		{"ID", id},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if !tc.id.IsUnique() {
				t.Error("ID is not unique")
			}
		})
	}
}
