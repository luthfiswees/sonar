package handler

import (
		"encoding/json"
		"net/http"

		"github.com/luthfiswees/sonar/checker"
)

type Response struct {
	Data string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	result := checker.Check()
	data   := &Response{Data: result}

	js, err := json.Marshal(data)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
	}
	
	w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}