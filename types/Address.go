package types

type Address struct {
	Flat    string `json:"flat"`
	House   string `json:"house"`
	Street  string `json:"street"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}
