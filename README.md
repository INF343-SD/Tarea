# Integreantes:
- Bastián Navarrete         Rol: 202073510-7
- Felipe Rojas              Rol: 201873112-9

## Consideraciones:
- Máquina 033 está el director y un DataNode
- Máquina 034 DoshBank y un DataNode
- Máquina 035 NameNode
- Máquina 034 Mercenarios y un DataNode

Debe ejecutar los comandos de abajo considerando estos puntos, de otro modo la tarea no funciona

# Instrucciones de ejecución
Corre el Director
```
sudo docker compose up
```

Corre el DoshBank
```
cd DoshBank
sudo docker compose up
```

Corre el NameNode
```
cd NameNode
sudo docker compose up
```

Corre los mercenarios
```
cd Mercenarios
sudo docker run -it mercenarios-server
```

Corre un DataNode
```
cd DataNode
sudo docker compose up
```

Para compilar director.proto:

```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative director.proto
```
