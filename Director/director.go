package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"sync"

	pb "Lab4444/proto"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

var monto_solicitado int32

type KillingFloorStruct struct {
	pb.UnimplementedKillingFloorServer
	mutex sync.Mutex
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func publishMessage(nombre string, piso int32) {
	conn, err := amqp.Dial("amqp://guest:guest@10.35.169.44:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"mercenario_muerto", // name
		false,               // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	failOnError(err, "Failed to declare a queue")

	message := fmt.Sprintf("%s,%d", nombre, piso)

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	failOnError(err, "Failed to publish a message")
}

func pedirMonto(ctx context.Context, nombre string) int32 {
	conn2, err := grpc.Dial("10.35.169.44:50052", grpc.WithInsecure())
	if err != nil {
		fmt.Println("No se pudo contactar con el servidor: " + err.Error())
	}
	defer conn2.Close()
	sv2 := pb.NewKillingFloorClient(conn2)
	resp, err := sv2.Pet_Monto(ctx, &pb.Dir_DBank{Nombre: nombre})
	if err != nil {
		fmt.Println("No se pudo realizar la solicitud: \n" + err.Error())
	}

	monto_solicitado = resp.Monto
	return monto_solicitado
}

func (s *KillingFloorStruct) Merc_Dir_Monto(ctx context.Context, req *pb.Merc_Dir2) (*pb.Dir_Merc2, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	monto_solicitado := pedirMonto(ctx, req.Nombre)
	return &pb.Dir_Merc2{Nombre: req.Nombre, Monto: monto_solicitado}, nil
}

func enviarDecision(ctx context.Context, decision struct {
	nombre   string
	piso     int32
	decision int32
}) {
	conn3, err := grpc.Dial("10.35.169.45:50053", grpc.WithInsecure())
	if err != nil {
		fmt.Println("No se pudo contactar con el servidor: " + err.Error())
	}
	defer conn3.Close()
	sv3 := pb.NewKillingFloorClient(conn3)
	rep, err := sv3.Dir_NN_Dec(ctx, &pb.Dir_NameNode{Nombre: decision.nombre, Piso: decision.piso, Decision: decision.decision})
	if err != nil {
		fmt.Println("No se pudo realizar la solicitud: \n" + err.Error())
	}

	if rep.Ack != 1 {
		fmt.Println("El mercenario ", decision.nombre, " ha llegado al piso ", decision.piso, " con la decisión ", decision.decision)
	}
}

func (s *KillingFloorStruct) Decision(ctx context.Context, req *pb.Merc_Dir) (*pb.Dir_Merc, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	datos := struct {
		nombre   string
		piso     int32
		decision int32
	}{
		nombre:   req.Nombre,
		piso:     req.Piso,
		decision: req.Decision,
	}
	enviarDecision(ctx, datos)

	if req.Piso == 1 {
		x := int32(rand.Intn(101) + 1)
		fmt.Println("Probabilidad de muerte: ", x)
		if req.Decision < x {
			fmt.Printf("%s ha muerto en el Piso %d! :(. Su decisión fue: %d\n", req.Nombre, req.Piso, req.Decision)
			publishMessage(req.Nombre, req.Piso)
			return &pb.Dir_Merc{Vivo: false}, nil
		}
	} else if req.Piso == 2 {
		x := int32(rand.Intn(2))
		if req.Decision != x {
			fmt.Printf("%s ha muerto en el Piso %d! :(. Su decisión fue: %d\n", req.Nombre, req.Piso, req.Decision)
			publishMessage(req.Nombre, req.Piso)
			return &pb.Dir_Merc{Vivo: false}, nil
		}
	} else if req.Piso == 3 {
		x := int32(rand.Intn(16) + 1)
		if req.Decision == x {

			publishMessage(req.Nombre, req.Piso)
			actual := req.Aciertos + 1
			if req.Aciertos == 3 {
				fmt.Printf("%s ha sobrevivido en el Piso %d. Su decisión fue: %d\n", req.Nombre, req.Piso, req.Decision)
				return &pb.Dir_Merc{Vivo: true}, nil
			}

			return &pb.Dir_Merc{Aciertos: actual}, nil
		}
		publishMessage(req.Nombre, req.Piso)
		return &pb.Dir_Merc{Vivo: false}, nil
	}

	fmt.Printf("%s ha sobrevivido en el Piso %d. Su decisión fue: %d\n", req.Nombre, req.Piso, req.Decision)
	return &pb.Dir_Merc{Vivo: true}, nil
}

func main() {
	mercenarios := &KillingFloorStruct{}

	port := ":50051"

	conn, err := net.Listen("tcp", port)

	if err != nil {
		fmt.Println("No se pudo crear la conexión TCP: " + err.Error())
		return
	}
	defer conn.Close()
	fmt.Println("Servidor escuchando en el puerto: ", port)
	sv := grpc.NewServer()
	pb.RegisterKillingFloorServer(sv, mercenarios)

	if err = sv.Serve(conn); err != nil {
		fmt.Println("Error con el levantamiento del servidor: " + err.Error())
	}

}
