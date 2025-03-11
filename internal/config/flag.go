package config

import (
	"flag"
)

const (
	flagAddr            = "a"
	flagBaseURL         = "b"
	flagFileStoragePath = "f"
	flagDatabaseDSN     = "d"
	flagEnableHTTPS     = "s"
	flagConfigFile      = "c"
	flagTrustedSubnet   = "t"
	flagGRPCPort        = "g"
)

type flagsValue struct {
	addr            string
	baseReturnedURL string
	fileStoragePath string
	databaseDSN     string
	enableHTTPS     bool
	configFilePath  string
	trustedSubnet   string
	grpcPort        string
}

func newFlagsValue() flagsValue {
	var f flagsValue
	f.parseFlagsIfNeeded()

	return f
}

func (v *flagsValue) parseFlagsIfNeeded() {
	registerFlag(
		flagAddr,
		"the address to listen on for HTTP requests, in format [host:port]",
		&v.addr,
	)

	registerFlag(
		flagBaseURL,
		"base full url returned in response when url is shorted, in format [protocol://host:port]",
		&v.baseReturnedURL,
	)

	registerFlag(
		flagFileStoragePath,
		"disc path for url storage, in format [./filename.format]",
		&v.fileStoragePath,
	)

	registerFlag(
		flagDatabaseDSN,
		"database DSN, in format [postgres://username:password@host:port/dbName]",
		&v.databaseDSN,
	)

	registerBoolFlag(
		flagEnableHTTPS,
		"enable HTTPS for server, in format [-s]",
		&v.enableHTTPS,
	)

	registerFlag(
		flagConfigFile,
		"the name of config file, in format [filename]",
		&v.configFilePath,
	)

	registerFlag(
		flagTrustedSubnet,
		"trusted subnet for handler [filename]",
		&v.trustedSubnet,
	)

	registerFlag(
		flagGRPCPort,
		"gRPC port for launch server, in format [:port]",
		&v.grpcPort,
	)

	flag.Parse()
}

func registerFlag(name, usage string, target *string) {
	if flag.Lookup(name) == nil {
		flag.StringVar(target, name, "", usage)
	} else {
		*target = flag.Lookup(name).Value.(flag.Getter).Get().(string)
	}
}

func registerBoolFlag(name, usage string, target *bool) {
	if flag.Lookup(name) == nil {
		flag.BoolVar(target, name, false, usage)
	} else {
		*target = flag.Lookup(name).Value.(flag.Getter).Get().(bool)
	}
}
