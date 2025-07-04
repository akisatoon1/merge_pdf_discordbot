package config

import "flag"

type flags struct {
	Emode bool
}

func loadFlags() *flags {
	f := &flags{}

	flag.BoolVar(&f.Emode, "e", false, "Enable eMode to load environment variables from .env file")
	flag.Parse()

	return f
}
