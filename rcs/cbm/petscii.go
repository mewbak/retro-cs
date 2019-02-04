package cbm

import "github.com/blackchip-org/retro-cs/rcs"

// PetsciiDecoder converts byte values to PETSCII equivilents in
// Unicode.
var PetsciiDecoder = func(code uint8) (rune, bool) {
	ch, printable := petsciiUnshifted[code]
	return ch, printable
}

// PetsciiShiftedDecoder converts byte values to PETSCII equivilents in
// Unicode.
var PetsciiShiftedDecoder = func(code uint8) (rune, bool) {
	ch, printable := petsciiShifted[code]
	return ch, printable
}

// http://sta.c64.org/cbm64pettoscr.html

var ScreenDecoder = func(code uint8) (rune, bool) {
	return decoder(code, PetsciiDecoder)
}

var ScreenShiftedDecoder = func(code uint8) (rune, bool) {
	return decoder(code, PetsciiShiftedDecoder)
}

func decoder(code uint8, decode rcs.CharDecoder) (rune, bool) {
	switch {
	case code == 0x5e:
		return decode(0xff)
	case code >= 0x00 && code <= 0x1f:
		return decode(code + 64)
	case code >= 0x20 && code <= 0x3f:
		return decode(code)
	case code >= 0x40 && code <= 0x5f:
		return decode(code + 32)
	case code >= 0x60 && code <= 0x7f:
		return decode(code + 64)
	case code >= 0x80 && code <= 0x9f:
		return decode(code - 128)
	case code >= 0xc0 && code <= 0xdf:
		return decode(code - 64)
	}
	return decode(code)
}
