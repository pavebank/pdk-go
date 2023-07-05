package paveapps

import (
	"encoding/json"

	"github.com/extism/go-pdk"
)

type Memory struct {
	pdk.Memory
}

func Input(v interface{}) error {
	input := pdk.Input()
	return json.Unmarshal(input, &v)
}

func MemoryToBytes(mem Memory) []byte {
	buf := make([]byte, mem.Length())
	mem.Load(buf)
	return buf
}

func MemoryToStruct(mem Memory, v interface{}) error {
	bytes := MemoryToBytes(mem)
	return json.Unmarshal(bytes, &v)
}

func BytesToMemory(b []byte) Memory {
	return Memory{pdk.AllocateBytes(b)}
}

func FindMemory(offset uint64) Memory {
	return Memory{pdk.FindMemory(offset)}
}

func StructToMemory(v interface{}) (mem Memory, err error) {
	b, err := json.Marshal(&v)

	if err != nil {
		return
	}

	mem = BytesToMemory(b)
	return
}

func GetConfig(key string) (string, bool) {
	return pdk.GetConfig(key)
}
