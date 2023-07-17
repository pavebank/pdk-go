package pdk

import (
	"encoding/json"

	"github.com/extism/go-pdk"
)

// Return the byte slices based on the Args in linear memory
func argsToBytes(a Args) []byte {
	mem := pdk.FindMemory(uint64(a))
	buf := make([]byte, mem.Length())
	mem.Load(buf)
	return buf
}

// Unmarshal data in linear memory into struct based on Args
func argsToStruct(a Args, v interface{}) error {
	bytes := argsToBytes(a)
	return json.Unmarshal(bytes, &v)
}

// Allocate and store bytes in linear memory, and return the Args to that allocation
func bytesToArgs(b []byte) Args {
	mem := pdk.AllocateBytes(b)
	return Args(mem.Offset())
}

// Marshal struct to byte slice and store in linear memory, returning the Args
func structToArgs(v interface{}) (a Args, err error) {
	b, err := json.Marshal(&v)

	if err != nil {
		return
	}

	a = bytesToArgs(b)
	return
}
