# grpc-stream-bi-direcional-example

# Instalação

```
sudo apt install protobuf-compiler 
go mod init github.com/<seu_user>/<repo_name>
```

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go get google.golang.org/grpc
```

# Executar
```
go run cmd/server/server.go
go run cmd/server/client.go
```

# Client Evans
[Repositório Github](https://github.com/ktr0731/evans#from-github-releases)

```
tar -zxvf evans_linux_amd64.tar.gz
```
> Mova o arquivo para a PATH

```
evans -r repl --host localhost --port 50051
```
![Exemplo do Evans](../../assets/evans-example.png)