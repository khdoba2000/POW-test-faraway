package controller

import (
	"log"
	"math/rand"
	"test-faraway/repository"
	"time"
)

const (
	defaultQuote = "Some beliefs are so strongly held that some believers are ready to lose their heads rather than the belief they carry therein."
)

type Controller struct {
	Repo repository.Repo
}

func (c *Controller) GetRandomWOW() string {

	quotestCount := c.Repo.GetQuotesCount()

	rand.New(rand.NewSource(time.Now().UnixNano()))
	quoteIndex := rand.Intn(quotestCount)

	quote, err := c.Repo.GetQuoteByIndex(quoteIndex)
	if err != nil {
		log.Printf("GetQuoteByIndex error: %v\n", err)
		log.Println("Returning defaultQuote")
		quote = defaultQuote
	}

	return quote
}
