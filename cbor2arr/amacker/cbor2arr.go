package cbor2arr

import (
	"io"

	ac "github.com/fxamacker/cbor/v2"

	c2a "github.com/takanoriyanagitani/go-cbor2json/cbor2arr"
)

type CborToArray struct {
	*ac.Decoder
}

func (c CborToArray) DecodeToArray(arr *[]any) error {
	return c.Decoder.Decode(arr)
}

func (c CborToArray) ToConverter() c2a.CborToArray {
	return c.DecodeToArray
}

func CborToArrayNew(rdr io.Reader) CborToArray {
	return CborToArray{
		Decoder: ac.NewDecoder(rdr),
	}
}
