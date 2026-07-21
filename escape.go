package jsonparser

const supplementalPlanesOffset = 0x10000
const highSurrogateOffset = 0xD800
const lowSurrogateOffset = 0xDC00

const basicMultilingualPlaneReservedOffset = 0xDFFF
const basicMultilingualPlaneOffset = 0xFFFF

func combineUTF16Surrogates(high, low rune) rune { _ = "STUB: not implemented"; return 0 }

const badHex = -1

func h2I(c byte) int { _ = "STUB: not implemented"; return 0 }

func decodeSingleUnicodeEscape(in []byte) (rune, bool) { _ = "STUB: not implemented"; return 0, false }

func isUTF16EncodedRune(r rune) bool { _ = "STUB: not implemented"; return false }

func decodeUnicodeEscape(in []byte) (rune, int) { _ = "STUB: not implemented"; return 0, 0 }

var backslashCharEscapeTable = [...]byte{
	'"':  '"',
	'\\': '\\',
	'/':  '/',
	'b':  '\b',
	'f':  '\f',
	'n':  '\n',
	'r':  '\r',
	't':  '\t',
}

func unescapeToUTF8(in, out []byte) (inLen int, outLen int) { _ = "STUB: not implemented"; return 0, 0 }

func Unescape(in, out []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
