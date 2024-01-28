package api

import (
	"fmt"
	"net/http"

	"encoder"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	text := "Hi there, this is 3310 encoder program!"
	fmt.Fprintf(w, "<h1>Welcome to 3310 Encoder!</h1>\n<p>English: %s\n3310 version: %s</p>",
		text,
		encoder.Encode(text),
	)
}
