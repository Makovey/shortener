package config

import (
	"flag"
)

const (
	flagAddr            = "a"
	flagBaseURL         = "b"
	flagFileStoragePath = "f"
)

type flagsValue struct {
	addr            string
	baseReturnedURL string
	fileStoragePath string
}

func (v *flagsValue) parseFlagsIfNeeded() {
	if flag.Lookup(flagAddr) == nil {
		flag.StringVar(&v.addr, flagAddr, "", "the address to listen on for HTTP requests")
	}

	if flag.Lookup(flagBaseURL) == nil {
		flag.StringVar(&v.baseReturnedURL, flagBaseURL, "", "base url returned in response when url is shorted")
	}

	if flag.Lookup(flagFileStoragePath) == nil {
		flag.StringVar(&v.fileStoragePath, flagFileStoragePath, "", "file path for url storage")
	}

	flag.Parse()
}

func (v *flagsValue) getFlagsValue() {
	v.addr = flag.Lookup(flagAddr).Value.(flag.Getter).Get().(string)
	v.baseReturnedURL = flag.Lookup(flagBaseURL).Value.(flag.Getter).Get().(string)
	v.fileStoragePath = flag.Lookup(flagFileStoragePath).Value.(flag.Getter).Get().(string)
}

func newFlagsValue() flagsValue {
	var f flagsValue
	f.parseFlagsIfNeeded()

	if f.addr == "" || f.baseReturnedURL == "" || f.fileStoragePath == "" {
		f.getFlagsValue()
	}

	return f
}
