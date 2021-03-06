package rcs

import (
	"fmt"
	"testing"
)

func TestFormats(t *testing.T) {
	tests := []struct {
		in  int
		out string
		fn  func(int) string
	}{
		{15, "$f", func(v int) string { return X(v) }},
		{15, "$0f", func(v int) string { return X8(uint8(v)) }},
		{15, "$000f", func(v int) string { return X16(uint16(v)) }},
		{15, "%1111", func(v int) string { return B(v) }},
		{15, "%0000.1111", func(v int) string { return B8(uint8(v)) }},
		{15, "%0000.0000:0000.1111", func(v int) string { return B16(uint16(v)) }},
	}
	for i, test := range tests {
		name := fmt.Sprintf("%v", i)
		t.Run(name, func(t *testing.T) {
			have := test.fn(test.in)
			want := test.out
			if have != want {
				t.Errorf("\n have: %v \n want: %v\n", have, want)
			}
		})
	}
}

func ExampleFromBCD() {
	v := FromBCD(0x42)
	fmt.Println(v)
	// Output: 42
}

func ExampleToBCD() {
	v := ToBCD(42)
	fmt.Printf("%02x", v)
	// Output: 42
}

func TestToBCDOverflow(t *testing.T) {
	want := uint8(0x12)
	have := ToBCD(112)
	if want != have {
		t.Errorf("\n want: %02x \n have: %02x\n", want, have)
	}
}

func TestSliceBits(t *testing.T) {
	b := ParseBits
	tests := []struct {
		lo   int
		hi   int
		in   uint8
		out  uint8
		name string
	}{
		{6, 7, b("11000000"), b("011"), "high one"},
		{6, 7, b("00111111"), b("000"), "high zero"},
		{3, 5, b("00111000"), b("111"), "middle one"},
		{3, 5, b("11000111"), b("000"), "middle zero"},
		{0, 2, b("00000111"), b("111"), "low one"},
		{0, 2, b("11111000"), b("000"), "low zero"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			slice := SliceBits(test.in, test.lo, test.hi)
			if slice != test.out {
				t.Errorf("\n have: %08b \n want: %08b", slice, test.out)
			}
		})
	}
}

func ExampleSliceBits() {
	value := ParseBits("00111000")
	fmt.Printf("%03b", SliceBits(value, 3, 5))
	// Output: 111
}

func TestBitPlane(t *testing.T) {
	b := ParseBits
	tests := []struct {
		offset int
		in     uint8
		out    uint8
	}{
		{0, b("00010001"), b("11")},
		{1, b("00100010"), b("11")},
		{2, b("01000100"), b("11")},
		{3, b("10001000"), b("11")},
	}
	for _, test := range tests {
		out := BitPlane4(test.in, test.offset)
		if out != test.out {
			t.Errorf("\n have: %08b \n want: %08b", out, test.out)
		}
	}
}
