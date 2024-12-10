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

// func TestD2(t *testing.T) {
// 	input := `89010123
// 78121874
// 87430965
// 96549874
// 45678903
// 32019012
// 01329801
// 10456732`
//
// 	d := NewD()
// 	d.Input(bytes.NewBufferString(input))
// 	got := d.Run2()
//
// 	want := 2858
// 	if got != want {
// 		t.Errorf("got %v want %v", got, want)
// 	}
// }
