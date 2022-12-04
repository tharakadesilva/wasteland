package day2

import (
	"testing"
)

func TestPar1(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input string
		want  int
	}{
		"sample input 1": {
			input: "inputs/input1.txt",
			want:  15,
		},
		"sample input 2": {
			input: "inputs/input2.txt",
			want:  14827,
		},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := part1(tc.input)
			if err != nil {
				t.Fatalf("failed to run: %v", err)
			}

			if got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestPar2(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input string
		want  int
	}{
		"sample input 1": {
			input: "inputs/input1.txt",
			want:  12,
		},
		"sample input 2": {
			input: "inputs/input2.txt",
			want:  13889,
		},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := part2(tc.input)
			if err != nil {
				t.Fatalf("failed to run: %v", err)
			}

			if got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
		})
	}
}
