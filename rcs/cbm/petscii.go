package cbm

var petsciiUnshifted = map[uint8]rune{
	0x20: ' ',
	0x21: '!',
	0x22: '"',
	0x23: '#',
	0x24: '$',
	0x25: '%',
	0x26: '&',
	0x27: '\'',
	0x28: '(',
	0x29: ')',
	0x2a: '*',
	0x2b: '+',
	0x2c: ',',
	0x2d: '-',
	0x2e: '.',
	0x2f: '/',
	0x30: '0',
	0x31: '1',
	0x32: '2',
	0x33: '3',
	0x34: '4',
	0x35: '5',
	0x36: '6',
	0x37: '7',
	0x38: '8',
	0x39: '9',
	0x3a: ':',
	0x3b: ';',
	0x3c: '<',
	0x3d: '=',
	0x3e: '>',
	0x3f: '?',
	0x40: '@',
	0x41: 'A',
	0x42: 'B',
	0x43: 'C',
	0x44: 'D',
	0x45: 'E',
	0x46: 'F',
	0x47: 'G',
	0x48: 'H',
	0x49: 'I',
	0x4a: 'J',
	0x4b: 'K',
	0x4c: 'L',
	0x4d: 'M',
	0x4e: 'N',
	0x4f: 'O',
	0x50: 'P',
	0x51: 'Q',
	0x52: 'R',
	0x53: 'S',
	0x54: 'T',
	0x55: 'U',
	0x56: 'V',
	0x57: 'W',
	0x58: 'X',
	0x59: 'Y',
	0x5a: 'Z',
	0x5b: '[',
	0x5c: '£',
	0x5d: ']',
	0x5e: '↑',
	0x5f: '←',
	0x60: '─',
	0x61: '♠',
	0x62: '│',
	0x63: '─',
	0x64: '�',
	0x65: '�',
	0x66: '�',
	0x67: '�',
	0x68: '�',
	0x69: '╮',
	0x6a: '╰',
	0x6b: '╯',
	0x6c: '�',
	0x6d: '╲',
	0x6e: '╱',
	0x6f: '�',
	0x70: '�',
	0x71: '●',
	0x72: '�',
	0x73: '♥',
	0x74: '�',
	0x75: '╭',
	0x76: '╳',
	0x77: '○',
	0x78: '♣',
	0x79: '�',
	0x7a: '♦',
	0x7b: '┼',
	0x7c: '�',
	0x7d: '│',
	0x7e: 'π',
	0x7f: '◥',
	0xa0: ' ',
	0xa1: '▌',
	0xa2: '▄',
	0xa3: '▔',
	0xa4: '▁',
	0xa5: '▏',
	0xa6: '▒',
	0xa7: '▕',
	0xa8: '�',
	0xa9: '◤',
	0xaa: '�',
	0xab: '├',
	0xac: '▗',
	0xad: '└',
	0xae: '┐',
	0xaf: '▂',
	0xb0: '┌',
	0xb1: '┴',
	0xb2: '┬',
	0xb3: '┤',
	0xb4: '▎',
	0xb5: '▍',
	0xb6: '�',
	0xb7: '�',
	0xb8: '�',
	0xb9: '▃',
	0xba: '�',
	0xbb: '▖',
	0xbc: '▝',
	0xbd: '┘',
	0xbe: '▘',
	0xbf: '▚',
}
var petsciiShifted = map[uint8]rune{
	0x20: ' ',
	0x21: '!',
	0x22: '"',
	0x23: '#',
	0x24: '$',
	0x25: '%',
	0x26: '&',
	0x27: '\'',
	0x28: '(',
	0x29: ')',
	0x2a: '*',
	0x2b: '+',
	0x2c: ',',
	0x2d: '-',
	0x2e: '.',
	0x2f: '/',
	0x30: '0',
	0x31: '1',
	0x32: '2',
	0x33: '3',
	0x34: '4',
	0x35: '5',
	0x36: '6',
	0x37: '7',
	0x38: '8',
	0x39: '9',
	0x3a: ':',
	0x3b: ';',
	0x3c: '<',
	0x3d: '=',
	0x3e: '>',
	0x3f: '?',
	0x40: '@',
	0x41: 'a',
	0x42: 'b',
	0x43: 'c',
	0x44: 'd',
	0x45: 'e',
	0x46: 'f',
	0x47: 'g',
	0x48: 'h',
	0x49: 'i',
	0x4a: 'j',
	0x4b: 'k',
	0x4c: 'l',
	0x4d: 'm',
	0x4e: 'n',
	0x4f: 'o',
	0x50: 'p',
	0x51: 'q',
	0x52: 'r',
	0x53: 's',
	0x54: 't',
	0x55: 'u',
	0x56: 'v',
	0x57: 'w',
	0x58: 'x',
	0x59: 'y',
	0x5a: 'z',
	0x5b: '[',
	0x5c: '£',
	0x5d: ']',
	0x5e: '↑',
	0x5f: '←',
	0x60: '─',
	0x61: 'A',
	0x62: 'B',
	0x63: 'C',
	0x64: 'D',
	0x65: 'E',
	0x66: 'F',
	0x67: 'G',
	0x68: 'H',
	0x69: 'I',
	0x6a: 'J',
	0x6b: 'K',
	0x6c: 'L',
	0x6d: 'M',
	0x6e: 'N',
	0x6f: 'O',
	0x70: 'P',
	0x71: 'Q',
	0x72: 'R',
	0x73: 'S',
	0x74: 'T',
	0x75: 'U',
	0x76: 'V',
	0x77: 'W',
	0x78: 'X',
	0x79: 'Y',
	0x7a: 'Z',
	0x7b: '┼',
	0x7c: '�',
	0x7d: '│',
	0x7e: '▒',
	0x7f: '�',
	0xa0: ' ',
	0xa1: '▌',
	0xa2: '▄',
	0xa3: '▔',
	0xa4: '▁',
	0xa5: '▏',
	0xa6: '▒',
	0xa7: '▕',
	0xa8: '�',
	0xa9: '�',
	0xaa: '�',
	0xab: '├',
	0xac: '▗',
	0xad: '└',
	0xae: '┐',
	0xaf: '▂',
	0xb0: '┌',
	0xb1: '┴',
	0xb2: '┬',
	0xb3: '┤',
	0xb4: '▎',
	0xb5: '▍',
	0xb6: '�',
	0xb7: '�',
	0xb8: '�',
	0xb9: '▃',
	0xba: '✓',
	0xbb: '▖',
	0xbc: '▝',
	0xbd: '┘',
	0xbe: '▘',
	0xbf: '▚',
}

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
