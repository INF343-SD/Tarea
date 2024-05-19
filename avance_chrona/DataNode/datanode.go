package main

import (
	pb "Lab4444/proto"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"
)

type KillingFloorStruct struct {
	pb.UnimplementedKillingFloorServer
}

func escribir(datos struct {
	nombre   string
	piso     int32
	decision int32
}) {
	piso_string := fmt.Sprintf("%d", datos.piso)
	// crea un archivo de texto con el nombre nombre
	archivo := datos.nombre + "_" + piso_string + ".txt"
	err := ioutil.WriteFile(archivo, []byte(""), 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Archivo creado")

	file, err := ioutil.ReadFile(archivo)

	if err != nil {
		fmt.Println(err)
		return
	}

	decision_string := fmt.Sprintf("%d", datos.decision)

	file = append(file, []byte(decision_string+"\n")...)
	err = ioutil.WriteFile(archivo, file, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Archivo actualizado")

}
func (s *KillingFloorStruct) NameNode_DN(ctx context.Context, req *pb.NameNode_DataNode) (res *pb.DataNode_NameNode, err error) {

	nombre := req.Nombre
	piso := req.Piso
	decision := req.Decision

	escribir(struct {
		nombre   string
		piso     int32
		decision int32
	}{nombre, piso, decision})

	return &pb.DataNode_NameNode{Ack: 1}, nil
}

func main() {
	mercenarios := &KillingFloorStruct{}

	lis, err := net.Listen("tcp", ":50055")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterKillingFloorServer(s, mercenarios)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
