package d

import (
	"bytes"
	"testing"
)

func TestD(t *testing.T) {
	input := ``

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run1()
	want := 1
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestD2(t *testing.T) {
	input := ``

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run2()
	want := 0
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
