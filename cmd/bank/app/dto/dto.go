package dto

type CardDTO struct {
	Id int64 `json:"id"`
	Issuer string `json:"issuer"`
	Number string `json:"number"`
	Currency string `json:"currency"`
	Virtual  bool `json:"virtual"`
}
