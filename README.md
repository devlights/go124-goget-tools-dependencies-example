# go124-goget-tools-dependencies-example

Go 1.24で導入された go get -tool によるツール依存関係インストールのサンプルです。

```go get -tool```でインストールしたツールは ```go tool [ライブラリ名]``` で実行出来ます。

## Run

```yaml
# https://taskfile.dev

version: '3'

tasks:
  default:
    cmds:
      - go generate
      - go vet
      - go tool goimports -w main.go
      - go build
      - ./app
```

```sh
$ task
task: [default] go generate
task: [default] go vet
task: [default] go tool goimports -w main.go
task: [default] go build
task: [default] ./app
  Java(1)
Golang(0)
CSharp(2)
```