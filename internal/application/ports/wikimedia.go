package ports

type WikiMedia interface {
	Search(query, language string) (interface{}, error)
}
