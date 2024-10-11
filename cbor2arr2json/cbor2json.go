package cbor2arr2json

import (
	"context"
	"errors"
	"io"

	a2j "github.com/takanoriyanagitani/go-cbor2json/arr2json"
	c2a "github.com/takanoriyanagitani/go-cbor2json/cbor2arr"
)

type CborToArrayToJson struct {
	c2a.CborToArray
	a2j.ArrayToJson
}

func (c CborToArrayToJson) Convert(buf *[]any) error {
	eser := c.CborToArray(buf)
	if nil != eser {
		return eser
	}

	return c.ArrayToJson(*buf)
}

func (c CborToArrayToJson) ConvertAll(ctx context.Context) error {
	var buf []any
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		e := c.Convert(&buf)
		if nil != e {
			if !errors.Is(e, io.EOF) {
				return e
			}
			return nil
		}
	}
}
