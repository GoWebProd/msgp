package msgp

import "math"

// The sizes provided
// are the worst-case
// encoded sizes for
// each type. For variable-
// length types ([]byte, string),
// the total encoded size is
// the prefix size plus the
// length of the object.
const (
	Int8Size       = 2
	Int16Size      = 3
	Int32Size      = 5
	Int64Size      = 9
	ByteSize       = Uint8Size
	Uint8Size      = 2
	Uint16Size     = 3
	Uint32Size     = 5
	Uint64Size     = Int64Size
	Float64Size    = 9
	Float32Size    = 5
	Complex64Size  = 10
	Complex128Size = 18

	TimeSize = 15
	BoolSize = 1
	NilSize  = 1

	MapHeaderSize   = 5
	ArrayHeaderSize = 5

	ExtensionPrefixSize = 6
)

func StringSize(sz int) int {
	return sz + stringPrefixSize(sz)
}

func IntSize(u uint64) int {
	switch {
	case u <= (1<<7)-1:
		return 1
	case u <= math.MaxUint8:
		return 2
	case u <= math.MaxUint16:
		return 3
	case u <= math.MaxUint32:
		return 5
	default:
		return 9
	}
}

func stringPrefixSize(sz int) int {
	switch {
	case sz <= 31:
		return 1
	case sz <= math.MaxUint8:
		return 2
	case sz <= math.MaxUint16:
		return 3
	default:
		return 5
	}
}
