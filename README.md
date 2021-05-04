# reviewdog-playground

## Requirements

- reviewdog

## golangci-lint

run on local

```sh
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.36.0 golangci-lint run ./... | reviewdog -f=golangci-lint -diff="git diff FETCH_HEAD"
```
