package d

import (
	"bytes"
	"testing"
)

func TestD(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run1()
	want := 18
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

// func TestD2(t *testing.T) {
// 	input := `MMMSXXMASM
// MSAMXMSMSA
// AMXSXMAAMM
// MSAMASMSMX
// XMASAMXAMM
// XXAMMXXAMA
// SMSMSASXSS
// SAXAMASAAA
// MAMMMXMMMM
// MXMXAXMASX`
//
// 	d := NewD()
// 	d.Input(bytes.NewBufferString(input))
// 	got := d.Run2()
// 	want := 9
// 	if got != want {
// 		t.Errorf("got %v want %v", got, want)
// 	}
// }
