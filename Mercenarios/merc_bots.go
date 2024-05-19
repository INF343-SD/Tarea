package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"

	pb "Lab4444/proto"

	"google.golang.org/grpc"
)

func enviarDecision(sv pb.KillingFloorClient, wg *sync.WaitGroup, mercenarios struct {
	nombre   string
	piso     int32
	decision int32
	vivo     bool
	aciertos int32
}) {
	defer wg.Done()
	for {
		if !mercenarios.vivo {
			fmt.Printf("%s murió en el piso %d! :(\n", mercenarios.nombre, mercenarios.piso)
		}

		if mercenarios.piso == 1 {
			if mercenarios.nombre == "User" {
				println("Escoge un arma: ")
				println("1) Escopeta")
				println("2) Rifle automático")
				println("3) Puños eléctricos")
				fmt.Scanln(&mercenarios.decision)
			}

			if mercenarios.decision == 1 {
				mercenarios.decision = int32(rand.Intn(101) + 1)
			} else if mercenarios.decision == 2 {
				x := int32(rand.Intn(101) + 1)
				y := int32(rand.Intn(101) + 1)
				for x > y {
					x = int32(rand.Intn(101) + 1)
					y = int32(rand.Intn(101) + 1)
				}
				mercenarios.decision = y - x
			} else if mercenarios.decision == 3 {
				y := int32(rand.Intn(101) + 1)
				mercenarios.decision = 100 - y
			}
		}

		respond, err := sv.Decision(context.Background(), &pb.Merc_Dir{
			Nombre:   mercenarios.nombre,
			Piso:     mercenarios.piso,
			Decision: mercenarios.decision,
			Vivo:     mercenarios.vivo,
			Aciertos: mercenarios.aciertos,
		})
		if err != nil {
			fmt.Println("No se pudo realizar la solicitud: \n" + err.Error())
			continue
		}

		if respond.Vivo {
			fmt.Printf("%s sobrevivió al Piso %d! :D\n", mercenarios.nombre, mercenarios.piso)
			if mercenarios.piso == 3 {

				if mercenarios.nombre == "" {
					fmt.Println("Llegaste al tercer nivel")
				}

				for i := 0; i < 5; i++ {
					if respond.Vivo {
						if mercenarios.nombre == "User" {
							println("Escoge un número entre 1 y 15")
							fmt.Scanln(&mercenarios.decision)
						} else {
							mercenarios.decision = int32(rand.Intn(16) + 1)
						}

						respond, err = sv.Decision(context.Background(), &pb.Merc_Dir{
							Nombre:   mercenarios.nombre,
							Piso:     mercenarios.piso,
							Decision: mercenarios.decision,
							Vivo:     mercenarios.vivo,
							Aciertos: mercenarios.aciertos,
						})
						if err != nil {
							fmt.Println("No se pudo realizar la solicitud: \n" + err.Error())
							continue
						}

					}
				}

			} else if mercenarios.piso == 1 {
				mercenarios.piso += 1
				if mercenarios.nombre == "User" {
					println("Estás en el segundo piso, el camino se divide en dos. ¿Por cuál pasillo irás?")
					println("1) Pasillo A")
					println("2) Pasillo B")
					fmt.Scanln(&mercenarios.decision)
				}

				mercenarios.decision = int32(rand.Intn(2))
			} else if mercenarios.piso == 2 {
				mercenarios.piso += 1
				mercenarios.decision = int32(rand.Intn(16) + 1)
			}

		} else {
			fmt.Printf("%s murió en el Piso %d! :(\n", mercenarios.nombre, mercenarios.piso)
			break
		}
	}
}

func solicitarMonto(sv pb.KillingFloorClient, nombre string) {
	respond, err := sv.Merc_Dir_Monto(context.Background(), &pb.Merc_Dir2{
		Nombre: nombre,
	})
	if err != nil {
		fmt.Println("No se pudo realizar la solicitud: \n" + err.Error())
	}
	fmt.Printf("El monto acumulado actual de %s es de: %d\n", nombre, respond.Monto)
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("No se pudo contactar con el servidor: " + err.Error())
	}
	defer conn.Close()

	fmt.Println("Conexión establecida")
	sv := pb.NewKillingFloorClient(conn)
	var ver_monto int
	var decision_inicial int
	fmt.Println("-----------------------------------------------------\n¡Bienvenido usuario! Antes de comenzar el piso 1 ¿te gustaría conocer tu monto acumulado inicial?:\n1) Sí.\n2) No.\nElige: ")
	fmt.Scanln(&ver_monto)

	nombre := "User"
	if ver_monto == 1 {
		solicitarMonto(sv, nombre)
	} else {
		decision_inicial = 0
	}

	mercenarios := []struct {
		nombre   string
		piso     int32
		decision int32
		vivo     bool
		aciertos int32
	}{
		{nombre: "User", piso: 1, decision: int32(decision_inicial), vivo: true, aciertos: 0},
		{nombre: "Reverend Alberts", piso: 1, decision: int32(rand.Intn(3) + 1), vivo: true, aciertos: 0},
		{nombre: "Tom Banner", piso: 1, decision: int32(rand.Intn(3) + 1), vivo: true, aciertos: 0},
		{nombre: "PC Rob Briar", piso: 1, decision: int32(rand.Intn(3) + 1), vivo: true, aciertos: 0},
		{nombre: "Classic Briar", piso: 1, decision: int32(rand.Intn(3) + 1), vivo: true, aciertos: 0},
		{nombre: "Mr. Foster", piso: 1, decision: int32(rand.Intn(3) + 1), vivo: true, aciertos: 0},
		{nombre: "Ana Larive", piso: 1, decision: int32(rand.Intn(3) + 1), vivo: true, aciertos: 0},
		{nombre: "Lt. Bill Masterson", piso: 1, decision: int32(rand.Intn(3) + 1), vivo: true, aciertos: 0},
	}

	var wg sync.WaitGroup

	for _, mercenario := range mercenarios {
		wg.Add(1)
		go enviarDecision(sv, &wg, mercenario)
	}

	wg.Wait()
	fmt.Println("Todas las goroutines han terminado.")
}
