package models

type LongPoolServerResponse struct {
	Response LongPoolServer `json:"response"`
}

// указатели на string использованы для того, чтобы удобно реализовать проверку на то, что у нас успешно распарсился json
// и всё классно записалось в структуру. Может же быть такое, что не распарсиллся json, тогда в структуре будут значения
// по умолчанию, пустые строки и всё сломается. В случае с указателем там будет nil
type LongPoolServer struct {
	Key    string `json:"key"`
	Server string `json:"server"`
	Ts     string `json:"ts"`
}
