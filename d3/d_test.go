package d

import (
	"bytes"
	"testing"
)

func TestD(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run1()
	want := 1
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestD2(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
1 2 3 4 5`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run2()
	want := 1
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
