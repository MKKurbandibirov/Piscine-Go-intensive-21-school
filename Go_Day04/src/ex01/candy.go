package ex00

type CandyRequest struct {
	Money      int    `json:"money"`
	CandyType  string `json:"candyType"`
	CandyCount int    `json:"candyCount"`
}

type CandyResponse struct {
	Change int    `json:"change"`
	Thanks string `json:"thanks"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type Candy struct {
	Name  string
	Price int
}
