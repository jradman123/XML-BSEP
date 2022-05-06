package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

// renderJSON renders 'v' as JSON and writes it as a response into w.
func renderJSON(ctx context.Context, w http.ResponseWriter, v interface{}) {

	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func decodeBody(ctx context.Context, r io.Reader) (*NewUser, error) {

	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()
	var rt NewUser //kreiramo svoju sema za nasu strukturu koja govori kako ce se json namapirati na nas structure
	//meni ovo lici na dto pa valjda je to to
	//dole prosledjujemo adresu ovjektaa

	if err := dec.Decode(&rt); err != nil {
		return nil, err
	}
	return &rt, nil
}
