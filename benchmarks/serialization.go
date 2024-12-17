package benchmarks

import (
	"bytes"
	"encoding/gob"
	"encoding/json"

	jsoniter "github.com/json-iterator/go"
	"github.com/vmihailenco/msgpack/v5"
)

func SerializeJSON(data map[string]string) ([]byte, error) {
	return json.Marshal(data)
}

func DeserializeJSON(data []byte) (map[string]string, error) {
	var result map[string]string

	err := json.Unmarshal(data, &result)

	return result, err
}

func SerializeBinary(data map[string]string) ([]byte, error) {
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)

	return buf.Bytes(), err
}

func DeserializeBinary(data []byte) (map[string]string, error) {
	var result map[string]string

	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(&result)

	return result, err
}

func SerializeMsgPack(data map[string]string) ([]byte, error) {
	return msgpack.Marshal(data)
}

func DeserializeMsgPack(data []byte) (map[string]string, error) {
	var result map[string]string

	err := msgpack.Unmarshal(data, &result)

	return result, err
}

func SerializeJSONIter(data map[string]string) ([]byte, error) {
	return jsoniter.Marshal(data)
}

func DeserializeJSONIter(data []byte) (map[string]string, error) {
	var result map[string]string

	err := jsoniter.Unmarshal(data, &result)

	return result, err
}
