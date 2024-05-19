package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	pb "github.com/ChronaNavarrete/2024/tree/main/2024-01/Distribuidos/Lab3/proto"
	"google.golang.org/grpc"
)

type MessageStruct struct {
	pb.UnimplementedMercenaryTaskServiceServer
	mutex    sync.Mutex
	codigo   []string
	piso     []string
	decision []string
}

type Node struct {
	Name    string
	Address string
}

func (s *MessageStruct) RecieveTask(ctx context.Context, req *pb.TaskReq) (*pb.TaskResp, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return &pb.TaskResp{
		Codigo:   s.codigo,
		Piso:     s.piso,
		Decision: s.decision,
	}, nil
}

func (s *MessageStruct) SendTask(ctx context.Context, req *pb.TaskReq) (*pb.TaskResp, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.codigo = append(s.codigo, req.GetCodigo())
	s.piso = append(s.piso, req.GetPiso())
	s.decision = append(s.decision, req.GetDecision())

	return &pb.TaskResp{}, nil
}

func main() {

	lis, err := net.Listen("tcp", "500051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	s := grpc.NewServer()
	pb.RegisterMercenaryTaskServiceServer(s, &MessageStruct{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	rand.Seed(time.Now().UnixNano())

	// Lista de nodos
	nodes := []Node{
		{Name: "Nodo1", Address: "10.35.169.43:5432"},
		{Name: "Nodo2", Address: "10.35.169.44:5432"},
		{Name: "Nodo3", Address: "10.35.169.45:5432"},
	}

	// Recibir información (simulado aquí)
	info := "Información importante"

	randomIndex := rand.Intn(len(nodes))
	selectedNode := nodes[randomIndex]

	fmt.Printf("Enviando información a %s en %s\n", selectedNode.Name, selectedNode.Address)

}
