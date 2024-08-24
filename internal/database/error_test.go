package database

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsNotFoundError(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		in     error
		expect bool
	}{
		"true": {
			in:     NewNotFoundError("user"),
			expect: true,
		},
		"false": {
			in:     errors.New("err!"),
			expect: false,
		},
		"false(nil)": {
			in:     nil,
			expect: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := IsNotFoundError(tt.in)
			if diff := cmp.Diff(tt.expect, got); diff != "" {
				t.Errorf("mismatch. (-expect +got)\n%s", diff)
			}
		})
	}
}
