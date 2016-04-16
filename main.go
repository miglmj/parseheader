package main

import (
	"encoding/json"
	"github.com/miguelmejiamontes/tools"
	"net/http"
	"strings"
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
	ipadr := strings.Split(r.RemoteAddr, ":")[0]
	lang := strings.Split(r.Header.Get("Accept-Language"), ",")[0]
	softw := extractSoftware(r.UserAgent())

	builtHeader := headers{ipadr, lang, softw}
	obj, _ := json.Marshal(builtHeader)
	w.Write([]byte(obj))
}

func extractSoftware(useragent string) string {
	preTrim := strings.Split(useragent, "(")[1]
	postTrim := strings.Split(preTrim, ")")[0]
	return postTrim
}
