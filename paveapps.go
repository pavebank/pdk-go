package paveapps

import (
	"encoding/json"

	"github.com/extism/go-pdk"
)

// Offset type is a uint64 type
// Offset to the location in linear memory where allocations begins
type Offset uint64

// Unmarshal PaveApps Trigger handler input into struct
func Input(v interface{}) error {
	input := pdk.Input()
	return json.Unmarshal(input, &v)
}

// Return the byte slices based on the Offset in linear memory
func OffsetToBytes(o Offset) []byte {
	mem := pdk.FindMemory(uint64(o))
	buf := make([]byte, mem.Length())
	mem.Load(buf)
	return buf
}

// Unmarshal data in linear memory into struct based on Offset
func OffsetToStruct(o Offset, v interface{}) error {
	bytes := OffsetToBytes(o)
	return json.Unmarshal(bytes, &v)
}

// Allocate and store bytes in linear memory, and return the Offset to that allocation
func BytesToOffset(b []byte) Offset {
	mem := pdk.AllocateBytes(b)
	return Offset(mem.Offset())
}

// Marshal struct to byte slice and store in linear memory, returning the Offset
func StructToOffset(v interface{}) (o Offset, err error) {
	b, err := json.Marshal(&v)

	if err != nil {
		return
	}

	o = BytesToOffset(b)
	return
}

// Get config that is stored during creation of PaveApps
func GetConfig(key string) (string, bool) {
	return pdk.GetConfig(key)
}
