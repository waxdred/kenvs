package parse

import (
	"bufio"
	"os"
	"strings"
)

func ParseEnvFile(envFile string) (map[string]string, error) {
	var env = make(map[string]string)
	tmpKey := ""
	file, err := os.Open(envFile)
	if err != nil {
		return env, nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			env[tmpKey] = env[tmpKey] + line
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		value = strings.Trim(value, `"'`)
		env[key] = value
		tmpKey = key
	}
	if err := scanner.Err(); err != nil {
		return env, err
	}

	return env, nil
}
