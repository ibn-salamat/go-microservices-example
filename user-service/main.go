package main

import (
	"fmt"
	"net/http"
)

const port = ":80"

func main() {
	fmt.Println("running on port %s", port)
	http.ListenAndServe(port, nil)
}
