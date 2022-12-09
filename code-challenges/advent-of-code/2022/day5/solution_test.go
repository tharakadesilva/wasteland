package day5

import (
	"testing"
)

func TestPart1(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input string
		want  string
	}{
		"sample input 1": {
			input: "testdata/input1.txt",
			want:  "CMZ",
		},
		"sample input 2": {
			input: "testdata/input2.txt",
			want:  "MQSHJMWNH",
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

func TestPart2(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input string
		want  string
	}{
		"sample input 1": {
			input: "testdata/input1.txt",
			want:  "MCD",
		},
		"sample input 2": {
			input: "testdata/input2.txt",
			want:  "LLWJRBHVZ",
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
