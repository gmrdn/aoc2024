package d

import (
	"bytes"
	"testing"
)

func TestD(t *testing.T) {
	input := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run1()
	want := 161
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestD2(t *testing.T) {
	input := `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run2()
	want := 48
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
