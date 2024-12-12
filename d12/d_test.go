package d

import (
	"bytes"
	"testing"
)

func TestD(t *testing.T) {
	input := `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run1()

	want := 1930
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestD2(t *testing.T) {
	input := `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run2()

	want := 236
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
