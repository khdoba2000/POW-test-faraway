package repository

type Repo interface {
	GetQuoteByIndex(int) (string, error)
	GetQuotesCount() int
}
