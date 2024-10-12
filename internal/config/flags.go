package config

import (
	"flag"
)

const (
	flagAddr    = "a"
	flagBaseURL = "b"
)

type flagsValue struct {
	addr            string
	baseReturnedURL string
}

func (v *flagsValue) parseFlagsIfNeeded() {
	if flag.Lookup(flagAddr) == nil {
		flag.StringVar(&v.addr, flagAddr, "localhost:8080", "the address to listen on for HTTP requests")
	}

	if flag.Lookup(flagBaseURL) == nil {
		flag.StringVar(&v.baseReturnedURL, flagBaseURL, "http://localhost:8080", "base url returned in response when url is shorted")
	}

	flag.Parse()
}

func (v *flagsValue) getFlagsValue() {
	v.addr = flag.Lookup(flagAddr).Value.(flag.Getter).Get().(string)
	v.baseReturnedURL = flag.Lookup(flagBaseURL).Value.(flag.Getter).Get().(string)
}

func newFlagsValue() flagsValue {
	f := flagsValue{}
	f.parseFlagsIfNeeded()

	if f.addr == "" || f.baseReturnedURL == "" {
		f.getFlagsValue()
	}

	return f
}
