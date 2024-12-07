package domain

type Jwt interface {
	Read() (*string, error)
	Generate(metadata *string) (string)
}
