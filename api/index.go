package api

import (
	"fmt"
	"net/http"

	"encoder"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	text := "Hi there, this is 3310 encoder program!"
	fmt.Fprintf(w, `
			<!DOCTYPE html>
			<html lang="en">
				<head>
					<meta charset="UTF-8">
					<meta name="viewport" content="width=device-width, initial-scale=1.0">
					<title>3310 Decoder in Golang!</title>
				</head>
				<body>
					<div style="display: grid; place-items: center; text-align: center;">
						<h1>Welcome to 3310 Encoder!</h1>
						<div>
							<strong>English</strong>
							<pre
								style="
									white-space: pre-wrap;
									word-break: keep-all;
									background: grey;
									color: white;
									border: 1px solid black;
									padding: 5px;
								"
							>%s</pre>
						</div>
						<div>
							<strong>3310 version</strong>
							<pre
								style="
									white-space: pre-wrap;
									word-break: keep-all;
									background: grey;
									color: white;
									border: 1px solid black;
									padding: 5px;
								"
							>%s</pre>
						</div>
					</div>
				</body>
			</html>
		`,
		text,
		encoder.Encode(text),
	)
}
