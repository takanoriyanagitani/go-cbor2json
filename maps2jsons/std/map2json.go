package map2json

import (
	"encoding/json"
	"io"

	m2j "github.com/takanoriyanagitani/go-cbor2json/maps2jsons"
)

type MapToJsonStd struct {
	*json.Encoder
}

func (s MapToJsonStd) MapToJson(m map[string]any) error {
	return s.Encoder.Encode(m)
}

func (s MapToJsonStd) ToConverter() m2j.MapToJson {
	return s.MapToJson
}

func MapToJsonStdNew(wtr io.Writer) MapToJsonStd {
	return MapToJsonStd{
		Encoder: json.NewEncoder(wtr),
	}
}
