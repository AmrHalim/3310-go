package api

import (
	"fmt"
	"net/http"

	"encoder"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	text := "Hi there, this is 3310 encoder program!"
	fmt.Fprintf(w, "<h1>Welcome to 3310 Encoder!</h1><div><strong>English:</strong>%s</div><div><strong>3310 version:</strong>%s</div>",
		text,
		encoder.Encode(text),
	)
}
