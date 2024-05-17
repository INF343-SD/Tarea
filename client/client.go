package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	pb "github.com/ChronaNavarrete/2024/tree/main/2024-01/Distribuidos/Lab3/proto"
	"google.golang.org/grpc"
)

func sendRequests(sv pb.AmmoClient, wg *sync.WaitGroup, teams struct {
	ID      int32
	AtCount int32
	MpCount int32
}) {
	defer wg.Done()
	for {
		time.Sleep(10 * time.Second)

		respond, err := sv.ReqAmmo(context.Background(), &pb.AmmoReq{
			Team: teams.ID,
			At:   teams.AtCount,
			Np:   teams.MpCount,
		})
		if err != nil {
			fmt.Println("No se pudo realizar la solicitud: " + err.Error())
			continue
		}
		if respond.App {
			fmt.Printf("Equipo %d solicitando %d AT y %d MP; Resolución: -- APROBADA -- ; Conquista Exitosa!, cerrando comunicación\n", teams.ID, teams.AtCount, teams.MpCount)
			break
		} else {
			fmt.Printf("Equipo %d solicitando %d AT y %d MP; Resolución: -- DENEGADA -- ; Reintentando en 3 segundos...\n", teams.ID, teams.AtCount, teams.MpCount)
			time.Sleep(3 * time.Second)
		}
	}
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())	# Cambiar localhost por la direccion del servidor
	if err != nil {
		fmt.Println("No se pudo contactar con el servidor: " + err.Error())
	}
	defer conn.Close()

	fmt.Println("Conexión establecida con la Tierra")
	sv := pb.NewAmmoClient(conn)

	teams := []struct {
		ID      int32
		AtCount int32
		MpCount int32
	}{
		{ID: 1, AtCount: int32(20 + rand.Intn(11)), MpCount: int32(10 + rand.Intn(6))},
		{ID: 2, AtCount: int32(20 + rand.Intn(11)), MpCount: int32(10 + rand.Intn(6))},
		{ID: 3, AtCount: int32(20 + rand.Intn(11)), MpCount: int32(10 + rand.Intn(6))},
		{ID: 4, AtCount: int32(20 + rand.Intn(11)), MpCount: int32(10 + rand.Intn(6))},
	}

	var wg sync.WaitGroup

	for _, team := range teams {
		wg.Add(1)
		go sendRequests(sv, &wg, team)
	}
	wg.Wait()
}
