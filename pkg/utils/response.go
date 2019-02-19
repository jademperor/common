package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/jademperor/common/pkg/code"
)

func response(w http.ResponseWriter, s string) {
	log.Println(s)
	_, err := io.WriteString(w, s)
	if err != nil {
		log.Printf("Error: response err: %s\n", err.Error())
	}
}

// ResponseString ...
func ResponseString(w http.ResponseWriter, s string) {
	response(w, s)
}

// ResponseJSON ...
func ResponseJSON(w http.ResponseWriter, i interface{}) {
	bs, err := json.Marshal(i)
	if err != nil {
		bs, _ = json.Marshal(code.NewCodeInfo(code.CodeSystemErr, err.Error()))
		log.Printf("Error: get an err: %s\n", err.Error())
	}

	// set header
	w.Header().Set("Content-Type", "application/json")

	response(w, string(bs))
}
