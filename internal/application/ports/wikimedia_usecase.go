package ports

type WikiMediaUseCase interface {
	Search(query, language string) (interface{}, error)
}
