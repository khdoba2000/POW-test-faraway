package handler

import (
	"context"
	"log"
	"net"
	"test-faraway/entity"
	"test-faraway/pkg/pow"
)

type Handler struct {
	MaxLengthChallenge int
	MinLengthChallenge int
	DifficultyLength   int
	SolutionLength     int
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
	_, err = conn.Write([]byte(challangeStr))
	if err != nil {
		log.Println("failed to send challenge message error:", err)
		return err
	}

	log.Println("challenge message send")

	solutionRespBytes := make([]byte, challenge.SolutionLength)
	n, err := conn.Read(solutionRespBytes)
	if err != nil {
		log.Println("failed to read from the server err:", err)
		return err
	}
	solution := solutionRespBytes[:n]

	log.Println("solution response:", solution)

	if pow.VerifySolution(challenge.Challenge, solution, challenge.Difficulty) {
		log.Println("verify solution succeeded")

		//write quote to the client
		conn.Write([]byte("quote"))
	} else {
		log.Println("verify solution failed")
	}

	return nil
}