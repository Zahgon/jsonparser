package jsonparser

import (
	"errors"
)

var (
	KeyPathNotFoundError       = errors.New("Key path not found")
	UnknownValueTypeError      = errors.New("Unknown value type")
	MalformedJsonError         = errors.New("Malformed JSON error")
	MalformedStringError       = errors.New("Value is string, but can't find closing '\"' symbol")
	MalformedArrayError        = errors.New("Value is array, but can't find closing ']' symbol")
	MalformedObjectError       = errors.New("Value looks like object, but can't find closing '}' symbol")
	MalformedValueError        = errors.New("Value looks like Number/Boolean/None, but can't find its end: ',' or '}' symbol")
	OverflowIntegerError       = errors.New("Value is number, but overflowed while parsing")
	MalformedStringEscapeError = errors.New("Encountered an invalid escape sequence in a string")
	NullValueError             = errors.New("Value is null")
)

const unescapeStackBufSize = 64

func tokenEnd(data []byte) int { _ = "STUB: not implemented"; return 0 }

func findTokenStart(data []byte, token byte) int { _ = "STUB: not implemented"; return 0 }

func findKeyStart(data []byte, key string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func tokenStart(data []byte) int { _ = "STUB: not implemented"; return 0 }

func nextToken(data []byte) int { _ = "STUB: not implemented"; return 0 }

func lastToken(data []byte) int { _ = "STUB: not implemented"; return 0 }

func stringEnd(data []byte) (int, bool) { _ = "STUB: not implemented"; return 0, false }

func blockEnd(data []byte, openSym byte, closeSym byte) int { _ = "STUB: not implemented"; return 0 }

func searchKeys(data []byte, keys ...string) int { _ = "STUB: not implemented"; return 0 }

func sameTree(p1, p2 []string) bool { _ = "STUB: not implemented"; return false }

const stackArraySize = 128

func EachKey(data []byte, cb func(int, []byte, ValueType, error), paths ...[]string) int {
	_ = "STUB: not implemented"
	return 0
}

type ValueType int

const (
	NotExist = ValueType(iota)
	String
	Number
	Object
	Array
	Boolean
	Null
	Unknown
)

func (vt ValueType) String() string { _ = "STUB: not implemented"; return "" }

var (
	trueLiteral  = []byte("true")
	falseLiteral = []byte("false")
	nullLiteral  = []byte("null")
)

func createInsertComponent(keys []string, setValue []byte, comma, object bool) []byte {
	_ = "STUB: not implemented"
	return nil
}

func calcAllocateSpace(keys []string, setValue []byte, comma, object bool) int {
	_ = "STUB: not implemented"
	return 0
}

func WriteToBuffer(buffer []byte, str string) int { _ = "STUB: not implemented"; return 0 }

func Delete(data []byte, keys ...string) []byte { _ = "STUB: not implemented"; return nil }

func Set(data []byte, setValue []byte, keys ...string) (value []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getType(data []byte, offset int) ([]byte, ValueType, int, error) {
	_ = "STUB: not implemented"
	return nil, *new(ValueType), 0, nil
}

func Get(data []byte, keys ...string) (value []byte, dataType ValueType, offset int, err error) {
	_ = "STUB: not implemented"
	return nil, *new(ValueType), 0, nil
}

func internalGet(data []byte, keys ...string) (value []byte, dataType ValueType, offset, endOffset int, err error) {
	_ = "STUB: not implemented"
	return nil, *new(ValueType), 0, 0, nil
}

func ArrayEach(data []byte, cb func(value []byte, dataType ValueType, offset int, err error), keys ...string) (offset int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func ObjectEach(data []byte, callback func(key []byte, value []byte, dataType ValueType, offset int) error, keys ...string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func GetUnsafeString(data []byte, keys ...string) (val string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func GetString(data []byte, keys ...string) (val string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func GetFloat(data []byte, keys ...string) (val float64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func GetInt(data []byte, keys ...string) (val int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func GetBoolean(data []byte, keys ...string) (val bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func ParseBoolean(b []byte) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func ParseString(b []byte) (string, error) { _ = "STUB: not implemented"; return "", nil }

func ParseFloat(b []byte) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func ParseInt(b []byte) (int64, error) { _ = "STUB: not implemented"; return 0, nil }
