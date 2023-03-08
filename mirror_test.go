package mirror

import (
	"encoding/hex"
	"testing"
	"time"

	"github.com/fxamacker/cbor/v2"
)

func TestNop(t *testing.T) {
}

func TestIs(t *testing.T) {
	opts := cbor.CanonicalEncOptions()
	opts.TimeTag = cbor.EncTagRequired
	em, err := opts.EncMode()
	if err != nil {
		t.Fatalf("cbor encoding failed: %v", err)
	}

	t.Run("isPositive", func(t *testing.T) {
		data := 123
		b, err := em.Marshal(data)
		if err != nil {
			t.Fatalf("failed to marshal data: %v", err)
		}
		is := IsPositive(b)
		if !is {
			t.Errorf("Number %v (0x%v) incorrectly classified: %v", data, hex.EncodeToString(b), is)
		}
	})

	t.Run("isNegative", func(t *testing.T) {
		data := -123
		b, err := em.Marshal(data)
		if err != nil {
			t.Fatalf("failed to marshal data: %v", err)
		}
		is := IsNegative(b)
		if !is {
			t.Errorf("Number %v (0x%v) incorrectly classified: %v", data, hex.EncodeToString(b), is)
		}
	})

	t.Run("isInteger", func(t *testing.T) {
		data := 123
		b, err := em.Marshal(data)
		if err != nil {
			t.Fatalf("failed to marshal data: %v", err)
		}
		is := IsInteger(b)
		if !is {
			t.Errorf("Number %v (0x%v) incorrectly classified: %v", data, hex.EncodeToString(b), is)
		}
	})

	t.Run("isFloat", func(t *testing.T) {
		data := 123.456
		b, err := em.Marshal(data)
		if err != nil {
			t.Fatalf("failed to marshal data: %v", err)
		}
		is := IsFloat(b)
		if !is {
			t.Errorf("Number %v (0x%v) incorrectly classified: %v", data, hex.EncodeToString(b), is)
		}
	})

	t.Run("isBytes", func(t *testing.T) {
		data := []byte{0x1,0x2,0x3}
		b, err := em.Marshal(data)
		if err != nil {
			t.Fatalf("failed to marshal data: %v", err)
		}
		is := IsBytes(b)
		if !is {
			t.Errorf("Data %v (0x%v) incorrectly classified: %v", data, hex.EncodeToString(b), is)
		}
	})

	t.Run("isText", func(t *testing.T) {
		data := "hello world"
		b, err := em.Marshal(data)
		if err != nil {
			t.Fatalf("failed to marshal data: %v", err)
		}
		is := IsText(b)
		if !is {
			t.Errorf("String %v (0x%v) incorrectly classified: %v", data, hex.EncodeToString(b), is)
		}
	})

	t.Run("isArray", func(t *testing.T) {
		data := []int{1,2,3}
		b, err := em.Marshal(data)
		if err != nil {
			t.Fatalf("failed to marshal data: %v", err)
		}
		is := IsArray(b)
		if !is {
			t.Errorf("Array %v (0x%v) incorrectly classified: %v", data, hex.EncodeToString(b), is)
		}
	})

	t.Run("isMap", func(t *testing.T) {
		data := map[string]string{ "a": "hello", "b" : "world" }
		b, err := em.Marshal(data)
		if err != nil {
			t.Fatalf("failed to marshal data: %v", err)
		}
		is := IsMap(b)
		if !is {
			t.Errorf("Map %v (0x%v) incorrectly classified: %v", data, hex.EncodeToString(b), is)
		}
	})

	t.Run("isTag", func(t *testing.T) {
		data := time.Now() // time is tagged with when the option EncTagRequired is set
		b, err := em.Marshal(data)
		if err != nil {
			t.Fatalf("failed to marshal data: %v", err)
		}
		is := IsTag(b)
		if !is {
			t.Errorf("String %v (0x%v) incorrectly classified: %v", data, hex.EncodeToString(b), is)
		}
	})

	/*
	Simple encoding still in beta in fxamacker/cbor
	t.Run("isSimple", func(t *testing.T) {
		data := cbor.SimpleValue(123)
		b, err := em.Marshal(data)
		if err != nil {
			t.Fatalf("failed to marshal data: %v", err)
		}
		is := IsTag(b)
		if !is {
			t.Errorf("Simple %v (0x%v) incorrectly classified: %v", data, hex.EncodeToString(b), is)
		}
	})
	*/

	t.Run("isSimple", func(t *testing.T) {
		b := []byte{0xe3}
		is := IsSimple(b)
		if !is {
			t.Errorf("Simple %v (0x%v) incorrectly classified: %v", 3, hex.EncodeToString(b), is)
		}
	})


}
