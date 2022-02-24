package main

import "testing"

func TestFunc(t *testing.T) {
	tt := []struct {
		name           string
		expectedOutput string
	}{
		{"stefan", "hello, stefan"},
	}

	for _, tc := range tt {
		o := sayHello(tc.name)
		if o != tc.expectedOutput {
			t.Fatalf("expected [%s] but got [%s]", tc.expectedOutput, o)
		}
	}
}
