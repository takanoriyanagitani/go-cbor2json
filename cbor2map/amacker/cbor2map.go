package cbor2map

import (
	"io"

	ac "github.com/fxamacker/cbor/v2"

	c2m "github.com/takanoriyanagitani/go-cbor2json/cbor2map"
)

type CborToMap struct {
	*ac.Decoder
}

func (c CborToMap) DecodeToMap(m *map[string]any) error {
	return c.Decoder.Decode(m)
}

func (c CborToMap) ToConverter() c2m.CborToMap {
	return c.DecodeToMap
}

func CborToMapNew(rdr io.Reader) CborToMap {
	return CborToMap{
		Decoder: ac.NewDecoder(rdr),
	}
}
