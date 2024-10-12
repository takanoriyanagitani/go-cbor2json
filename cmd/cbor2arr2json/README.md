## Install

```
go install github.com/takanoriyanagitani/go-cbor2json/cmd/cbor2arr2json@v1.0.0
```

## Example

```
echo -n ggECggECggECggMEggMEggME |
	base64 --decode |
	cbor2arr2json
```

```
[1,2]
[1,2]
[1,2]
[3,4]
[3,4]
[3,4]
```
