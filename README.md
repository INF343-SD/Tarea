# Integreantes:
- Bastián Navarrete         Rol: 202073510-7
- Felipe Rojas              Rol: 201873112-9

# Instrucciones de ejecución
Corre el Director y un datanode en una máquina
```
docker-compose -f docker-compose.yml -f compose-datanode.yml up --build
```

Corre el DoshBank y un datanode en una máquina
```
docker-compose -f compose-doshbank.yml -f compose-datanode.yml up --build
```

Corre el NameNode en una máquina
```
docker-compose -f compose-namenode.yml up --build
```

Corre los mercenarios y un datanode en una máquina
```
docker-compose -f compose-mercenarios.yml -f compose-datanode.yml up --build
```


Para compilar director.proto:

```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative director.proto
```