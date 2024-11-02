package config

import (
	"flag"
)

const (
	flagAddr            = "a"
	flagBaseURL         = "b"
	flagFileStoragePath = "f"
	flagDatabaseDSN     = "d"
)

type flagsValue struct {
	addr            string
	baseReturnedURL string
	fileStoragePath string
	databaseDSN     string
}

func (v *flagsValue) parseFlagsIfNeeded() {
	if flag.Lookup(flagAddr) == nil {
		flag.StringVar(&v.addr, flagAddr, "", "the address to listen on for HTTP requests")
	} else {
		v.addr = flag.Lookup(flagAddr).Value.(flag.Getter).Get().(string)
	}

	if flag.Lookup(flagBaseURL) == nil {
		flag.StringVar(&v.baseReturnedURL, flagBaseURL, "", "base url returned in response when url is shorted")
	} else {
		v.baseReturnedURL = flag.Lookup(flagBaseURL).Value.(flag.Getter).Get().(string)
	}

	if flag.Lookup(flagFileStoragePath) == nil {
		flag.StringVar(&v.fileStoragePath, flagFileStoragePath, "", "disc path for url storage")
	} else {
		v.fileStoragePath = flag.Lookup(flagFileStoragePath).Value.(flag.Getter).Get().(string)
	}

	if flag.Lookup(flagDatabaseDSN) == nil {
		flag.StringVar(&v.databaseDSN, flagDatabaseDSN, "", "database DSN in format -> postgres://username:password@host:port/databaseName")
	}

	flag.Parse()
}

func newFlagsValue() flagsValue {
	var f flagsValue
	f.parseFlagsIfNeeded()

	return f
}
