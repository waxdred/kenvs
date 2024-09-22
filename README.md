# kenvs (Kubernetes ENV Secret)

kenvs is a command-line tool that simplifies the process of creating and updating Kubernetes Secrets from .env files. It allows for easy conversion between .env files and Kubernetes Secret YAML, with options for base64 encoding.

## Features

- Read .env files and convert them to Kubernetes Secret YAML
- Update existing Kubernetes Secret YAML files with new .env data
- Support for both `stringData` (plain text) and `data` (base64 encoded) Secret formats
- Flexible configuration options for customizing Secret metadata

## Installation

To install kenvs, make sure you have Go installed on your system, then run:

```bash
go get github.com/waxdred/kenvs
```

## Usage

Basic usage:

```bash
kenvs [options]
```

Options:

- `-env string`: Path to the .env file (default: ".env")
- `-encode`: Encode secret in base64 (uses `data` field instead of `stringData`)
- `-secret string`: Path to the existing secret YAML file to update (optional)

### Examples

1. Generate a new Secret from a .env file:

   ```bash
   kenvs -env myapp.env
   ```

2. Update an existing Secret with new .env data:

   ```bash
   kenvs -env myapp.env -secret existing-secret.yaml
   ```

3. Generate a Secret with base64 encoded values:
   ```bash
   kenvs -env myapp.env -encode
   ```

## Output

kenvs will generate or update a Kubernetes Secret YAML file. By default, it uses the `stringData` field for plain text values. When the `-encode` flag is used, it uses the `data` field with base64 encoded values.

Example output (without encoding):

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: my-secret
  namespace: default
type: Opaque
stringData:
  DB_PASSWORD: mysecretpassword
  API_KEY: abcdef123456
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Disclaimer

This tool is provided as-is, without any guarantees or warranty. The authors are not responsible for any damage or data loss that may occur from its use.
