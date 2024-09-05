package main

import (
	"fmt"
	"net/http"
)

const port = ":80"

func main() {
	fmt.Printf("running on port %s", port)
	http.ListenAndServe(port, nil)
}
