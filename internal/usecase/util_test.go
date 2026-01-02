package usecase

import (
	"errors"
	"testing"

	"github.com/kyu08/go-api-server-playground/internal/apperrors"
)

func TestHandleError(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		inputErr        error
		wantSameAsInput bool
		wantInternal    bool
	}{
		"Preconditionエラーの場合はそのまま返す": {
			inputErr:        apperrors.NewPreconditionError("invalid argument"),
			wantSameAsInput: true,
			wantInternal:    false,
		},
		"NotFoundエラーの場合はそのまま返す": {
			inputErr:        apperrors.NewNotFoundError("user"),
			wantSameAsInput: true,
			wantInternal:    false,
		},
		"通常のエラーの場合はInternalErrorでラップして返す": {
			inputErr:        errors.New("database connection failed"),
			wantSameAsInput: false,
			wantInternal:    true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := handleError(tt.inputErr)

			if tt.wantSameAsInput {
				if !errors.Is(got, tt.inputErr) {
					t.Errorf("handleError() = %v, want same as input %v", got, tt.inputErr)
				}
			} else {
				if errors.Is(got, tt.inputErr) {
					t.Errorf("handleError() returned same error, want wrapped error")
				}
			}

			if tt.wantInternal {
				if !apperrors.IsInternal(got) {
					t.Errorf("handleError() returned non-internal error, want internal error")
				}
			}
		})
	}
}
