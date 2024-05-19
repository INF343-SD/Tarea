package main

import (
	pb "Lab4444/proto"
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"

	"google.golang.org/grpc"
)

type KillingFloorStruct struct {
	pb.UnimplementedKillingFloorServer
}

func escribir(datos struct {
	nombre  string
	piso    int32
	maquina string
}) {
	//
	archivo := "registro.txt"
	file, err := ioutil.ReadFile(archivo)

	if err != nil {
		fmt.Println(err)
		return
	}

	piso_string := fmt.Sprintf("%d", datos.piso)

	file = append(file, []byte(datos.nombre+" Piso_"+piso_string+" "+datos.maquina+"\n")...)
	err = ioutil.WriteFile(archivo, file, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Archivo actualizado")

}

func enviarDecision_DataNode(ctx context.Context, decision struct {
	nombre    string
	piso      int32
	decision  int32
	direccion string
}) {

	conn, err := grpc.Dial(decision.direccion, grpc.WithInsecure())
	if err != nil {
		fmt.Println("No se pudo contactar con el servidor: " + err.Error())
	}
	defer conn.Close()

	sv := pb.NewKillingFloorClient(conn)
	rep, err := sv.NameNode_DN(ctx, &pb.NameNode_DataNode{Nombre: decision.nombre, Piso: decision.piso, Decision: decision.decision})

	if err != nil {
		fmt.Println("No se pudo contactar con el servidor: " + err.Error())
	}

	if rep.Ack != 1 {
		fmt.Println("No se pudo contactar con el servidor: " + err.Error())
	}
}

func (s *KillingFloorStruct) Dir_NN_Dec(ctx context.Context, req *pb.Dir_NameNode) (*pb.NameNode_Dir, error) {
	nombre := req.Nombre
	piso := req.Piso
	decision := req.Decision

	// se escoge una de las 3 direcciones al azar y luego se conecta a la direccion escogida
	indice := rand.Intn(3)
	var direccion string
	if indice == 0 {
		direccion = "10.35.169.43:50055"
	} else if indice == 1 {
		direccion = "10.35.169.44:50055"
	} else {
		direccion = "10.35.169.46:50055"
	}

	escribir(struct {
		nombre  string
		piso    int32
		maquina string
	}{nombre: nombre, piso: piso, maquina: direccion})

	enviarDecision_DataNode(ctx, struct {
		nombre    string
		piso      int32
		decision  int32
		direccion string
	}{nombre, piso, decision, direccion})

	return &pb.NameNode_Dir{Ack: 1}, nil
}

func main() {
	conn, err := net.Listen("tcp", ":50053")

	if err != nil {
		fmt.Println("No se pudo crear la conexión TCP: " + err.Error())
		return
	}
	defer conn.Close()
	fmt.Println("Servidor escuchando en el puerto: 50053")
	sv := grpc.NewServer()
	pb.RegisterKillingFloorServer(sv, &KillingFloorStruct{})
	if err = sv.Serve(conn); err != nil {
		fmt.Println("Error con el levantamiento del servidor: " + err.Error())
	}

}
