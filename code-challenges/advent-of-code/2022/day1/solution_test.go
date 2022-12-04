package day1

import (
	"testing"
)

func TestPart1(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input string
		want  int
	}{
		"sample input 1": {
			input: "testdata/input1.txt",
			want:  24000,
		},
		"sample input 2": {
			input: "testdata/input2.txt",
			want:  70613,
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
		want  int
	}{
		"sample input 1": {
			input: "testdata/input1.txt",
			want:  45000,
		},
		"sample input 2": {
			input: "testdata/input2.txt",
			want:  205805,
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
