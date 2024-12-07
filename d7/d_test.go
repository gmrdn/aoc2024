package d

import (
	"bytes"
	"math/big"
	"testing"
)

func TestDbabystep(t *testing.T) {
	input := `190: 180 10`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run1BigInt()
	want := new(big.Int).SetInt64(190)
	if got.Cmp(want) != 0 {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestDbabystep2(t *testing.T) {
	input := `190: 150 10 30`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run1BigInt()
	want := new(big.Int).SetInt64(190)
	if got.Cmp(want) != 0 {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestDbabystep3(t *testing.T) {
	input := `190: 150 10 25 15`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run1BigInt()
	want := new(big.Int).SetInt64(190)
	if got.Cmp(want) != 0 {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestDbabystep4(t *testing.T) {
	input := `190: 19 10`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run1BigInt()
	want := new(big.Int).SetInt64(190)
	if got.Cmp(want) != 0 {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestDbabystep5(t *testing.T) {
	input := `265: 25 10 15`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run1BigInt()
	want := new(big.Int).SetInt64(265)
	if got.Cmp(want) != 0 {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestDbabystep6(t *testing.T) {
	input := `23000000000001: 2 1000000000000 2 1000000000000 2 1000000000000 2 1000000000000 1`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run1BigInt()
	want := new(big.Int).SetInt64(23000000000001)
	if got.Cmp(want) != 0 {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestDbabystep7(t *testing.T) {
	input := `10000000000000: 1000000 10000000`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run1BigInt()
	want := new(big.Int).SetInt64(10000000000000)
	if got.Cmp(want) != 0 {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestDbabystep8(t *testing.T) {
	input := `390625: 5 5 5 5 5 5 5 5`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run1BigInt()
	want := new(big.Int).SetInt64(390625)
	if got.Cmp(want) != 0 {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestDbabystep9(t *testing.T) {
	input := `3240: 80 40 27`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run1BigInt()
	want := new(big.Int).SetInt64(3240)
	if got.Cmp(want) != 0 {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestDbabystep10(t *testing.T) {
	input := `3240: 81 40 27`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run1BigInt()
	want := new(big.Int).SetInt64(0)
	if got.Cmp(want) != 0 {
		t.Errorf("got %v want %v", got, want)
	}
}

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
	got := d.Run1BigInt()

	want := new(big.Int).SetInt64(3749)
	if got.Cmp(want) != 0 {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestD2debug(t *testing.T) {
	input := `7290: 6 8 6 15`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run2BigInt()

	want := new(big.Int).SetInt64(7290)
	if got.Cmp(want) != 0 {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestD2(t *testing.T) {
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
	got := d.Run2BigInt()

	want := new(big.Int).SetInt64(11387)
	if got.Cmp(want) != 0 {
		t.Errorf("got %v want %v", got, want)
	}
}
