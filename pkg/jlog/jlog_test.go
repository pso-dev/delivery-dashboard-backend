package jlog

import (
	"os"
	"testing"
)

func TestNewNotNil(t *testing.T) {
	l := New(os.Stdout, LevelInfo)
	if l == nil {
		t.Errorf("new logger should not be nil. expected instance, got mil")
	}
}
