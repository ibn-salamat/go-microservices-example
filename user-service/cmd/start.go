package cmd

import "users/internal/transport/http"

const port = ":80"

func Start() {
	http.Start(port)
}
