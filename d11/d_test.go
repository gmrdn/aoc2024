package d

import (
	"bytes"
	"testing"
)

func TestD(t *testing.T) {
	input := `125 17`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run1()

	want := 55312
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestD2(t *testing.T) {
	input := `125 17`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run2()

	want := 55312
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
