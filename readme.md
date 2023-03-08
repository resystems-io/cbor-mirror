# CBOR Mirror

A reflection library for [CBOR][cbor.io] ([RFC8949][rfc8949]) that is focused
on providing datatype information for CBOR.

It can be used in conjunction with CBOR marshaling libraries like
[fxamacker/cbor][fxamacker-cbor], when dynamic interpretation of CBOR is
required. That is, use `cbor-mirror` to obtain the type information and build up
interface types that can then be used drive decoding.

[cbor.io]:http://cbor.io/ "CBOR - Concise Binary Object Representation"
[rfc8949]:https://www.rfc-editor.org/rfc/rfc8949.html "RFC 8949 Concise Binary Object Representation (CBOR)"
[fxamacker-cbor]:https://github.com/fxamacker/cbor "fxamacker/cbor CBOR Library for Go"
