package d

import (
	"bytes"
	"testing"
)

func TestD(t *testing.T) {
	input := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run1()

	want := 36
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestD21(t *testing.T) {
	input := `.....0.
..4321.
..5..2.
..6543.
..7..4.
..8765.
..9....`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run2()

	want := 3
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestD22(t *testing.T) {
	input := `..90..9
...1.98
...2..7
6543456
765.987
876....
987....`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run2()

	want := 13
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestD2(t *testing.T) {
	input := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run2()

	want := 81
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
