package net

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type respWriter struct {
	http.ResponseWriter
	Code int
	buf  *bytes.Buffer
}

func (w *respWriter) Write(p []byte) (int, error) {
	return w.buf.Write(p)
}

func (w *respWriter) WriteHeader(code int) {
	w.Code = code
	w.ResponseWriter.WriteHeader(code)
}

func writeBytes(w http.ResponseWriter, code int, v []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(v)
}

func WriteJSON(w http.ResponseWriter, code int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}
