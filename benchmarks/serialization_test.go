package benchmarks_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/KazikovAP/benchmarks/benchmarks"
)

func generateData(size int) map[string]string {
	data := make(map[string]string, size)
	value := strings.Repeat("v", 10)

	for i := 0; i < size; i++ {
		key := "key" + strconv.Itoa(i)
		data[key] = value
	}

	return data
}

func BenchmarkSerialization(b *testing.B) {
	tests := []struct {
		name        string
		serialize   func(map[string]string) ([]byte, error)
		deserialize func([]byte) (map[string]string, error)
	}{
		{"JSON", benchmarks.SerializeJSON, benchmarks.DeserializeJSON},
		{"Binary", benchmarks.SerializeBinary, benchmarks.DeserializeBinary},
		{"MsgPack", benchmarks.SerializeMsgPack, benchmarks.DeserializeMsgPack},
		{"JSONIter", benchmarks.SerializeJSONIter, benchmarks.DeserializeJSONIter},
	}

	sizes := map[string]int{
		"1KB":  128,
		"1MB":  128000,
		"10MB": 1280000,
	}

	for _, test := range tests {
		for sizeName, size := range sizes {
			data := generateData(size)

			b.Run("Serialize_"+test.name+"_"+sizeName, func(b *testing.B) {
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					_, err := test.serialize(data)
					if err != nil {
						b.Fatalf("Serialization failed: %v", err)
					}
				}
			})

			serialized, err := test.serialize(data)
			if err != nil {
				b.Fatalf("Failed to serialize data: %v", err)
			}

			b.Run("Deserialize_"+test.name+"_"+sizeName, func(b *testing.B) {
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					_, err := test.deserialize(serialized)
					if err != nil {
						b.Fatalf("Deserialization failed: %v", err)
					}
				}
			})
		}
	}
}
