package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"

	a2j "github.com/takanoriyanagitani/go-cbor2json/arr2json"
	a2js "github.com/takanoriyanagitani/go-cbor2json/arr2json/std"

	c2a "github.com/takanoriyanagitani/go-cbor2json/cbor2arr"
	c2aa "github.com/takanoriyanagitani/go-cbor2json/cbor2arr/amacker"

	c2j "github.com/takanoriyanagitani/go-cbor2json/cbor2arr2json"
)

func rdr2wtr(rdr io.Reader, wtr io.Writer) error {
	var br io.Reader = bufio.NewReader(rdr)
	var bw *bufio.Writer = bufio.NewWriter(wtr)
	defer bw.Flush()

	var cbor2arr c2a.CborToArray = c2aa.CborToArrayNew(br).
		ToConverter()

	var arr2json a2j.ArrayToJson = a2js.ArrayToJsonNew(bw).
		ToConverter()

	var cbor2json c2j.CborToArrayToJson = c2j.CborToArrayToJson{
		ArrayToJson: arr2json,
		CborToArray: cbor2arr,
	}

	return cbor2json.ConvertAll(context.Background())
}

func stdin2stdout() error {
	return rdr2wtr(os.Stdin, os.Stdout)
}

func main() {
	e := stdin2stdout()
	if nil != e {
		log.Printf("%v\n", e)
	}
}
