## Integrantes:
- Bastian Navarrete     202073510-7
- Felipe Rojas          201873112-9

## Instrucciones de uso:
### Para el servidor (Tierra):
```
docker compose up --build --force-recreate
```
### Para el cliente:
```
cd cliente
go run client.go
```


To compile ammo.proto:

    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/ammo.proto