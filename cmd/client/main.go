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

const (
	powCalculationTimeSeconds = 100
)

func main() {
	cfg := configs.Config()
	conn, err := net.Dial("tcp", cfg.ServerHost+":"+cfg.ServerPort)
	if err != nil {
		log.Fatal("failed to connect to the server", err)
	}

	defer conn.Close()

	bytes := make([]byte, 1024)
	n, err := conn.Read(bytes)
	if err != nil {
		log.Fatal("failed to read from the server", err)
	}

	log.Println("challenge response:", string(bytes[:n]))
	challenge := entity.Challenge{}
	err = challenge.DecodeFromBytes(bytes[:n])
	if err != nil {
		log.Fatal("failed to decode to challenge error:", err)
	}

	log.Println("Mining started...")

	timeOut, cancel := context.WithTimeout(context.Background(), time.Second*powCalculationTimeSeconds)
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
		break

	case <-timeOut.Done():
		log.Print("pow calculation timed out.")
		break
	}

}
