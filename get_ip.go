package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ipBody struct {
	Origin string `json:"origin"`
}

func ShowIPAddress(w io.Writer, client *http.Client) error {
	url := "https://httpbin.org/ip"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var b ipBody
	err = json.NewDecoder(res.Body).Decode(&b)
	if err != nil {
		return err
	}

	fmt.Fprintf(w, "IP Address: %s", b.Origin)
	return nil
}
