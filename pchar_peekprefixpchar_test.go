package rfc3986_test

import (
	"testing"

	"github.com/reiver/go-rfc3986"
)

func TestPeekPrefixPChar(t *testing.T) {

	tests := []struct{
		String string
		ExpectedRune rune
		ExpectedLength int
		ExpectedFound bool
	}{
		{

		},



		{
			String: "",
		},
		{
			String: "^",
		},



		{
			String:       "0",
			ExpectedRune: '0',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "1",
			ExpectedRune: '1',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "2",
			ExpectedRune: '2',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "3",
			ExpectedRune: '3',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "4",
			ExpectedRune: '4',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "5",
			ExpectedRune: '5',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "6",
			ExpectedRune: '6',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "7",
			ExpectedRune: '7',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "8",
			ExpectedRune: '8',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "9",
			ExpectedRune: '9',
			ExpectedLength:1,
			ExpectedFound:true,
		},



		{
			String:       "0123456789",
			ExpectedRune: '0',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "1234567890",
			ExpectedRune: '1',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "2345678901",
			ExpectedRune: '2',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "3456789012",
			ExpectedRune: '3',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "4567890123",
			ExpectedRune: '4',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "5678901234",
			ExpectedRune: '5',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "6789012345",
			ExpectedRune: '6',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "7890123456",
			ExpectedRune: '7',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "8901234567",
			ExpectedRune: '8',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "9012345678",
			ExpectedRune: '9',
			ExpectedLength:1,
			ExpectedFound:true,
		},



		{
			String:       "A",
			ExpectedRune: 'A',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "B",
			ExpectedRune: 'B',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "C",
			ExpectedRune: 'C',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "D",
			ExpectedRune: 'D',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "E",
			ExpectedRune: 'E',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "F",
			ExpectedRune: 'F',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "G",
			ExpectedRune: 'G',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "H",
			ExpectedRune: 'H',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "I",
			ExpectedRune: 'I',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "J",
			ExpectedRune: 'J',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "K",
			ExpectedRune: 'K',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "L",
			ExpectedRune: 'L',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "M",
			ExpectedRune: 'M',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "N",
			ExpectedRune: 'N',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "O",
			ExpectedRune: 'O',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "P",
			ExpectedRune: 'P',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "Q",
			ExpectedRune: 'Q',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "R",
			ExpectedRune: 'R',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "S",
			ExpectedRune: 'S',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "T",
			ExpectedRune: 'T',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "U",
			ExpectedRune: 'U',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "V",
			ExpectedRune: 'V',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "W",
			ExpectedRune: 'W',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "X",
			ExpectedRune: 'X',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "Y",
			ExpectedRune: 'Y',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "Z",
			ExpectedRune: 'Z',
			ExpectedLength:1,
			ExpectedFound:true,
		},



		{
			String:       "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			ExpectedRune: 'A',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "BCDEFGHIJKLMNOPQRSTUVWXYZA",
			ExpectedRune: 'B',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "CDEFGHIJKLMNOPQRSTUVWXYZAB",
			ExpectedRune: 'C',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "DEFGHIJKLMNOPQRSTUVWXYZABC",
			ExpectedRune: 'D',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "EFGHIJKLMNOPQRSTUVWXYZABCD",
			ExpectedRune: 'E',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "FGHIJKLMNOPQRSTUVWXYZABCDE",
			ExpectedRune: 'F',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "GHIJKLMNOPQRSTUVWXYZABCDEF",
			ExpectedRune: 'G',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "HIJKLMNOPQRSTUVWXYZABCDEFG",
			ExpectedRune: 'H',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "IJKLMNOPQRSTUVWXYZABCDEFGH",
			ExpectedRune: 'I',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "JKLMNOPQRSTUVWXYZABCDEFGHI",
			ExpectedRune: 'J',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "KLMNOPQRSTUVWXYZABCDEFGHIJ",
			ExpectedRune: 'K',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "LMNOPQRSTUVWXYZABCDEFGHIJK",
			ExpectedRune: 'L',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "MNOPQRSTUVWXYZABCDEFGHIJKL",
			ExpectedRune: 'M',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "NOPQRSTUVWXYZABCDEFGHIJKLM",
			ExpectedRune: 'N',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "OPQRSTUVWXYZABCDEFGHIJKLMN",
			ExpectedRune: 'O',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "PQRSTUVWXYZABCDEFGHIJKLMNO",
			ExpectedRune: 'P',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "QRSTUVWXYZABCDEFGHIJKLMNOP",
			ExpectedRune: 'Q',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "RSTUVWXYZABCDEFGHIJKLMNOPQ",
			ExpectedRune: 'R',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "STUVWXYZABCDEFGHIJKLMNOPQR",
			ExpectedRune: 'S',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "TUVWXYZABCDEFGHIJKLMNOPQRS",
			ExpectedRune: 'T',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "UVWXYZABCDEFGHIJKLMNOPQRST",
			ExpectedRune: 'U',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "VWXYZABCDEFGHIJKLMNOPQRSTU",
			ExpectedRune: 'V',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "WXYZABCDEFGHIJKLMNOPQRSTUV",
			ExpectedRune: 'W',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "XYZABCDEFGHIJKLMNOPQRSTUVW",
			ExpectedRune: 'X',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "YZABCDEFGHIJKLMNOPQRSTUVWX",
			ExpectedRune: 'Y',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "ZABCDEFGHIJKLMNOPQRSTUVWXY",
			ExpectedRune: 'Z',
			ExpectedLength:1,
			ExpectedFound:true,
		},



		{
			String:       "a",
			ExpectedRune: 'a',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "b",
			ExpectedRune: 'b',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "c",
			ExpectedRune: 'c',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "d",
			ExpectedRune: 'd',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "e",
			ExpectedRune: 'e',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "f",
			ExpectedRune: 'f',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "g",
			ExpectedRune: 'g',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "h",
			ExpectedRune: 'h',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "i",
			ExpectedRune: 'i',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "j",
			ExpectedRune: 'j',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "k",
			ExpectedRune: 'k',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "l",
			ExpectedRune: 'l',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "m",
			ExpectedRune: 'm',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "n",
			ExpectedRune: 'n',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "o",
			ExpectedRune: 'o',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "p",
			ExpectedRune: 'p',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "q",
			ExpectedRune: 'q',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "r",
			ExpectedRune: 'r',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "s",
			ExpectedRune: 's',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "t",
			ExpectedRune: 't',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "u",
			ExpectedRune: 'u',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "v",
			ExpectedRune: 'v',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "w",
			ExpectedRune: 'w',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "x",
			ExpectedRune: 'x',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "y",
			ExpectedRune: 'y',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "z",
			ExpectedRune: 'z',
			ExpectedLength:1,
			ExpectedFound:true,
		},



		{
			String:       "abcdefghijklmnopqrstuvwxyz",
			ExpectedRune: 'a',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "bcdefghijklmnopqrstuvwxyza",
			ExpectedRune: 'b',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "cdefghijklmnopqrstuvwxyzab",
			ExpectedRune: 'c',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "defghijklmnopqrstuvwxyzabc",
			ExpectedRune: 'd',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "efghijklmnopqrstuvwxyzabcd",
			ExpectedRune: 'e',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "fghijklmnopqrstuvwxyzabcde",
			ExpectedRune: 'f',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "ghijklmnopqrstuvwxyzabcdef",
			ExpectedRune: 'g',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "hijklmnopqrstuvwxyzabcdefg",
			ExpectedRune: 'h',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "ijklmnopqrstuvwxyzabcdefgh",
			ExpectedRune: 'i',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "jklmnopqrstuvwxyzabcdefghi",
			ExpectedRune: 'j',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "klmnopqrstuvwxyzabcdefghij",
			ExpectedRune: 'k',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "lmnopqrstuvwxyzabcdefghijk",
			ExpectedRune: 'l',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "mnopqrstuvwxyzabcdefghijkl",
			ExpectedRune: 'm',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "nopqrstuvwxyzabcdefghijklm",
			ExpectedRune: 'n',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "opqrstuvwxyzabcdefghijklmn",
			ExpectedRune: 'o',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "pqrstuvwxyzabcdefghijklmno",
			ExpectedRune: 'p',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "qrstuvwxyzabcdefghijklmnop",
			ExpectedRune: 'q',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "rstuvwxyzabcdefghijklmnopq",
			ExpectedRune: 'r',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "stuvwxyzabcdefghijklmnopqr",
			ExpectedRune: 's',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "tuvwxyzabcdefghijklmnopqrs",
			ExpectedRune: 't',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "uvwxyzabcdefghijklmnopqrst",
			ExpectedRune: 'u',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "vwxyzabcdefghijklmnopqrstu",
			ExpectedRune: 'v',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "wxyzabcdefghijklmnopqrstuv",
			ExpectedRune: 'w',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "xyzabcdefghijklmnopqrstuvw",
			ExpectedRune: 'x',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "yzabcdefghijklmnopqrstuvwx",
			ExpectedRune: 'y',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "zabcdefghijklmnopqrstuvwxy",
			ExpectedRune: 'z',
			ExpectedLength:1,
			ExpectedFound:true,
		},



		{
			String:       "-",
			ExpectedRune: '-',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       ".",
			ExpectedRune: '.',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "_",
			ExpectedRune: '_',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "~",
			ExpectedRune: '~',
			ExpectedLength:1,
			ExpectedFound:true,
		},



		{
			String:       "-._~",
			ExpectedRune: '-',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "._~-",
			ExpectedRune: '.',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "_~-.",
			ExpectedRune: '_',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "~-._",
			ExpectedRune: '~',
			ExpectedLength:1,
			ExpectedFound:true,
		},



		{
			String: "%",
		},
		{
			String: "%A",
		},
		{
			String: "%a",
		},



		{
			String: "%HE",
		},
		{
			String: "%HELLO WORLD",
		},



		{
			String:       "%00",
			ExpectedRune: 0x00,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%01",
			ExpectedRune: 0x01,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%02",
			ExpectedRune: 0x02,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%03",
			ExpectedRune: 0x03,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%04",
			ExpectedRune: 0x04,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%05",
			ExpectedRune: 0x05,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%06",
			ExpectedRune: 0x06,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%07",
			ExpectedRune: 0x07,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%08",
			ExpectedRune: 0x08,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%09",
			ExpectedRune: 0x09,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%0A",
			ExpectedRune: 0x0A,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%0a",
		},
		{
			String:       "%0B",
			ExpectedRune: 0x0B,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%0b",
		},
		{
			String:       "%0C",
			ExpectedRune: 0x0C,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%0c",
		},
		{
			String:       "%0D",
			ExpectedRune: 0x0D,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%0d",
		},
		{
			String:       "%0E",
			ExpectedRune: 0x0E,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%0e",
		},
		{
			String:       "%0F",
			ExpectedRune: 0x0F,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%0f",
		},
		{
			String:       "%10",
			ExpectedRune: 0x10,
			ExpectedLength:3,
			ExpectedFound:true,
		},

		{
			String:       "%FF",
			ExpectedRune: 0xFF,
			ExpectedLength:3,
			ExpectedFound:true,
		},
		{
			String:       "%fF",
		},
		{
			String:       "%Ff",
		},
		{
			String:       "%ff",
		},



		{
			String:       "%DE%AD%C0%DE",
			ExpectedRune: 0xDE,
			ExpectedLength:3,
			ExpectedFound:true,
		},



		{
			String:       "!",
			ExpectedRune: '!',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "$",
			ExpectedRune: '$',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "&",
			ExpectedRune: '&',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "'",
			ExpectedRune: '\'',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "(",
			ExpectedRune: '(',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       ")",
			ExpectedRune: ')',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "*",
			ExpectedRune: '*',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "+",
			ExpectedRune: '+',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       ",",
			ExpectedRune: ',',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       ";",
			ExpectedRune: ';',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "=",
			ExpectedRune: '=',
			ExpectedLength:1,
			ExpectedFound:true,
		},



		{
			String:       "!$&'()*+,;=",
			ExpectedRune: '!',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "$&'()*+,;=!",
			ExpectedRune: '$',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "&'()*+,;=!$",
			ExpectedRune: '&',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "'()*+,;=!$&",
			ExpectedRune: '\'',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "()*+,;=!$&'",
			ExpectedRune: '(',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       ")*+,;=!$&'(",
			ExpectedRune: ')',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "*+,;=!$&'()",
			ExpectedRune: '*',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "+,;=!$&'()*",
			ExpectedRune: '+',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       ",;=!$&'()*+",
			ExpectedRune: ',',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       ";=!$&'()*+,",
			ExpectedRune: ';',
			ExpectedLength:1,
			ExpectedFound:true,
		},
		{
			String:       "=!$&'()*+,;",
			ExpectedRune: '=',
			ExpectedLength:1,
			ExpectedFound:true,
		},
	}

	for testNumber, test := range tests {

		actualRune, actualLength, actualFound := rfc3986.PeekPrefixPChar(test.String)

		{
			expected := test.ExpectedFound
			actual := actualFound

			if expected != actual {
				t.Errorf("For test #%d, the actual 'found' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:  %t", actual)
				t.Logf("STRING: %q", test.String )
				continue
			}
		}

		{
			expected := test.ExpectedLength
			actual := actualLength

			if expected != actual {
				t.Errorf("For test #%d, the actual 'length' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:  %d", actual)
				t.Logf("STRING: %q", test.String )
				continue
			}
		}

		{
			expected := test.ExpectedRune
			actual := actualRune

			if expected != actual {
				t.Errorf("For test #%d, the actual 'rune' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:  %d", actual)
				t.Logf("STRING: %q", test.String )
				continue
			}
		}
	}
}
