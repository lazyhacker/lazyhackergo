package lazyhackergo

import (
	"fmt"
	"net/http"
)

func init() {
	http.Handle("/.well-known/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<meta name=\"go-import\" content=\"%s%s git https://github.com/lazyhacker%s\">", r.Host, r.URL.Path, r.URL.Path)
}
