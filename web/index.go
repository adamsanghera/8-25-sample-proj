package web

import (
	"net/http"
)

func (ws *Web) Index(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Under construction"))
}
