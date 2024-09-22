package main

import (
	"flag"
	"fmt"

	"github.com/waxdred/kenvs/kenvs"
)

var (
	envFile    = flag.String("env", ".env", "path to env file")
	encode     = flag.Bool("encode", false, "encode secret in base64")
	secretFile = flag.String("secret", "", "secret to encrypt env file")
)

func main() {
	flag.Parse()
	k := kenvs.New(*envFile, *secretFile, *encode)
	err := k.Run()
	if err != nil {
		fmt.Errorf(err.Error())
	}
}
