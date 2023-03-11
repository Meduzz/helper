package herror

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestCreatingAnError(t *testing.T) {
	candidate := ErrorFromCode(500)

	if candidate == nil {
		t.Error("candidate was nil")
	}

	if !errors.Is(candidate, ErrInternalError) {
		t.Error("candidate was not ErrInternalError")
	}

	if candidate.Error() != "500 Internal Server Error" {
		t.Error("error was not 500 Internal Server Error", candidate.Error())
	}

	code := CodeFromError(candidate)

	if code != 500 {
		t.Error("code was not 500")
	}
}

func TestUnwrapAnError(t *testing.T) {
	wrapped := fmt.Errorf("Something is clearly missing! %w", ErrNotFound)

	if !errors.Is(wrapped, ErrNotFound) {
		t.Error("wrapped error was not ErrNotFound")
	}
}

func TestSomeOtherError(t *testing.T) {
	candidate := os.ErrNotExist

	code := CodeFromError(candidate)

	if code != 500 {
		t.Error("Code was not 500", code)
	}
}
