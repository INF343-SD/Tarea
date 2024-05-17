package main

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"

	pb "github.com/ChronaNavarrete/2024/tree/main/2024-01/Distribuidos/Lab3/proto"
	"google.golang.org/grpc"
)

type AmmoStruct struct {
	pb.UnimplementedAmmoServer
	av_at int32
	av_mp int32
	mutex sync.Mutex
}

func (s *AmmoStruct) updateMunitions() {
	for {
		time.Sleep(5 * time.Second)

		s.mutex.Lock()
		s.av_at += 10
		s.av_mp += 5

		if s.av_at > 50 {
			s.av_at = 50
		}

		if s.av_mp > 20 {
			s.av_mp = 20
		}

		fmt.Println("=========== ACTUALIZACION DE MUNICION =========")
		fmt.Printf("AT: %d, MP: %d\n", s.av_at, s.av_mp)

		s.mutex.Unlock()
	}
}

func (s *AmmoStruct) ReqAmmo(ctx context.Context, req *pb.AmmoReq) (*pb.AmmoResp, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.av_at < req.At || s.av_mp < req.Np {
		fmt.Printf("Recepción de solicitud desde equipo %s, %s AT y %s MP -- DENEGADA -- AT EN SISTEMA: %s ; MP EN SISTEMA: %s\n", strconv.Itoa(int(req.GetTeam())), strconv.Itoa(int(req.GetAt())), strconv.Itoa(int(req.GetNp())), strconv.Itoa(int(s.av_at)), strconv.Itoa(int(s.av_mp)))
		return &pb.AmmoResp{App: false}, nil
	}
	fmt.Printf("Recepción de solicitud desde equipo %s, %s AT y %s MP -- APROBADA -- AT EN SISTEMA: %s ; MP EN SISTEMA: %s\n", strconv.Itoa(int(req.GetTeam())), strconv.Itoa(int(req.GetAt())), strconv.Itoa(int(req.GetNp())), strconv.Itoa(int(s.av_at)), strconv.Itoa(int(s.av_mp)))
	s.av_at -= req.At
	s.av_mp -= req.Np
	return &pb.AmmoResp{App: true, AvAt: s.av_at, AvNp: s.av_mp}, nil
}

func main() {
	ammo := &AmmoStruct{
		av_at: 0,
		av_mp: 0,
	}
	go ammo.updateMunitions()

	port := ":50051"

	conn, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("No se pudo crear la conexión TCP: " + err.Error())
		return
	}
	defer conn.Close()

	fmt.Println("Servidor escuchando en el puerto: ", port)

	fmt.Println("=========== MUNICION INICIAL =========")
	fmt.Printf("AT: %d, MP: %d\n", ammo.av_at, ammo.av_mp)

	sv := grpc.NewServer()
	pb.RegisterAmmoServer(sv, ammo)

	if err = sv.Serve(conn); err != nil {
		fmt.Println("Error con el levantamiento del servidor: " + err.Error())
	}

}
