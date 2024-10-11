package cbor2map_test

import (
	"testing"

	"bytes"

	c2m "github.com/takanoriyanagitani/go-cbor2json/cbor2map"
	c2ma "github.com/takanoriyanagitani/go-cbor2json/cbor2map/amacker"
)

func TestCborToMapAmacker(t *testing.T) {
	t.Parallel()

	t.Run("CborToMap", func(t *testing.T) {
		t.Parallel()

		t.Run("helo", func(t *testing.T) {
			t.Parallel()

			var rdr = bytes.NewReader([]byte{
				0xa1,
				0x65,
				0x68,
				0x65,
				0x6c,
				0x6c,
				0x6f,
				0x65,
				0x77,
				0x6f,
				0x72,
				0x6c,
				0x64,
			})

			var cnv c2m.CborToMap = c2ma.CborToMapNew(rdr).ToConverter()

			var buf map[string]any

			e := cnv(&buf)
			if nil != e {
				t.Fatalf("unexpected err: %v\n", e)
			}

			val := buf["hello"]
			if val != "world" {
				t.Fatalf("unexpected value: %v\n", val)
			}
		})
	})
}
