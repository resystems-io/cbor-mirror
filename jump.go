package mirror

func initialMajorJump(b byte) (major CBORMajor, value int, next int) {

	// https://www.rfc-editor.org/rfc/rfc8949.html#section-appendix.b

	switch b {
	// 0x00..0x17 	unsigned integer 0x00..0x17 (0..23)
	case 0x00:
		return CBORMajor0NonNegative, 0, 0
	case 0x01:
		return CBORMajor0NonNegative, 1, 0
	case 0x02:
		return CBORMajor0NonNegative, 2, 0
	case 0x03:
		return CBORMajor0NonNegative, 3, 0
	case 0x04:
		return CBORMajor0NonNegative, 4, 0
	case 0x05:
		return CBORMajor0NonNegative, 5, 0
	case 0x06:
		return CBORMajor0NonNegative, 6, 0
	case 0x07:
		return CBORMajor0NonNegative, 7, 0
	case 0x08:
		return CBORMajor0NonNegative, 8, 0
	case 0x09:
		return CBORMajor0NonNegative, 9, 0
	case 0x0a:
		return CBORMajor0NonNegative, 10, 0
	case 0x0b:
		return CBORMajor0NonNegative, 11, 0
	case 0x0c:
		return CBORMajor0NonNegative, 12, 0
	case 0x0d:
		return CBORMajor0NonNegative, 13, 0
	case 0x0e:
		return CBORMajor0NonNegative, 14, 0
	case 0x0f:
		return CBORMajor0NonNegative, 15, 0
	case 0x10:
		return CBORMajor0NonNegative, 16, 0
	case 0x11:
		return CBORMajor0NonNegative, 17, 0
	case 0x12:
		return CBORMajor0NonNegative, 18, 0
	case 0x13:
		return CBORMajor0NonNegative, 19, 0
	case 0x14:
		return CBORMajor0NonNegative, 20, 0
	case 0x15:
		return CBORMajor0NonNegative, 21, 0
	case 0x16:
		return CBORMajor0NonNegative, 22, 0
	case 0x17:
		return CBORMajor0NonNegative, 23, 0
	// 0x18 	unsigned integer (one-byte uint8_t follows)
	case 0x18:
		return CBORMajor0NonNegative, 0, 1
	// 0x19 	unsigned integer (two-byte uint16_t follows)
	case 0x19:
		return CBORMajor0NonNegative, 0, 2
	// 0x1a 	unsigned integer (four-byte uint32_t follows)
	case 0x1a:
		return CBORMajor0NonNegative, 0, 4
	// 0x1b 	unsigned integer (eight-byte uint64_t follows)
	case 0x1b:
		return CBORMajor0NonNegative, 0, 8
	// 0x20..0x37 	negative integer -1-0x00..-1-0x17 (-1..-24)
	case 0x20:
		return CBORMajor1Negative, -1, 0
	case 0x21:
		return CBORMajor1Negative, -2, 0
	case 0x22:
		return CBORMajor1Negative, -3, 0
	case 0x23:
		return CBORMajor1Negative, -4, 0
	case 0x24:
		return CBORMajor1Negative, -5, 0
	case 0x25:
		return CBORMajor1Negative, -6, 0
	case 0x26:
		return CBORMajor1Negative, -7, 0
	case 0x27:
		return CBORMajor1Negative, -8, 0
	case 0x28:
		return CBORMajor1Negative, -9, 0
	case 0x29:
		return CBORMajor1Negative, -10, 0
	case 0x2a:
		return CBORMajor1Negative, -11, 0
	case 0x2b:
		return CBORMajor1Negative, -12, 0
	case 0x2c:
		return CBORMajor1Negative, -13, 0
	case 0x2d:
		return CBORMajor1Negative, -14, 0
	case 0x2e:
		return CBORMajor1Negative, -15, 0
	case 0x2f:
		return CBORMajor1Negative, -16, 0
	case 0x30:
		return CBORMajor1Negative, -17, 0
	case 0x31:
		return CBORMajor1Negative, -18, 0
	case 0x32:
		return CBORMajor1Negative, -19, 0
	case 0x33:
		return CBORMajor1Negative, -20, 0
	case 0x34:
		return CBORMajor1Negative, -21, 0
	case 0x35:
		return CBORMajor1Negative, -22, 0
	case 0x36:
		return CBORMajor1Negative, -23, 0
	case 0x37:
		return CBORMajor1Negative, -24, 0
	// 0x38 	negative integer -1-n (one-byte uint8_t for n follows)
	case 0x38:
		return CBORMajor1Negative, 0, 1
	// 0x39 	negative integer -1-n (two-byte uint16_t for n follows)
	case 0x39:
		return CBORMajor1Negative, 0, 2
	// 0x3a 	negative integer -1-n (four-byte uint32_t for n follows)
	case 0x3a:
		return CBORMajor1Negative, 0, 4
	// 0x3b 	negative integer -1-n (eight-byte uint64_t for n follows)
	case 0x3b:
		return CBORMajor1Negative, 0, 8
	// 0x40..0x57 	byte string (0x00..0x17 bytes follow)
	case 0x40:
		return CBORMajor2Bytes, 0, 0
	case 0x41:
		return CBORMajor2Bytes, 1, 0
	case 0x42:
		return CBORMajor2Bytes, 2, 0
	case 0x43:
		return CBORMajor2Bytes, 3, 0
	case 0x44:
		return CBORMajor2Bytes, 4, 0
	case 0x45:
		return CBORMajor2Bytes, 5, 0
	case 0x46:
		return CBORMajor2Bytes, 6, 0
	case 0x47:
		return CBORMajor2Bytes, 7, 0
	case 0x48:
		return CBORMajor2Bytes, 8, 0
	case 0x49:
		return CBORMajor2Bytes, 9, 0
	case 0x4a:
		return CBORMajor2Bytes, 10, 0
	case 0x4b:
		return CBORMajor2Bytes, 11, 0
	case 0x4c:
		return CBORMajor2Bytes, 12, 0
	case 0x4d:
		return CBORMajor2Bytes, 13, 0
	case 0x4e:
		return CBORMajor2Bytes, 14, 0
	case 0x4f:
		return CBORMajor2Bytes, 15, 0
	case 0x50:
		return CBORMajor2Bytes, 16, 0
	case 0x51:
		return CBORMajor2Bytes, 17, 0
	case 0x52:
		return CBORMajor2Bytes, 18, 0
	case 0x53:
		return CBORMajor2Bytes, 19, 0
	case 0x54:
		return CBORMajor2Bytes, 20, 0
	case 0x55:
		return CBORMajor2Bytes, 21, 0
	case 0x56:
		return CBORMajor2Bytes, 22, 0
	case 0x57:
		return CBORMajor2Bytes, 23, 0
	// 0x58 	byte string (one-byte uint8_t for n, and then n bytes follow)
	case 0x58:
		return CBORMajor2Bytes, 0, 1
	// 0x59 	byte string (two-byte uint16_t for n, and then n bytes follow)
	case 0x59:
		return CBORMajor2Bytes, 0, 2
	// 0x5a 	byte string (four-byte uint32_t for n, and then n bytes follow)
	case 0x5a:
		return CBORMajor2Bytes, 0, 4
	// 0x5b 	byte string (eight-byte uint64_t for n, and then n bytes follow)
	case 0x5b:
		return CBORMajor2Bytes, 0, 8
	// 0x5f 	byte string, byte strings follow, terminated by "break"
	case 0x5f:
		return CBORMajor2Bytes, 0, -1
	// 0x60..0x77 	UTF-8 string (0x00..0x17 bytes follow)
	case 0x60:
		return CBORMajor3Text, 0, 0
	case 0x61:
		return CBORMajor3Text, 1, 0
	case 0x62:
		return CBORMajor3Text, 2, 0
	case 0x63:
		return CBORMajor3Text, 3, 0
	case 0x64:
		return CBORMajor3Text, 4, 0
	case 0x65:
		return CBORMajor3Text, 5, 0
	case 0x66:
		return CBORMajor3Text, 6, 0
	case 0x67:
		return CBORMajor3Text, 7, 0
	case 0x68:
		return CBORMajor3Text, 8, 0
	case 0x69:
		return CBORMajor3Text, 9, 0
	case 0x6a:
		return CBORMajor3Text, 10, 0
	case 0x6b:
		return CBORMajor3Text, 11, 0
	case 0x6c:
		return CBORMajor3Text, 12, 0
	case 0x6d:
		return CBORMajor3Text, 13, 0
	case 0x6e:
		return CBORMajor3Text, 14, 0
	case 0x6f:
		return CBORMajor3Text, 15, 0
	case 0x70:
		return CBORMajor3Text, 16, 0
	case 0x71:
		return CBORMajor3Text, 17, 0
	case 0x72:
		return CBORMajor3Text, 18, 0
	case 0x73:
		return CBORMajor3Text, 19, 0
	case 0x74:
		return CBORMajor3Text, 20, 0
	case 0x75:
		return CBORMajor3Text, 21, 0
	case 0x76:
		return CBORMajor3Text, 22, 0
	case 0x77:
		return CBORMajor3Text, 23, 0
	// 0x78 	UTF-8 string (one-byte uint8_t for n, and then n bytes follow)
	case 0x78:
		return CBORMajor3Text, 0, 1
	// 0x79 	UTF-8 string (two-byte uint16_t for n, and then n bytes follow)
	case 0x79:
		return CBORMajor3Text, 0, 2
	// 0x7a 	UTF-8 string (four-byte uint32_t for n, and then n bytes follow)
	case 0x7a:
		return CBORMajor3Text, 0, 4
	// 0x7b 	UTF-8 string (eight-byte uint64_t for n, and then n bytes follow)
	case 0x7b:
		return CBORMajor3Text, 0, 8
	// 0x7f 	UTF-8 string, UTF-8 strings follow, terminated by "break"
	case 0x7f:
		return CBORMajor3Text, 0, -1
	// 0x80..0x97 	array (0x00..0x17 data items follow)
	case 0x80:
		return CBORMajor4Array, 0, 0
	case 0x81:
		return CBORMajor4Array, 1, 0
	case 0x82:
		return CBORMajor4Array, 2, 0
	case 0x83:
		return CBORMajor4Array, 3, 0
	case 0x84:
		return CBORMajor4Array, 4, 0
	case 0x85:
		return CBORMajor4Array, 5, 0
	case 0x86:
		return CBORMajor4Array, 6, 0
	case 0x87:
		return CBORMajor4Array, 7, 0
	case 0x88:
		return CBORMajor4Array, 8, 0
	case 0x89:
		return CBORMajor4Array, 9, 0
	case 0x8a:
		return CBORMajor4Array, 10, 0
	case 0x8b:
		return CBORMajor4Array, 11, 0
	case 0x8c:
		return CBORMajor4Array, 12, 0
	case 0x8d:
		return CBORMajor4Array, 13, 0
	case 0x8e:
		return CBORMajor4Array, 14, 0
	case 0x8f:
		return CBORMajor4Array, 15, 0
	case 0x90:
		return CBORMajor4Array, 16, 0
	case 0x91:
		return CBORMajor4Array, 17, 0
	case 0x92:
		return CBORMajor4Array, 18, 0
	case 0x93:
		return CBORMajor4Array, 19, 0
	case 0x94:
		return CBORMajor4Array, 20, 0
	case 0x95:
		return CBORMajor4Array, 21, 0
	case 0x96:
		return CBORMajor4Array, 22, 0
	case 0x97:
		return CBORMajor4Array, 23, 0
	// 0x98 	array (one-byte uint8_t for n, and then n data items follow)
	case 0x98:
		return CBORMajor4Array, 0, 1
	// 0x99 	array (two-byte uint16_t for n, and then n data items follow)
	case 0x99:
		return CBORMajor4Array, 0, 2
	// 0x9a 	array (four-byte uint32_t for n, and then n data items follow)
	case 0x9a:
		return CBORMajor4Array, 0, 4
	// 0x9b 	array (eight-byte uint64_t for n, and then n data items follow)
	case 0x9b:
		return CBORMajor4Array, 0, 8
	// 0x9f 	array, data items follow, terminated by "break"
	case 0x9f:
		return CBORMajor4Array, 0, -1
	// 0xa0..0xb7 	map (0x00..0x17 pairs of data items follow)
	case 0xa0:
		return CBORMajor5Map, 0, 0
	case 0xa1:
		return CBORMajor5Map, 1, 0
	case 0xa2:
		return CBORMajor5Map, 2, 0
	case 0xa3:
		return CBORMajor5Map, 3, 0
	case 0xa4:
		return CBORMajor5Map, 4, 0
	case 0xa5:
		return CBORMajor5Map, 5, 0
	case 0xa6:
		return CBORMajor5Map, 6, 0
	case 0xa7:
		return CBORMajor5Map, 7, 0
	case 0xa8:
		return CBORMajor5Map, 8, 0
	case 0xa9:
		return CBORMajor5Map, 9, 0
	case 0xaa:
		return CBORMajor5Map, 10, 0
	case 0xab:
		return CBORMajor5Map, 11, 0
	case 0xac:
		return CBORMajor5Map, 12, 0
	case 0xad:
		return CBORMajor5Map, 13, 0
	case 0xae:
		return CBORMajor5Map, 14, 0
	case 0xaf:
		return CBORMajor5Map, 15, 0
	case 0xb0:
		return CBORMajor5Map, 16, 0
	case 0xb1:
		return CBORMajor5Map, 17, 0
	case 0xb2:
		return CBORMajor5Map, 18, 0
	case 0xb3:
		return CBORMajor5Map, 19, 0
	case 0xb4:
		return CBORMajor5Map, 20, 0
	case 0xb5:
		return CBORMajor5Map, 21, 0
	case 0xb6:
		return CBORMajor5Map, 22, 0
	case 0xb7:
		return CBORMajor5Map, 23, 0
	// 0xb8 	map (one-byte uint8_t for n, and then n pairs of data items follow)
	case 0xb8:
		return CBORMajor5Map, 0, 2
	// 0xb9 	map (two-byte uint16_t for n, and then n pairs of data items follow)
	case 0xb9:
		return CBORMajor5Map, 0, 4
	// 0xba 	map (four-byte uint32_t for n, and then n pairs of data items follow)
	case 0xba:
		return CBORMajor5Map, 0, 8
	// 0xbb 	map (eight-byte uint64_t for n, and then n pairs of data items follow)
	case 0xbb:
		return CBORMajor5Map, 0, -1
	// 0xbf 	map, pairs of data items follow, terminated by "break"
	case 0xbf:
		return CBORMajor5Map, 23, 0
	// 0xc0 	text-based date/time (data item follows; see Section 3.4.1)
	case 0xc0:
		return CBORMajor6Tag, 1, 1
	// 0xc1 	epoch-based date/time (data item follows; see Section 3.4.2)
	case 0xc1:
		return CBORMajor6Tag, 2, 1
	// 0xc2 	unsigned bignum (data item "byte string" follows)
	case 0xc2:
		return CBORMajor6Tag, 3, 1
	// 0xc3 	negative bignum (data item "byte string" follows)
	case 0xc3:
		return CBORMajor6Tag, 4, 1
	// 0xc4 	decimal Fraction (data item "array" follows; see Section 3.4.4)
	case 0xc4:
		return CBORMajor6Tag, 5, 1
	// 0xc5 	bigfloat (data item "array" follows; see Section 3.4.4)
	case 0xc5:
		return CBORMajor6Tag, 6, 1
	// 0xc6..0xd4 	(tag)
	case 0xc6:
		return CBORMajor6Tag, 7, 1
	case 0xc7:
		return CBORMajor6Tag, 7, 1
	case 0xc8:
		return CBORMajor6Tag, 7, 1
	case 0xc9:
		return CBORMajor6Tag, 7, 1
	case 0xca:
		return CBORMajor6Tag, 7, 1
	case 0xcb:
		return CBORMajor6Tag, 7, 1
	case 0xcc:
		return CBORMajor6Tag, 7, 1
	case 0xcd:
		return CBORMajor6Tag, 7, 1
	case 0xce:
		return CBORMajor6Tag, 7, 1
	case 0xcf:
		return CBORMajor6Tag, 7, 1
	case 0xd0:
		return CBORMajor6Tag, 7, 1
	case 0xd1:
		return CBORMajor6Tag, 7, 1
	case 0xd2:
		return CBORMajor6Tag, 7, 1
	case 0xd3:
		return CBORMajor6Tag, 7, 1
	case 0xd4:
		return CBORMajor6Tag, 7, 1
	// 0xd5..0xd7 	expected conversion (data item follows; see Section 3.4.5.2)
	case 0xd5:
		return CBORMajor6Tag, 7, 1
	case 0xd6:
		return CBORMajor6Tag, 7, 1
	case 0xd7:
		return CBORMajor6Tag, 7, 1
	// 0xd8..0xdb 	(more tags; 1/2/4/8 bytes of tag number and then a data item follow)
	case 0xd8:
		return CBORMajor6Tag, 7, 1
	case 0xd9:
		return CBORMajor6Tag, 7, 1
	case 0xda:
		return CBORMajor6Tag, 7, 1
	case 0xdb:
		return CBORMajor6Tag, 7, 1
	// 0xe0..0xf3 	(simple value)
	case 0xe0:
		return CBORMajor7Simple, 0, 0
	case 0xe1:
		return CBORMajor7Simple, 1, 0
	case 0xe2:
		return CBORMajor7Simple, 2, 0
	case 0xe3:
		return CBORMajor7Simple, 3, 0
	case 0xe4:
		return CBORMajor7Simple, 4, 0
	case 0xe5:
		return CBORMajor7Simple, 5, 0
	case 0xe6:
		return CBORMajor7Simple, 6, 0
	case 0xe7:
		return CBORMajor7Simple, 7, 0
	case 0xe8:
		return CBORMajor7Simple, 8, 0
	case 0xe9:
		return CBORMajor7Simple, 9, 0
	case 0xea:
		return CBORMajor7Simple, 10, 0
	case 0xeb:
		return CBORMajor7Simple, 11, 0
	case 0xec:
		return CBORMajor7Simple, 12, 0
	case 0xed:
		return CBORMajor7Simple, 13, 0
	case 0xee:
		return CBORMajor7Simple, 14, 0
	case 0xef:
		return CBORMajor7Simple, 15, 0
	case 0xf0:
		return CBORMajor7Simple, 16, 0
	case 0xf1:
		return CBORMajor7Simple, 17, 0
	case 0xf2:
		return CBORMajor7Simple, 18, 0
	case 0xf3:
		return CBORMajor7Simple, 19, 0
	// 0xf4 	false
	case 0xf4:
		return CBORMajor7Simple, 0xf4, 1
	// 0xf5 	true
	case 0xf5:
		return CBORMajor7Simple, 0xf5, 1
	// 0xf6 	null
	case 0xf6:
		return CBORMajor7Simple, 0xf6, 1
	// 0xf7 	undefined
	case 0xf7:
		return CBORMajor7Simple, 0xf7, 1
	// 0xf8 	(simple value, one byte follows)
	case 0xf8:
		return CBORMajor7Simple, 0xf8, 1
	// 0xf9 	half-precision float (two-byte IEEE 754)
	case 0xf9:
		return CBORMajor7Simple, 0xf9, 1
	// 0xfa 	single-precision float (four-byte IEEE 754)
	case 0xfa:
		return CBORMajor7Simple, 0xfa, 1
	// 0xfb 	double-precision float (eight-byte IEEE 754)
	case 0xfb:
		return CBORMajor7Simple, 0xfb, 1
	// 0xff 	"break" stop code
	case 0xff:
		return CBORMajor7Simple, 0xff, 1
	}
	return 0, 0, 0
}
