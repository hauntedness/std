package caller

import "testing"

type testCase1 struct {
	namedAssert bool
	name        string
	want        string
}

func assertFunc(tt testCase1, t *testing.T) {
	if got := Name(); got != tt.want {
		t.Errorf("Name() = %v, want %v", got, tt.want)
	}
}

// TestName
func TestName(t *testing.T) {
	tests := []testCase1{
		{
			namedAssert: true,
			name:        "test regular function",
			want:        "github.com/hauntedness/std/caller.assertFunc:12",
		},
		{
			namedAssert: false,
			name:        "test anonymous function",
			want:        "github.com/hauntedness/std/caller.TestName.func1:36",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.namedAssert {
				assertFunc(tt, t)
			} else {
				if got := Name(); got != tt.want {
					t.Errorf("Name() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
