package day6

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
			want:  7,
		},
		"sample input 2": {
			input: "testdata/input2.txt",
			want:  5,
		},
		"sample input 3": {
			input: "testdata/input3.txt",
			want:  6,
		},
		"sample input 4": {
			input: "testdata/input4.txt",
			want:  10,
		},
		"sample input 5": {
			input: "testdata/input5.txt",
			want:  11,
		},
		"sample input 6": {
			input: "testdata/input6.txt",
			want:  1757,
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
			want:  19,
		},
		"sample input 2": {
			input: "testdata/input2.txt",
			want:  23,
		},
		"sample input 3": {
			input: "testdata/input3.txt",
			want:  23,
		},
		"sample input 4": {
			input: "testdata/input4.txt",
			want:  29,
		},
		"sample input 5": {
			input: "testdata/input5.txt",
			want:  26,
		},
		"sample input 6": {
			input: "testdata/input6.txt",
			want:  2950,
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
