package kenvs

type Kenvs struct {
	envFile    string
	secretFile string
	encode     bool
	data       DataSecret
	env        map[string]string
}

type DataSecret struct {
	ApiVersion string            `yaml:"apiVersion"`
	Kind       string            `yaml:"kind"`
	Metadata   map[string]string `yaml:"metadata"`
	Type       string            `yaml:"type"`
	StringData map[string]string `yaml:"stringData,omitempty"`
	Data       map[string]string `yaml:"data,omitempty"`
}
