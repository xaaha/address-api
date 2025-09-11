package db

// Address represents the structure of the "address" table in the databse
type Address struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	FullAddress string `json:"full_address"`
	Phone       string `json:"phone"`
	CountryCode string `json:"country_code"`
	County      string `json:"country"`
}
