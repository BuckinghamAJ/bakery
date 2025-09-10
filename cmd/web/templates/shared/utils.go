package shared

import (
	"encoding/json"
	"log/slog"
	"reflect"
)

type Pair struct {
	Key, Value string
}

// Helper function to create json alpine.js x-data can understand
func MakeAlpineData(pairs ...Pair) string {
	var alpineMap = make(map[string]string)

	for _, pair := range pairs {
		if reflect.TypeOf(pair.Key).String() != "string" || reflect.TypeOf(pair.Value).String() != "string" {
			slog.Error("Not passing strings into the MakeAlpineData!", "key", pair.Key, "value", pair.Value)
			panic("Not passing strings into the MakeAlpineData")
		}

		alpineMap[pair.Key] = pair.Value
	}

	jEncoded, err := json.Marshal(alpineMap)
	if err != nil {
		slog.Error("Failed to marshal to JSON", slog.String("error", err.Error()))
		return ""
	}

	return string(jEncoded)

}

// Helper function to create json alpine.js x-data can understand
func SetAplineData(key string, value string) string {
	alpineMap := map[string]string{
		key: value,
	}

	jEncoded, err := json.Marshal(alpineMap)
	if err != nil {
		slog.Error("Failed to marshal to JSON", slog.String("error", err.Error()))
		return ""
	}

	return string(jEncoded)
}
