package anchor

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"

	"github.com/vmihailenco/msgpack"
)

func DecodeProof(data string) (interface{}, error) {
	b, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	z, err := zlib.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	defer z.Close()
	var m map[string]interface{}
	if err := msgpack.NewDecoder(z).UseJSONTag(true).Decode(&m); err != nil {
		return nil, err
	}
	return m, nil
}
