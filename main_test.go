package main

import (
	"bytes"
	"errors"
	"testing"
)

func TestMultipline(t *testing.T) {
	in := bytes.NewBufferString("line")
	expected := "line\nline\nline\n"
	out := bytes.NewBuffer(nil)

	err := multipline(in, out, 3)
	if err != nil {
		t.Fatalf("Unexpected error %s", err.Error())
	}
	result := out.String()
	if result != expected {
		t.Fatalf("Expected %s, got %s", expected, result)
	}
}

type brokenReader struct {
}

func (r brokenReader) Read(p []byte) (int, error) {
	return 0, errors.New("i'm broken")
}

func TestMultiplineError(t *testing.T) {
	in := brokenReader{}
	out := bytes.NewBuffer(nil)

	err := multipline(in, out, 3)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}
