package ports

type Parser interface {
	Parse(content string) (string, error)
}
