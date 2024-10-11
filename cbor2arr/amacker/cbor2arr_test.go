package cbor2arr_test

import (
	"testing"

	"bytes"
	"io"

	c2a "github.com/takanoriyanagitani/go-cbor2json/cbor2arr"
	c2aa "github.com/takanoriyanagitani/go-cbor2json/cbor2arr/amacker"
)

func TestCborToArrAmacker(t *testing.T) {
	t.Parallel()

	t.Run("CborToArr", func(t *testing.T) {
		t.Parallel()

		t.Run("numbers", func(t *testing.T) {
			t.Parallel()

			var rdr io.Reader = bytes.NewReader([]byte{
				0x82,
				0x01,
				0x02,
			})
			var cnv c2a.CborToArray = c2aa.CborToArrayNew(rdr).
				ToConverter()

			var buf []any
			e := cnv(&buf)
			if nil != e {
				t.Fatalf("unexpected error: %v\n", e)
			}

			if 2 != len(buf) {
				t.Fatalf("unexpected length: %v\n", len(buf))
			}

			var a0 any = buf[0]
			var a1 any = buf[1]

			switch i0 := a0.(type) {
			case uint64:
				if 1 != i0 {
					t.Fatalf("unexpected value: %v\n", i0)
				}
			default:
				t.Fatalf("unexpected type: %v\n", i0)
			}

			switch i1 := a1.(type) {
			case uint64:
				if 2 != i1 {
					t.Fatalf("unexpected value: %v\n", i1)
				}
			default:
				t.Fatalf("unexpected type: %v\n", i1)
			}
		})
	})
}
