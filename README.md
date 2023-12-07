# secrets-init

This is a simple CLI that reads secrets from Secrets Manager. It's a perfect "init" container in Kubernetes, it can create a file on a shared volume so the other containers can use that file. secrets-init can filter one or more secrets by name using a regular expression, it also parses the secret content as plain text or json.

## CLI

```sh
./secrets-init sync \
    --provider YOUR_CLOUD_PROVIDER \
    --project YOUR_PROJECT_ID \
    --filter YOUR_FILTER \
    --data-parser json
```

### Example

Given a secret called `myapp` with the content below:

```json
{
  "username": "root",
  "password": "s3cr3t",
  "host": "localhost",
  "port": "5432"
}

```

Running secrets-init with the flags:

```bash
./secrets-init sync \
    --provider gcp \
    --project myproject \
    --filter=^myapp*" \
    --data-parser json
```

Output:

```bash
export MYAPP_PASSWORD="s3cr3t"
export MYAPP_HOST="localhost"
export MYAPP_PORT="5432"
export MYAPP_USERNAME="root"
```
## Init container

Check the examples directory

- [GCP](https://github.com/leocomelli/secrets-init/blob/main/examples/gcp.yml)
- [AWS](https://github.com/leocomelli/secrets-init/blob/main/examples/aws.yml)


## Providers

- [x] Google Cloud Platform
- [x] AWS
- [ ] Azure

## Filter

Use the flag `--filter` to filter one or more secrets, a regular expression should be provided ([regexp/syntax](https://pkg.go.dev/regexp/syntax)).

## Parser

Use the flag `--data-parser` to parse the secret content. There are two predefined parsers, the default is `plaintext` the other one is `json`. Both parses are associated with a template to render the output.

* **plaintext:** `export {{ .Name  | ToUpper }}="{{ .Data }}`, where `Name` is the secret name and `Data` is the full content.
* **json:** `export {{ .Name  | ToUpper }}_{{ .ContentKey | ToUpper }}="{{ .ContentValue }}`, where `Name` is the secret name, `ContentKey`/`ContentValue` are the key and value of each json property. 

But when necessary, the template can be reset ([text/template](https://pkg.go.dev/text/template)). Use the flag `--template`, for example, to generate an output file in key/value format.

```bash
--template {{ .Name | ToLower }}_{{.ContentKey | ToLower }}={{ .ContentValue }}
```

## Output

Use the `--output` to write output file to a specific path, `stdout` if it is empty.

