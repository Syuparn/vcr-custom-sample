package main

import (
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}
	ShowIPAddress(os.Stdout, client)
}
