package mirror

import (
)

// CBOR Initial Byte Jump Table - https://www.rfc-editor.org/rfc/rfc8949.html#section-appendix.b

func IsInteger(rawcbor []byte) bool {
	if len(rawcbor) < 1 {
		return false
	}
	m, _, _ := initialMajorJump(rawcbor[0])
	return (m == CBORMajor0NonNegative || m == CBORMajor1Negative)
}

func IsPositive(rawcbor []byte) bool {
	if len(rawcbor) < 1 {
		return false
	}
	m, v, n := initialMajorJump(rawcbor[0])
	// The integer is positive if:
	// - the major type is non-negative, and
	//   - the "next bytes" are encoded (n > 0), or
	//   - the "next bytes" are not used (n == 0) but the value is non-zero (v != 0)
	return (m == CBORMajor0NonNegative && (v != 0 || n > 0))
}

func IsNonNegative(rawcbor []byte) bool {
	if len(rawcbor) < 1 {
		return false
	}
	m, _, _ := initialMajorJump(rawcbor[0])
	return m == CBORMajor0NonNegative
}

func IsNegative(rawcbor []byte) bool {
	if len(rawcbor) < 1 {
		return false
	}
	m, _, _ := initialMajorJump(rawcbor[0])
	return m == CBORMajor1Negative
}

func IsBytes(rawcbor []byte) bool {
	if len(rawcbor) < 1 {
		return false
	}
	m, _, _ := initialMajorJump(rawcbor[0])
	return m == CBORMajor2Bytes
}

func IsText(rawcbor []byte) bool {
	if len(rawcbor) < 1 {
		return false
	}
	m, _, _ := initialMajorJump(rawcbor[0])
	return m == CBORMajor3Text
}

func IsArray(rawcbor []byte) bool {
	if len(rawcbor) < 1 {
		return false
	}
	m, _, _ := initialMajorJump(rawcbor[0])
	return m == CBORMajor4Array
}

func IsMap(rawcbor []byte) bool {
	if len(rawcbor) < 1 {
		return false
	}
	m, _, _ := initialMajorJump(rawcbor[0])
	return m == CBORMajor5Map
}

func IsTag(rawcbor []byte) bool {
	if len(rawcbor) < 1 {
		return false
	}
	m, _, _ := initialMajorJump(rawcbor[0])
	return m == CBORMajor6Tag
}

func IsSimple(rawcbor []byte) bool {
	if len(rawcbor) < 1 {
		return false
	}
	m, _, _ := initialMajorJump(rawcbor[0])
	return m == CBORMajor7Simple
}

func Major(rawcbor []byte) CBORMajor {
	if len(rawcbor) < 1 {
		return CBORMajorUnknown
	}
	m, _, _ := initialMajorJump(rawcbor[0])
	return m
}

func Len(rawcbor []byte) int {
	if len(rawcbor) < 1 {
		return 0
	}
	_, _, n := initialMajorJump(rawcbor[0])
	return n
}

func IsBool(rawcbor []byte) bool {
	if len(rawcbor) < 1 {
		return false
	}

	switch rawcbor[0] {
	// 0xf4 	false
	case 0xf4:
		return true
	// 0xf5 	true
	case 0xf5:
		return true
	default:
		return false
	}
}

func IsTrue(rawcbor []byte) bool {
	if len(rawcbor) < 1 {
		return false
	}

	switch rawcbor[0] {
	// 0xf4 	false
	case 0xf4:
		return false
	// 0xf5 	true
	case 0xf5:
		return true
	default:
		return false
	}
}

func IsFalse(rawcbor []byte) bool {
	if len(rawcbor) < 1 {
		return false
	}

	switch rawcbor[0] {
	// 0xf4 	false
	case 0xf4:
		return true
	// 0xf5 	true
	case 0xf5:
		return false
	default:
		return false
	}
}

func IsNull(rawcbor []byte) bool {
	if len(rawcbor) < 1 {
		return false
	}

	switch rawcbor[0] {
	// 0xf6 	null
	case 0xf6:
		return true
	default:
		return false
	}
}

func IsUndefined(rawcbor []byte) bool {
	if len(rawcbor) < 1 {
		return false
	}

	switch rawcbor[0] {
	// 0xf7 	undefined
	case 0xf7:
		return true
	default:
		return false
	}
}

func IsStop(rawcbor []byte) bool {
	if len(rawcbor) < 1 {
		return false
	}

	switch rawcbor[0] {
	// 0xff 	"break" stop code
	case 0xff:
		return true
	default:
		return false
	}
}

func IsFloat(rawcbor []byte) bool {
	if len(rawcbor) < 1 {
		return false
	}

	switch rawcbor[0] {
	// 0xf9 	half-precision float (two-byte IEEE 754)
	case 0xf9:
		return true
	// 0xfa 	single-precision float (four-byte IEEE 754)
	case 0xfa:
		return true
	// 0xfb 	double-precision float (eight-byte IEEE 754)
	case 0xfb:
		return true
	default:
		return false
	}
}
