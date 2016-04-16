package main

import (
	"encoding/json"
	"github.com/miguelmejiamontes/tools"
	"net/http"
)

type headers struct {
	IPAddress string `json:"ipaddress"`
	Lang      string `json:"language"`
	Software  string `json:"software"`
}

func main() {
	http.HandleFunc("/api/whoami", returnHeaders)
	http.HandleFunc("/", showHome)
	http.ListenAndServe(tools.GetPort(), nil)
}

func showHome(w http.ResponseWriter, r *http.Request) {
	w.Write(tools.RenderHome())
}

func returnHeaders(w http.ResponseWriter, r *http.Request) {
	ipadr := r.RemoteAddr
	lang := r.Header.Get("Accept-Language")
	softw := r.UserAgent()

	builtHeader := headers{ipadr, lang, softw}
	obj, _ := json.Marshal(builtHeader)
	w.Write([]byte(obj))
}
