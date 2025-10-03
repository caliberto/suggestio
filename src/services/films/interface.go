package films

type FilmService interface {
	GetRecommendations(text any) ([]any, error)
}
