package main

import (
	"context"
	"log"
	"net"
	"test-faraway/configs"
	"test-faraway/entity"
	"test-faraway/pkg/pow"
	"time"
)

func main() {
	cfg := configs.Config()
	serverAddress := cfg.ServerHost + ":" + cfg.ServerPort
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		log.Fatal("failed to connect to the server", err)
	}

	defer conn.Close()

	bytes := make([]byte, 1024)
	n, err := conn.Read(bytes)
	if err != nil {
		log.Fatal("failed to read from the server", err)
	}

	challenge := entity.Challenge{}
	err = challenge.DecodeFromBytes(bytes[:n])
	if err != nil {
		log.Fatal("failed to decode to challenge error:", err)
	}

	log.Println("Mining started...")

	timeOut, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(cfg.PowCalculationTimeSeconds))
	defer cancel()
	solutionChan := make(chan []byte)
	go func() {
		solution := pow.PerformPoW(challenge.Challenge, challenge.Difficulty, challenge.SolutionLength)
		log.Printf("PoW completed successfully!\nChallenge: %s\nSolution: %v\n", challenge.Challenge, solution)

		solutionChan <- solution
	}()

	select {
	case sol := <-solutionChan:
		_, err = conn.Write(sol)
		if err != nil {
			log.Fatal("failed to write challenge solution:", err)
			break
		}
		log.Print("solution written successfully")

		n, err = conn.Read(bytes)
		if err != nil {
			log.Fatal("failed to read server response:", err)
			break
		}
		log.Print("read message from server: ", string(bytes[:n]))

	case <-timeOut.Done():
		log.Print("pow calculation timed out.")
	}

}
