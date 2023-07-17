package pdk

import (
	"encoding/json"

	"github.com/extism/go-pdk"
)

// Args type is a uint64 type
// Args to the location in linear memory where allocations begins
type Args uint64

// Unmarshal PaveApps Trigger handler input into struct
func Input(v interface{}) error {
	input := pdk.Input()
	return json.Unmarshal(input, &v)
}

// Return the byte slices based on the Args in linear memory
func ArgsToBytes(a Args) []byte {
	mem := pdk.FindMemory(uint64(a))
	buf := make([]byte, mem.Length())
	mem.Load(buf)
	return buf
}

// Unmarshal data in linear memory into struct based on Args
func ArgsToStruct(a Args, v interface{}) error {
	bytes := ArgsToBytes(a)
	return json.Unmarshal(bytes, &v)
}

// Allocate and store bytes in linear memory, and return the Args to that allocation
func BytesToArgs(b []byte) Args {
	mem := pdk.AllocateBytes(b)
	return Args(mem.Offset())
}

// Marshal struct to byte slice and store in linear memory, returning the Args
func StructToArgs(v interface{}) (a Args, err error) {
	b, err := json.Marshal(&v)

	if err != nil {
		return
	}

	a = BytesToArgs(b)
	return
}

// Get config that is stored during creation of PaveApps
func GetConfig(key string) (string, bool) {
	return pdk.GetConfig(key)
}
