package cbor2arr2json_test

import (
	"testing"

	"context"
	"io"

	a2j "github.com/takanoriyanagitani/go-cbor2json/arr2json"
	c2a "github.com/takanoriyanagitani/go-cbor2json/cbor2arr"

	c2j "github.com/takanoriyanagitani/go-cbor2json/cbor2arr2json"
)

func TestCborToArrayToJson(t *testing.T) {
	t.Parallel()

	t.Run("CborToArrayToJson", func(t *testing.T) {
		t.Parallel()

		t.Run("ConvertAll", func(t *testing.T) {
			t.Parallel()

			t.Run("empty", func(t *testing.T) {
				t.Parallel()

				var c2 c2a.CborToArray = func(_ *[]any) error {
					return io.EOF
				}
				var tj a2j.ArrayToJson = func([]any) error { return nil }

				var cnv c2j.CborToArrayToJson = c2j.CborToArrayToJson{
					CborToArray: c2,
					ArrayToJson: tj,
				}

				e := cnv.ConvertAll(context.Background())
				if nil != e {
					t.Fatalf("unexpected err: %v\n", e)
				}
			})

			t.Run("single empty row", func(t *testing.T) {
				t.Parallel()

				var scnt int = 0
				var c2 c2a.CborToArray = func(_ *[]any) error {
					switch scnt {
					case 0:
						scnt += 1
						return nil
					default:
						return io.EOF
					}
				}

				var tcnt int = 0
				var tj a2j.ArrayToJson = func(a []any) error {
					if 0 != len(a) {
						t.Fatalf("unexpected length: %v\n", len(a))
					}
					tcnt += 1
					return nil
				}

				var cnv c2j.CborToArrayToJson = c2j.CborToArrayToJson{
					CborToArray: c2,
					ArrayToJson: tj,
				}

				e := cnv.ConvertAll(context.Background())
				if nil != e {
					t.Fatalf("unexpected err: %v\n", e)
				}

				if 1 != tcnt {
					t.Fatalf("unexpected cnt: %v\n", tcnt)
				}
			})

			t.Run("two rows", func(t *testing.T) {
				t.Parallel()

				var scnt int = 0
				var c2 c2a.CborToArray = func(a *[]any) error {
					switch scnt {
					case 0:
						*a = append(*a, 42)
						scnt += 1
						return nil
					case 1:
						*a = append(*a, 634)
						scnt += 1
						return nil
					default:
						return io.EOF
					}
				}

				var integers []int
				var tj a2j.ArrayToJson = func(a []any) error {
					if 1 != len(a) {
						t.Fatalf("unexpected length: %v\n", len(a))
					}
					var a0 any = a[0]
					switch i0 := a0.(type) {
					case int:
						integers = append(integers, i0)
					default:
						t.Fatalf("unexpected type: %v\n", i0)
					}
					return nil
				}

				var cnv c2j.CborToArrayToJson = c2j.CborToArrayToJson{
					CborToArray: c2,
					ArrayToJson: tj,
				}

				e := cnv.ConvertAll(context.Background())
				if nil != e {
					t.Fatalf("unexpected err: %v\n", e)
				}

				if 2 != len(integers) {
					t.Fatalf("unexpected cnt: %v\n", len(integers))
				}

				var i0 int = integers[0]
				if 42 != i0 {
					t.Fatalf("unexpected value: %v\n", i0)
				}

				var i1 int = integers[1]
				if 634 != i1 {
					t.Fatalf("unexpected value: %v\n", i1)
				}
			})

			t.Run("three rows with null", func(t *testing.T) {
				t.Parallel()

				var scnt int = 0
				var c2 c2a.CborToArray = func(a *[]any) error {
					switch scnt {
					case 0:
						*a = append(*a, 42)
						scnt += 1
						return nil
					case 1:
						*a = append(*a, 634)
						scnt += 1
						return nil
					case 2:
						*a = append(*a, nil)
						scnt += 1
						return nil
					default:
						return io.EOF
					}
				}

				var integers []int
				var tj a2j.ArrayToJson = func(a []any) error {
					if 1 != len(a) {
						t.Fatalf("unexpected length: %v\n", len(a))
					}
					var a0 any = a[0]
					switch i0 := a0.(type) {
					case int:
						integers = append(integers, i0)
					case nil:
					default:
						t.Fatalf("unexpected type: %v\n", i0)
					}
					return nil
				}

				var cnv c2j.CborToArrayToJson = c2j.CborToArrayToJson{
					CborToArray: c2,
					ArrayToJson: tj,
				}

				e := cnv.ConvertAll(context.Background())
				if nil != e {
					t.Fatalf("unexpected err: %v\n", e)
				}

				if 2 != len(integers) {
					t.Fatalf("unexpected cnt: %v\n", len(integers))
				}

				var i0 int = integers[0]
				if 42 != i0 {
					t.Fatalf("unexpected value: %v\n", i0)
				}

				var i1 int = integers[1]
				if 634 != i1 {
					t.Fatalf("unexpected value: %v\n", i1)
				}
			})
		})
	})
}
