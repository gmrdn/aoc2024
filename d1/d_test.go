package d

import (
	"bytes"
	"testing"
)

func TestD(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3
`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run1()
	want := 11
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestDotherSide(t *testing.T) {
	input := `4   3
3   4
5   2
3   1
9   3
3   3
`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run1()
	want := 11
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestD2(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3
`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run2()
	want := 31
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
