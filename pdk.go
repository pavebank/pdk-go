package pdk

import (
	"encoding/json"
	"fmt"

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

// ToArgs function that takes in pointer and transforms it to Args
// Accepts *[]byte and struct{} type only
func ToArgs(input interface{}) (Args, error) {
	switch v := input.(type) {
	case *[]byte:
		return bytesToArgs(*v), nil
	case struct{}:
		a, err := structToArgs(v)
		if err != nil {
			return 0, err
		}
		return a, nil
	default:
		return 0, fmt.Errorf("Invalid type: only *[]byte and struct are accepted")
	}
}

// FromArgs function that takes in pointer and returns data back
// Accepts *[]byte and struct{} type only
func FromArgs(a Args, v interface{}) error {
	switch v := v.(type) {
	case *[]byte:
		*v = argsToBytes(a)
		return nil
	case struct{}:
		err := argsToStruct(a, &v)
		if err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("Invalid type: only *[]byte and struct are accepted")
	}
}

// Get config that is stored during creation of PaveApps
func GetConfig(key string) (string, bool) {
	return pdk.GetConfig(key)
}
