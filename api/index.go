package handler

import (
	"fmt"
	"github.com/AmrHalim/encoder"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>\n<p>%s</p>", encoder.Encode("hi"))
}
