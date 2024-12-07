package d

import (
	"bytes"
	"testing"
)

func TestD(t *testing.T) {
	input := `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run1()

	want := 0
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

// func func TestD2(t *testing.T) {
// 	input := `190: 10 19
// 3267: 81 40 27
// 83: 17 5
// 156: 15 6
// 7290: 6 8 6 15
// 161011: 16 10 13
// 192: 17 8 14
// 21037: 9 7 18 13
// 292: 11 6 16 20`
//
// 	d := NewD()
// 	d.Input(bytes.NewBufferString(input))
// 	got := d.Run2()
//
// 	want := 11387
// 	if got != want {
// 		t.Errorf("got %v want %v", got, want)
// 	}
// }
