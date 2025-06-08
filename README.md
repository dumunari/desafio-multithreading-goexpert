# desafio-multithreading-goexpert

## Server
Para subir o servidor, na raiz do projeto, você pode:
1) Executar o binário cep-server que se encontra na pasta bin, com o comando abaixo:
```
./bin/cep-server
```

2) Ou então, rodar o arquivo main.go, seguindo os passos abaixos.
```
go mod tidy
go run cmd/main.go
```

## Requisições
Para utilizar o servidor: 
1) Caso possua o VSCode com a extensão Rest Client, está disponível o arquivo cep.http na pasta test.
2) Também é possível realizar chamadas via cURL, como no exemplo abaixo:
```
curl http://localhost:8080/cep/01153000
```