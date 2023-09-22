package txt

import (
	"fmt"
	"log"
	wowIO "test-faraway/pkg/io"
	"test-faraway/repository"
)

type txtRepo struct {
	quotes map[int]string
}

// NewTxtRepo initializes a new quote repository scanning the txt file located in the given file path
func NewTxtRepo(filePath string) repository.Repo {
	lines, err := wowIO.ReadFile(filePath)
	if err != nil {
		log.Fatal("error reading file: ", err)
	}

	quotes := make(map[int]string, len(lines))
	for i, line := range lines {
		quote := line
		quotes[i] = quote
	}

	fmt.Println("quotes:", quotes)
	return &txtRepo{quotes: quotes}
}

func (r *txtRepo) GetQuoteByIndex(index int) (string, error) {
	quote, ok := r.quotes[index]
	if !ok {
		return "", fmt.Errorf("not found quote with index %d", index)
	}
	return quote, nil
}

func (r *txtRepo) GetQuotesCount() int {
	return len(r.quotes)
}
