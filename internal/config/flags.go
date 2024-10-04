package config

import "flag"

type flagsValue struct {
	addr            string
	baseReturnedURL string
}

func (v *flagsValue) parseFlags() {
	flag.StringVar(&v.addr, "a", "localhost:8080", "the address to listen on for HTTP requests")
	flag.StringVar(&v.baseReturnedURL, "b", "http://localhost:8080", "base url returned in response when url is shorted")
	flag.Parse()
}

func newFlagsValue() flagsValue {
	f := flagsValue{}
	f.parseFlags()

	return f
}
