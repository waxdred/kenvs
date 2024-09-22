package kenvs

import (
	"encoding/base64"
	"os"

	parse "github.com/waxdred/kenvs/Parse"
	"gopkg.in/yaml.v2"
)

func (k *Kenvs) parseSecretFile() error {
	b, err := os.ReadFile(k.secretFile)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(b, &k.data)
	if err != nil {
		return err
	}
	return nil
}

func New(envFile string, secretFile string, encode bool) *Kenvs {
	return &Kenvs{
		envFile:    envFile,
		encode:     encode,
		secretFile: secretFile,
	}
}

func (k *Kenvs) UpdateSecret() error {
	if k.secretFile == "" {
		k.secretFile = "secret.yaml"
	}
	if k.encode {
		k.data.Data = k.env
		k.data.StringData = nil
	} else {
		k.data.StringData = k.env
		k.data.Data = nil
	}
	secretByte, err := yaml.Marshal(&k.data)
	if err != nil {
		return err
	}
	if err := os.WriteFile(k.secretFile, secretByte, 0644); err != nil {
		return err
	}
	return nil
}

func (k *Kenvs) Run() error {
	var err error
	k.env, err = parse.ParseEnvFile(k.envFile)
	if err != nil {
		return err
	}
	if k.encode {
		for key, e := range k.env {
			k.env[key] = base64.StdEncoding.EncodeToString([]byte(e))
		}
	}
	if k.secretFile != "" {
		err := k.parseSecretFile()
		if err != nil {
			return err
		}
	}
	return k.UpdateSecret()
}
