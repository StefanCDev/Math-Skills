package main

import "testing"

func Test_calculateStandardDeviation(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateStandardDeviation(tt.args.data); got != tt.want {
				t.Errorf("calculateStandardDeviation() = %v, want %v", got, tt.want)
			}
		})
	}
}
