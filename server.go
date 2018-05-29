package lazyhackergo

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, "<meta name=\"go-import\" content=\"%s%s git https://github.com/lazyhacker%s\">", r.Host, r.URL.Path, r.URL.Path)
}

func init() {
	http.HandleFunc("/", handler)
}
