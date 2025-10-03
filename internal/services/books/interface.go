package books

type BookService interface {
	GetRecommendations(text any) ([]any, error)
}
