package main

import (
	"bytes"
	"strings"
	"testing"

	"gopkg.in/dnaeon/go-vcr.v3/cassette"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"
)

func patchGetIPResponse(i *cassette.Interaction) error {
	if strings.Contains(i.Request.URL, "/ip") {
		i.Response.Body = `{"origin": "127.0.0.1"}`
	}

	return nil
}

func runRecorder(cassette string, hooks []func(*cassette.Interaction) error) (*recorder.Recorder, error) {
	r, err := recorder.New(cassette)

	for _, hook := range hooks {
		r.AddHook(hook, recorder.AfterCaptureHook)
	}

	return r, err
}

func TestCustomizeResponse(t *testing.T) {
	hooks := []func(*cassette.Interaction) error{
		patchGetIPResponse,
	}

	r, err := runRecorder("testdata/get_ip", hooks)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop()

	var out bytes.Buffer
	err = ShowIPAddress(&out, r.GetDefaultClient())
	if err != nil {
		t.Error(err)
	}

	// assert output
	expected := "IP Address: 127.0.0.1"
	actual := out.String()

	if expected != actual {
		t.Errorf("%q != %q", expected, actual)
	}
}
