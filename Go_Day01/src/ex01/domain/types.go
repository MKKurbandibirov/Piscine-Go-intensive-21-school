package domain

type DBReader interface {
	ConvertFile(file []byte) error
	GetResult() ([]byte, error)
}
