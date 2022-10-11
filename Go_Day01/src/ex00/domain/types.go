package domain

type DBReader interface {
	ConvertFile(file []byte) error
	GetRecipes() DBReader
	GetResult() ([]byte, error)
}
