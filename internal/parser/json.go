package parser

import (
	"encoding/json"
	"streams/internal/model"
)

func Parse(in chan []byte, out chan string) {
	for msg := range in {
		var topupJSON model.TopupSchema
		_ = json.Unmarshal(msg, &topupJSON)
		jsonString, _ := json.Marshal(topupJSON)
		out <- string(jsonString)
	}
}
