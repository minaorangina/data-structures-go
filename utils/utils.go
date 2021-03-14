package utils

import "testing"

func AssertEqual(t *testing.T, got, want interface{}) {
	t.Helper()

	if got != want {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}

func AssertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
}

func AssertNotNil(t *testing.T, v interface{}) {
	t.Helper()

	if v == nil {
		t.Fatalf("unexpected nil value")
	}
}

func AssertErrored(t *testing.T, err error) {
	t.Helper()

	if err == nil {
		t.Fatalf("expected an error, but didn't get one")
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()

	if got != true {
		t.Errorf("expected to be true, but wasn't")
	}
}
