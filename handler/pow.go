package handler

import (
	"context"
	"fmt"
	"log"
	"net"
	"test-faraway/controller"
	"test-faraway/entity"
	"test-faraway/pkg/pow"
	"time"
)

type Handler struct {
	MaxLengthChallenge        int
	MinLengthChallenge        int
	DifficultyLength          int
	SolutionLength            int
	PowCalculationTimeSeconds int
	Controller                controller.Controller
}

func (h *Handler) Handle(ctx context.Context, conn net.Conn) error {

	defer conn.Close()
	go func() {
		<-ctx.Done()
		conn.Close()
	}()

	log.Println("received conn:", conn)

	challenge := entity.Challenge{
		Challenge:      pow.GenerateChallengeStr(h.MinLengthChallenge, h.MaxLengthChallenge),
		Difficulty:     h.DifficultyLength,
		SolutionLength: h.SolutionLength,
	}
	challangeStr, err := challenge.EncodeToString()
	if err != nil {
		log.Println("failed to encode challenge to string error:", err)
		return err
	}

	//send challenge to client
	_, err = conn.Write([]byte(challangeStr))
	if err != nil {
		log.Println("failed to send challenge message error:", err)
		return err
	}

	log.Println("challenge message send")

	// conn.SetReadDeadline(time.Now().Add(time.Second * time.Duration(h.PowCalculationTimeSeconds)))

	solutionChan := make(chan []byte)
	go func() {
		solutionRespBytes := make([]byte, h.SolutionLength)
		n, err := conn.Read(solutionRespBytes)
		if err != nil {
			log.Println("failed to read from the server err:", err)
			return
		}
		solution := solutionRespBytes[:n]
		solutionChan <- solution
	}()

	select {
	case solution := <-solutionChan:
		log.Println("solution response:", solution)
		if pow.VerifySolution(challenge.Challenge, solution, challenge.Difficulty) {
			log.Println("verify solution succeeded, wring WoW")
			err := h.WriteWOW(conn)
			if err != nil {
				return err
			}
		} else {
			log.Println("verify solution failed")
		}

	case <-time.After(time.Second * time.Duration(h.PowCalculationTimeSeconds)):
		log.Println("pow calculation timed out")
	}

	return nil
}

func (h Handler) WriteWOW(conn net.Conn) error {

	//get random quote
	quote := h.Controller.GetRandomWOW()

	//write quote to the client
	_, err := conn.Write([]byte(quote))
	if err != nil {
		return fmt.Errorf("error writing quote to client: %v", err)
	}

	return nil
}
