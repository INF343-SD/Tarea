package main

import (
	pb "Lab4444/proto"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

var montos_mercenarios = map[string]int{
	"Reverend Alberts":   100000000,
	"Tom Banner":         100000000,
	"PC Rob Briar":       100000000,
	"Classic Briar":      100000000,
	"Mr. Foster":         100000000,
	"Ana Larive":         100000000,
	"Lt. Bill Masterson": 100000000,
	"User":               100000000,
}

type KillingFloorStruct struct {
	pb.UnimplementedKillingFloorServer
}

func (s *KillingFloorStruct) Pet_Monto(ctx context.Context, req *pb.Dir_DBank) (*pb.DBank_Dir, error) {
	nombre := req.Nombre
	monto, ok := montos_mercenarios[nombre]
	if !ok {
		return nil, fmt.Errorf("el mercenario %s no encontrado", nombre)
	}
	return &pb.DBank_Dir{Nombre: nombre, Monto: int32(monto)}, nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func consumeMessages() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
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

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		txt_s := ""
		for d := range msgs {
			data := string(d.Body)
			parts := strings.Split(data, ",")
			name := parts[0]
			floor := parts[1] // Esto es un string, debes convertirlo a int si es necesario.
			log.Printf("Received death notice for %s on floor %s", name, floor)
			//montos_mercenarios[name] = 0
			floor_s := fmt.Sprint(floor)
			montos_s := fmt.Sprint(montos_mercenarios[name])
			txt_s = txt_s + name + " Piso_" + floor_s + " " + montos_s + "\n"
			b := []byte(txt_s)
			err := ioutil.WriteFile("mercenarios_eliminados.txt", b, 0644)
			if err != nil {
				log.Fatal(err)
			}
			for k, v := range montos_mercenarios {
				if k != name {
					montos_mercenarios[k] = v + 100000000
				}
			}
			log.Printf("Monto actualizado para los mercenarios vivos")
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func main() {
	port := ":50052"

	conn, err := net.Listen("tcp", port)

	if err != nil {
		fmt.Println("No se pudo crear la conexión TCP: " + err.Error())
		return
	}
	defer conn.Close()
	fmt.Println("Servidor escuchando en el puerto: ", port)
	sv := grpc.NewServer()
	pb.RegisterKillingFloorServer(sv, &KillingFloorStruct{})
	if err = sv.Serve(conn); err != nil {
		fmt.Println("Error con el levantamiento del servidor: " + err.Error())
	}

	consumeMessages()

}
