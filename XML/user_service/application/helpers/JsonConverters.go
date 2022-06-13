package helpers

import (
	"common/module/logger"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"user/module/domain/dto"
)

type JsonConverters struct {
	l *logger.Logger
}

func NewJsonConverters(l *logger.Logger) *JsonConverters {
	return &JsonConverters{l}
}

func RenderJSON(ctx context.Context, w http.ResponseWriter, v interface{}) {

	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func DecodeBody(ctx context.Context, r io.Reader) (*dto.NewUser, error) {

	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()
	var rt dto.NewUser //kreiramo svoju sema za nasu strukturu koja govori kako ce se json namapirati na nas structure
	//meni ovo lici na dto pa valjda je to to
	//dole prosledjujemo adresu ovjektaa

	if err := dec.Decode(&rt); err != nil {
		return nil, err
	}
	return &rt, nil
}
