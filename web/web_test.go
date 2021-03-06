package web

import (
	"testing"
)

func TestNew_Normal(t *testing.T) {
	_, err := New(&Config{
		HTTPAddr: ":3030",
	})
	if err != nil {
		t.Fatalf("Should have returned a nil error, instead returned {%+v}", err)
	}
}
