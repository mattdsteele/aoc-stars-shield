package aoc

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Shield struct {
	SchemaVersion int    `json:"schemaVersion"`
	Label         string `json:"label"`
	Message       string `json:"message"`
}

func Shields(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stars := Stars()
	val := Shield{SchemaVersion: 1, Label: "advent-of-code", Message: fmt.Sprintf("%d stars", stars)}
	j, _ := json.Marshal(val)
	fmt.Fprint(w, string(j))
}
