package debugerrorce

import (
	"encoding/json"
	"errors"
	"fmt"
)

type DestinationFormat int

const (
	FormatText DestinationFormat = iota
	FormatJSON
	FormatPrettyJson
)

// from fiber
func Format(format DestinationFormat, body interface{}) (string, error) {
	switch format {
	case FormatText:
		switch body.(type) {
		case string:
			return body.(string), nil
		case []byte:
			return string(body.([]byte)), nil
		default:
			return fmt.Sprintf("%v", body), nil
		}
	case FormatJSON:
		switch body.(type) {
		case string:
			return "{\"" + body.(string) + "\"}", nil
		case []byte:
		default:
		}
		json, err := json.Marshal(body)
		if err != nil {
			return "", err
		}
		return string(json), nil
	case FormatPrettyJson:
		switch body.(type) {
		case string:
			return "{ \"" + body.(string) + "\" }", nil
		case []byte:
		default:
		}
		json, err := json.MarshalIndent(body, "", "    ")
		if err != nil {
			return "", err
		}
		return string(json), nil
	}
	return "", errors.New("Undefined handling in Format")
}
