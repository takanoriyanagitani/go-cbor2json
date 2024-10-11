package arr2json

import (
	"encoding/json"
	"io"

	a2j "github.com/takanoriyanagitani/go-cbor2json/arr2json"
)

type ArrayToJson struct {
	*json.Encoder
}

func (j ArrayToJson) ToJson(a []any) error {
	return j.Encoder.Encode(a)
}

func (j ArrayToJson) ToConverter() a2j.ArrayToJson {
	return j.ToJson
}

func ArrayToJsonNew(wtr io.Writer) ArrayToJson {
	return ArrayToJson{
		Encoder: json.NewEncoder(wtr),
	}
}
